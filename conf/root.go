package conf

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"os/user"
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
type JsonAccess struct {
	Access JsonToken `json:"access"`
}
type JsonToken struct {
	Token ConfToken `json:"token"`
}

func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func Read() (Conf, error) {
	config := Conf{}
	_, err := toml.DecodeFile((HomeDir() + "/.config/conoha.toml"), &config)
	return config, err
}

func Write(config *Conf) error {
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err := encoder.Encode(config)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile((HomeDir() + "/.config/conoha.toml"), buffer.Bytes(), 0777)
	if err != nil {
		// ファイルの生成を試みる
		file, err := os.Create(HomeDir() + "/.config/conoha.toml")
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write(buffer.Bytes())
	}
	return nil
}
