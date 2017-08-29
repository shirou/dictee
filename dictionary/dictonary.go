package dictionary

import (
	"context"
	"encoding/binary"

	"github.com/pkg/errors"
)

type DictType string

const (
	DictTypeApple    DictType = "apple"
	DictTypeStarDict          = "stardict"
	DictTypeEDICT             = "EDICT"
)

const IndexSuffix = "_index.db"

type Dictionary interface {
	Search(ctx context.Context, word string) ([]Title, error)
	Get(Title) (Entry, error)
	MakeIndex(string) error
	GetName() string
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
	ToRich() string // return Rich to show in QT
}

// makeDictName returns filessytem friendly name from name.
// This name is used to store index data in config
// TODO
func makeDictName(name string) string {
	return name
}

func encodeOffsetSize(offset int64, size int) []byte {
	o := make([]byte, 8)
	binary.PutVarint(o, offset)
	s := make([]byte, 8)
	binary.PutVarint(s, int64(size))
	return append(o, s...)
}

func decodeOffsetSize(b []byte) (int64, int, error) {
	offset, _ := binary.Varint(b[:7])
	size, n := binary.Varint(b[8:])
	if n < 1 {
		return 0, 0, errors.Errorf("size decode failed, %d", n)
	}

	return offset, int(size), nil
}
