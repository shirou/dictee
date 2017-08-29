package dictionary

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppleSearch(t *testing.T) {
	assert := assert.New(t)
	confRoot := "/tmp/"
	path := filepath.Join("testcase", "apple", "jp-eg.dictionary", "Contents")
	dict := NewAppleDictionary("test", path, confRoot)
	dict.IndexPath = "/tmp/test_index.db"

	if err := dict.MakeIndex(confRoot); err != nil {
		t.Fatal(err)
	}
	t.Run("search", func(t *testing.T) {
		ctx := context.Background()
		titles, err := dict.Search(ctx, "gol") // golang is not included
		assert.Nil(err)
		assert.Len(titles, 47)
		assert.Equal(int64(426308), titles[0].Offset)
		assert.Equal("gold", titles[0].Title)
	})

	t.Run("get", func(t *testing.T) {
		title := Title{
			Title:  "gold",
			Offset: 426308,
			Size:   32780,
		}
		entry, err := dict.Get(title)
		assert.Nil(err)
		assert.Equal(`<h1>gold</h1>
<p>
go-rudo, kin, kogane<br/>
</p>`, entry.ToText())
	})
}

func TestAppleEntry(t *testing.T) {
	assert := assert.New(t)
	path := filepath.Join("testcase", "apple", "entry_test.xml")
	b, err := ioutil.ReadFile(path)
	assert.Nil(err)

	var r AppleXMLEntry
	assert.Nil(xml.Unmarshal(b, &r))
	assert.Equal("gold", r.AttrTitle)
	ap := &AppleEntry{}
	ap.Body = r
	ap.Original = string(b)
	fmt.Println(ap.ToRich())
}
