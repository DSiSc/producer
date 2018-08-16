package config

import (
	"strings"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

var ConfigName = "config.json"

func New(path string) Config {
	return Config{filePath: path}
}

type Config struct {
	filePath string
	maps map[string] interface{}
}

// Read the given json file.
func (config *Config) read() {
	if !filepath.IsAbs(config.filePath) {
		filePath, err := filepath.Abs(config.filePath)
		if err != nil {
			panic(err)
		}
		config.filePath = filePath
	}

	bts, err := ioutil.ReadFile(config.filePath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bts, &config.maps)

	if err != nil {
		panic(err)
	}
}

// If we want to get item in a stucture, which like this:
//{
//	"classs": {
//		"student":{
//			"name": "john"
//         }
//     }
//}
// { class: {}}
// You can get it by call Get("class.student.name")
func (config *Config) GetConfigItem(name string) interface{} {
	if config.maps == nil {
		config.read()
	}

	if config.maps == nil {
		return nil
	}

	keys := strings.Split(name, ".")
	length := len(keys)
	if length == 1 {
		return config.maps[name]
	}

	var ret interface{}
	for i := 0; i < length; i++ {
		if i == 0 {
			ret = config.maps[keys[i]]
			if ret == nil {
				return nil
			}
		} else {
			if m, ok := ret.(map[string]interface{}); ok {
				ret = m[keys[i]]
			} else {
				if length == i-1 {
					return ret
				}
				return nil
			}
		}
	}
	return ret
}
