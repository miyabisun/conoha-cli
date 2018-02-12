package status

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Id       string
	Name     string
	KeyName  string
	IpAddr   string
	ImageId  string
	FlavorId string
}

func path() string {
	return "status.toml"
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
