package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeOffsetSize(t *testing.T) {
	o := int64(10)
	s := 9999999999

	b := encodeOffsetSize(o, s)
	do, ds, err := decodeOffsetSize(b)
	assert.Nil(t, err)
	assert.Equal(t, o, do)
	assert.Equal(t, s, ds)
}
