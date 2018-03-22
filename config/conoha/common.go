package conoha

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	home "github.com/mitchellh/go-homedir"
)

func path() string {
	config_path, err := home.Expand("~/.config/conoha.toml")
	if err != nil {
		panic(err)
	}
	return config_path
}

func Read(config *Config) error {
	_, err := toml.DecodeFile(path(), &config)
	return err
}

func Write(config *Config) error {
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err := encoder.Encode(config)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path(), buffer.Bytes(), 0777)
	if err != nil {
		// ファイルの生成を試みる
		file, err := os.Create(path())
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write(buffer.Bytes())
	}
	return nil
}
