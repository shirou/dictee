package dictionary

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

type DictType string

const (
	DictTypeApple    DictType = "apple"
	DictTypeStarDict          = "stardict"
	DictTypeEDICT             = "EDICT"
)

type Dictionary interface {
	Search(word string) ([]Title, error)
	Get(Title) (Entry, error)
	MakeIndex(string) error
	//	IsIndexCreated() bool
}

type Title struct {
	Title      string
	Offset     int64
	Size       int
	Dictionary Dictionary
}

type Entry interface {
	ToText() string // return text to show in terminal
}

// makeDictName returns filessytem friendly name from name.
// This name is used to store index data in config
// TODO
func makeDictName(name string) string {
	return name
}

func encodeOffsetSize(offset int64, size int) []byte {
	o := make([]byte, 4)
	binary.PutVarint(o, offset)
	s := make([]byte, 4)
	binary.PutVarint(s, int64(size))
	return append(o, s...)
}

func decodeOffsetSize(b []byte) (int64, int, error) {
	offset, n := binary.Varint(b[:3])
	if n < 1 {
		return 0, 0, errors.Errorf("offset decode failed, %d", n)
	}
	size, n := binary.Varint(b[4:])
	if n < 1 {
		return 0, 0, errors.Errorf("size decode failed, %d", n)
	}

	return offset, int(size), nil
}

type ConfigDictionary struct {
	Name        string   `json:"name"`
	DictType    DictType `json:"dict_type"`
	DisplayName string   `json:"display_name"`
	BodyPath    string   `json:"body_path"`
	IndexPath   string   `json:"index_path"`
}

func (c ConfigDictionary) NewDictionary() (dic Dictionary, err error) {
	switch c.DictType {
	case DictTypeApple:
		dic, err = NewAppleDictionary(c.Name, c.DictPath)
	default:
		return nil, errors.Errorf("unknown dict type: %s", c.DictType)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "New Dictonary failed: %s", c.DisplayName)
	}

}
