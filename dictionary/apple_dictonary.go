package dictionary

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type AppleDictionary struct {
	Name        string `json:"name"`
	DictType    string `json:"dict_type"`
	DisplayName string `json:"display_name"`
	BodyPath    string `json:"body_path"`
	IndexPath   string `json:"index_path"`
}

type AppleEntry struct {
	Title  string
	Body   AppleXMLEntry
	Offset int64
	Size   int
}

func (e *AppleEntry) ToText() string {
	return ""
}

func NewAppleDictionary(name, dir string) (*AppleDictionary, error) {
	body := filepath.Join(dir, "Body.data")

	return &AppleDictionary{
		Name:     name,
		BodyPath: body,
	}, nil
}

func (dict *AppleDictionary) MakeIndex(root string) error {
	db, err := leveldb.OpenFile(dbpath, nil)
	defer db.Close()

	f, err := os.Open(dict.BodyPath)
	if err != nil {
		return err
	}
	defer f.Close()

	lbuf := make([]byte, 4)
	_, err = f.ReadAt(lbuf, 0x40)
	if err != nil {
		return err
	}
	limit := int64(0x40 + binary.LittleEndian.Uint32(lbuf))

	_, err = f.Seek(0x60, 0)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`<d:entry(.+)</d.entry>`)

	buf := make([]byte, 4)
	var count int
	for {
		offset, err := f.Seek(0, 1)
		if offset >= limit {
			break
		}
		_, err = f.Read(buf)
		if err != nil {
			return err
		}
		size := binary.LittleEndian.Uint32(buf)
		body := make([]byte, size)
		_, err = f.Read(body)
		if err != nil {
			return err
		}
		r, err := decompress(body[8:])
		if err != nil {
			return err
		}
		for _, x := range re.FindAll(r, -1) {
			var r AppleXMLEntry
			if err := xml.Unmarshal(x, &r); err != nil {
				return errors.Wrapf(err, "xml unmarshal failed")
			}

			//			fmt.Printf("%s; %d,%d\n", r.AttrTitle, offset, size)
			v := encodeOffsetSize(offset, int(size))
			db.Put([]byte(strings.ToLower(r.AttrTitle)), v, nil)
		}
		count++
		if count%100 == 0 {
			fmt.Print(".")
		}
	}

	return nil
}

func (dict *AppleDictionary) Search(word string) ([]Title, error) {
	db, err := leveldb.OpenFile(dict.IndexPath, nil)
	defer db.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "open failed:%s", dict.IndexPath)
	}

	var ret []Title
	iter := db.NewIterator(util.BytesPrefix([]byte(word)), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Println(value)
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

		fmt.Println(string(key), string(value))
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
		return nil, err
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
	fmt.Println(string(r))

	return &AppleEntry{}, nil
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
