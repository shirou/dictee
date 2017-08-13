package dictee

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"github.com/shirou/dictee/dictionary"
)

type Config struct {
	Dictionaries []dictionary.ConfigDictionary `json:"dictionaries"`
}

func ReadConfig(path string) (Dictee, error) {
	if path == "" {
		path = getDefaultConfigPath()
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return Dictee{}, errors.Wrap(err, "ReadConfig readfile")
	}

	var config Config
	if err := json.Unmarshal(buf, &config); err != nil {
		return Dictee{}, errors.Wrap(err, "ReadConfig unmarshal")
	}

	ret := Dictee{
		ConfigRoot:   filepath.Dir(path),
		Dictionaries: make([]dictionary.Dictionary, len(config.Dictionaries)),
	}

	for i, conf := range config.Dictionaries {
		t, err := conf.NewDictionary()
		if err != nil {
			return ret, errors.Wrap(err, "new dictionary")
		}
		ret.Dictionaries[i] = t
	}

	return ret, nil
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
