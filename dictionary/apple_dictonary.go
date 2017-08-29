package dictionary

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var appleEntryRe = regexp.MustCompile("<d:entry (?Us:.+)</d:entry>")

type AppleDictionary struct {
	Name        string `json:"name"`
	DictType    string `json:"dict_type"`
	DisplayName string `json:"display_name"`
	BodyPath    string `json:"body_path"`
	IndexPath   string `json:"index_path"`
}

type AppleEntry struct {
	Title    string
	Body     AppleXMLEntry
	Offset   int64
	Size     int
	Original string
}

func (e *AppleEntry) ToText() string {
	if len(e.Body.AppleXMLSpan) == 0 {
		// if len is 0, assume it is just a HTML
		return e.truncateEntryTag(e.Original)
	}

	var buf bytes.Buffer
	for _, span := range e.Body.AppleXMLSpan {
		buf.WriteString(span.Text)
	}

	return buf.String()
}

func (e *AppleEntry) ToRich() string {
	if len(e.Body.AppleXMLSpan) == 0 {
		// if len is 0, assume it is just a HTML
		return e.truncateEntryTag(e.Original)
	}
	var buf bytes.Buffer
	for _, span := range e.Body.AppleXMLSpan {
		e.Dump(buf, span, true)
	}
	return buf.String()
}

func (e *AppleEntry) Dump(buf bytes.Buffer, span *AppleXMLSpan, rich bool) {
	for _, sp := range span.AppleXMLSpan {
		e.Dump(buf, sp, rich)
	}
	b, err := xml.MarshalIndent(span, "", " ")
	if err != nil {
		return
	}
	buf.Write(b)
}

func (e *AppleEntry) truncateEntryTag(original string) string {
	f := strings.Index(original, ">")
	return strings.TrimSpace(strings.Replace(original[f+1:], "</d:entry>", "", 1))
}

func NewAppleDictionary(name, dir, confRoot string) *AppleDictionary {
	body := filepath.Join(dir, "Body.data")
	index := filepath.Join(confRoot, name+IndexSuffix)

	return &AppleDictionary{
		Name:      name,
		BodyPath:  body,
		IndexPath: index,
	}
}

func (dict *AppleDictionary) GetName() string { return dict.Name }

func (dict *AppleDictionary) MakeIndex(root string) error {
	log.Debugf("make index: %s, body: %s", dict.IndexPath, dict.BodyPath)
	db, err := leveldb.OpenFile(dict.IndexPath, nil)
	defer db.Close()

	fstat, err := os.Stat(dict.BodyPath)
	if err != nil {
		return errors.Wrapf(err, "body stat failed: %s", dict.BodyPath)
	}
	max := int(fstat.Size())

	f, err := os.Open(dict.BodyPath)
	if err != nil {
		return errors.Wrapf(err, "body Open failed: %s", dict.BodyPath)
	}
	defer f.Close()

	lbuf := make([]byte, 4)
	_, err = f.ReadAt(lbuf, 0x40)
	if err != nil {
		return errors.Wrapf(err, "first read at failed")
	}
	limit := int64(0x40 + binary.LittleEndian.Uint32(lbuf))

	_, err = f.Seek(0x60, 0)
	if err != nil {
		return errors.Wrapf(err, "seek to 0x60 failed")
	}

	buf := make([]byte, 4)
	for {
		offset, err := f.Seek(0, 1)
		if offset >= limit {
			break
		}
		_, err = f.Read(buf)
		if err != nil {
			return errors.Wrapf(err, "read header")
		}
		size := binary.LittleEndian.Uint32(buf)
		body := make([]byte, size)
		_, err = f.Read(body)
		if err != nil {
			return errors.Wrapf(err, "read body")
		}
		r, err := decompress(body[8:])
		if err != nil {
			return errors.Wrapf(err, "decompress")
		}
		for _, x := range appleEntryRe.FindAll(r, -1) {
			var r AppleXMLEntry
			if err := xml.Unmarshal(x, &r); err != nil {
				return errors.Wrapf(err, "xml unmarshal failed")
			}
			v := encodeOffsetSize(offset, int(size))
			db.Put([]byte(strings.ToLower(r.AttrTitle)), v, nil)
		}
		IndexCallback(int(offset), max)
	}
	log.Infof("create index done")

	return nil
}

