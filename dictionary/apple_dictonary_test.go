package dictionary

import (
	"fmt"
	"testing"
)

func TestAppleDictionaryMakeIndex(t *testing.T) {
	dict := NewAppleDictionary("test", "/System/Library/Assets/com_apple_MobileAsset_DictionaryServices_dictionaryOSX/1a184f41ee27bbb203ba596a25c8ea104c13d98f.asset/AssetData/Sanseido The WISDOM English-Japanese Japanese-English Dictionary.dictionary/Contents/Resources/")

	if err := dict.MakeIndex(); err != nil {
		t.Fatal(err)
	}
}

func TestAppleDictionarySearch(t *testing.T) {
	a := &AppleDictionary{
		IndexPath: "/tmp/hoge.db",
	}

	fmt.Println(a.Search("a"))
}

func TestAppleDictionaryGet(t *testing.T) {
	dict := NewAppleDictionary("test", "/System/Library/Assets/com_apple_MobileAsset_DictionaryServices_dictionaryOSX/1a184f41ee27bbb203ba596a25c8ea104c13d98f.asset/AssetData/Sanseido The WISDOM English-Japanese Japanese-English Dictionary.dictionary/Contents/Resources/")

	title := Title{
		Offset: 784027,
		Size:   32790,
	}
	//	332127,32811
	// 96,33017

	entry, err := dict.Get(title)
	fmt.Println(entry, err)
}
