package dictee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/shirou/dictee/dictionary"
)

type Config struct {
	Path         string             `json:"-"`
	Root         string             `json:"-"`
	Dictionaries []ConfigDictionary `json:"dictionaries"`
}

type ConfigDictionary struct {
	Name        string              `json:"name"`
	DictType    dictionary.DictType `json:"dict_type"`
	DisplayName string              `json:"display_name"`
	DictPath    string              `json:"dict_path"`
}

func (c ConfigDictionary) NewDictionary(confRoot string) (dic dictionary.Dictionary, err error) {
	switch c.DictType {
	case dictionary.DictTypeApple:
		dic, err = dictionary.NewAppleDictionary(c.Name, c.DictPath, confRoot)
	default:
		return nil, errors.Errorf("unknown dict type: %s", c.DictType)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "New Dictonary failed: %s", c.DisplayName)
	}

	return dic, nil
}

func (c ConfigDictionary) GetIndexPath(confRoot string) string {
	return filepath.Join(confRoot, c.Name+dictionary.IndexSuffix)
}

func ReadConfig(path string) (*Config, Dictee, error) {
	if path == "" {
		path = getDefaultConfigPath()
	}
	log.Debugf("config path:%s", path)

	var config Config
	if fileExists(path) {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, Dictee{}, errors.Wrap(err, "ReadConfig readfile")
		}
		if err := json.Unmarshal(buf, &config); err != nil {
			return nil, Dictee{}, errors.Wrap(err, "ReadConfig unmarshal")
		}
	} else {
		config.Path = path
		config.Dictionaries = make([]ConfigDictionary, 0)
	}

	ret := Dictee{
		ConfigRoot:   filepath.Dir(path),
		Dictionaries: make([]dictionary.Dictionary, len(config.Dictionaries)),
	}

	for i, conf := range config.Dictionaries {
		t, err := conf.NewDictionary(path)
		if err != nil {
			return nil, ret, errors.Wrap(err, "new dictionary")
		}
		ret.Dictionaries[i] = t
	}
	config.Path = path
	config.Root = filepath.Dir(path)

	return &config, ret, nil
}

func (c *Config) Dump() error {
	j, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return errors.Wrap(err, "config dump")
	}

	log.Debugf("dump to %s", c.Path)

	return ioutil.WriteFile(c.Path, j, os.ModePerm)
}

func (c *Config) AddDictionary(name, dictType, dictPath string) (err error) {
	log.Debugf("add dict, %s, %s, path: %s", name, dictType, dictPath)

	var dict dictionary.Dictionary
	switch dictionary.DictType(dictType) {
	case dictionary.DictTypeApple:
		dict, err = dictionary.NewAppleDictionary(name, dictPath, c.Root)
		if err != nil {
			return errors.Wrap(err, "AddDictionary")
		}
	default:
		return fmt.Errorf("unknown dict type: %s", dictType)
	}
	log.Infof("%s is created. making index", name)

	if err := dict.MakeIndex(c.Root); err != nil {
		return errors.Wrap(err, "make index")
	}
	confDict := ConfigDictionary{
		Name:        name,
		DisplayName: name, // TODO:
		DictType:    dictionary.DictType(dictType),
		DictPath:    dictPath,
	}

	c.Dictionaries = append(c.Dictionaries, confDict)
	fmt.Println(c.Dictionaries)

	return nil
}

func (c *Config) RemoveDictionary(name string) error {
	log.Debugf("remove dict, %s", name)

	var target ConfigDictionary
	deleted := make([]ConfigDictionary, 0)
	for _, dict := range c.Dictionaries {
		if dict.Name == name {
			target = dict
			continue
		}
		deleted = append(deleted, dict)
	}
	if target.Name == "" {
		log.Errorf("no such dictionary: %s", name)
		// actually, this is not a error.
		return nil
	}

	if err := os.RemoveAll(target.GetIndexPath(c.Root)); err != nil {
		return errors.Wrap(err, "RemoveAll")
	}

	c.Dictionaries = deleted

	return nil
}
func (c *Config) ListDictionary() error {
	for _, dict := range c.Dictionaries {
		fmt.Println(dict)
	}

	return nil
}

func (c *Config) Detect(dictpath string) error {
	return nil
}

func getDefaultConfigPath() string {
	dir := filepath.Join(userHomeDir(), ".config", "dictee")

	return filepath.Join(dir, "config.json")
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