func (dict *AppleDictionary) Search(ctx context.Context, word string) ([]Title, error) {
	db, err := leveldb.OpenFile(dict.IndexPath, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "open failed:%s", dict.IndexPath)
	}
	defer db.Close()

	var ret []Title
	iter := db.NewIterator(util.BytesPrefix([]byte(word)), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		o, sz, err := decodeOffsetSize(value)
		if err != nil {
			return nil, errors.Wrap(err, "apple dict search")
		}

		ret = append(ret, Title{
			Title:      string(key),
			Offset:     o,
			Size:       sz,
			Dictionary: dict,
		})
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return nil, errors.Wrap(err, "apple dict search, iter")
	}

	return ret, nil
}

func (dict *AppleDictionary) Get(title Title) (Entry, error) {
	f, err := os.Open(dict.BodyPath)
	if err != nil {
		return nil, errors.Wrapf(err, "get failed: %v", title)
	}
	defer f.Close()

	_, err = f.Seek(0x60, 0)
	if err != nil {
		return nil, errors.Wrap(err, "seek")
	}

	buf := make([]byte, title.Size)
	// TODO: what is this 12?
	if _, err := f.ReadAt(buf, title.Offset+12); err != nil {
		return nil, errors.Wrapf(err, "get read failed: %v", title)
	}
	r, err := decompress(buf)
	if err != nil {
		return nil, errors.Wrapf(err, "get decompress failed: %v", title)
	}

	for _, x := range appleEntryRe.FindAll(r, -1) {
		var r AppleXMLEntry
		if err := xml.Unmarshal(x, &r); err != nil {
			return nil, errors.Wrapf(err, "xml unmarshal failed")
		}
		if strings.ToLower(r.AttrTitle) == title.Title {
			ret := &AppleEntry{
				Title:    r.AttrTitle,
				Body:     r,
				Offset:   title.Offset,
				Size:     title.Size,
				Original: string(x),
			}
			return ret, nil
		}
	}

	return nil, fmt.Errorf("title not found")
}

func decompress(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)
	z, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	p, err := ioutil.ReadAll(z)
	if err != nil {
		return nil, err
	}
	return p, nil
}

type AppleXMLEntry struct {
	AttrClass    string          `xml:"class,attr"  json:",omitempty"`
	AttrXmlnsD   string          `xml:"xmlns:d,attr"  json:",omitempty"`
	AttrId       string          `xml:"id,attr"  json:",omitempty"`
	AttrTitle    string          `xml:"title,attr"  json:",omitempty"`
	AppleXMLSpan []*AppleXMLSpan `xml:"span,omitempty" json:"span,omitempty"`
}

type AppleXMLSpan struct {
	AttrClass    string          `xml:"class,attr"  json:",omitempty"`
	AttrDhw      string          `xml:"dhw,attr"  json:",omitempty"`
	AttrId       string          `xml:"id,attr"  json:",omitempty"`
	AttrOrd      string          `xml:"ord,attr"  json:",omitempty"`
	AttrPr       string          `xml:"pr,attr"  json:",omitempty"`
	AttrRole     string          `xml:"role,attr"  json:",omitempty"`
	AppleXMLSpan []*AppleXMLSpan `xml:"span,omitempty" json:"span,omitempty"`
	Text         string          `xml:",chardata" json:",omitempty"`
}

func (span *AppleXMLSpan) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch span.AttrClass {
	case "gramb":
		start.Name = xml.Name{Space: "sp", Local: "h1"}
	default:
		start.Name = xml.Name{Space: "sp", Local: "vvv"}
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	e.EncodeToken(span.Text)
	for _, s := range span.AppleXMLSpan {
		if err := e.Encode(s); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End)
}
