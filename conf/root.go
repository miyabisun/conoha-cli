package conf

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	home "github.com/mitchellh/go-homedir"
)

type Conf struct {
	Auth  ConfAuth  `toml:"auth"`
	Token ConfToken `toml:"token"`
}
type ConfAuth struct {
	User     string `toml:"username"`
	Pass     string `toml:"password"`
	TenantId string `toml:"tenant_id"`
}
type ConfToken struct {
	Id       string `json:"id" toml:"id"`
	IssuedAt string `json:"issued_at" toml:"issued_at"`
	Expires  string `json:"expires" toml:"expires"`
}

func ConfigPath() string {
	config_path, err := home.Expand("~/.config/conoha.toml")
	if err != nil {
		panic(err)
	}
	return config_path
}

func Read() (Conf, error) {
	config := Conf{}
	_, err := toml.DecodeFile(ConfigPath(), &config)
	return config, err
}

func Write(config *Conf) error {
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err := encoder.Encode(config)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(ConfigPath(), buffer.Bytes(), 0777)
	if err != nil {
		// ファイルの生成を試みる
		file, err := os.Create(ConfigPath())
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write(buffer.Bytes())
	}
	return nil
}
