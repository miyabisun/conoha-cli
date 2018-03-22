package conoha

type Config struct {
	Auth  ConfigAuth
	Token ConfigToken
	Ssh   []ConfigSsh
}
type ConfigAuth struct {
	User     string `toml:"username"`
	Pass     string `toml:"password"`
	TenantId string `toml:"tenant_id"`
}
type ConfigToken struct {
	Id       string `json:"id" toml:"id"`
	IssuedAt string `json:"issued_at" toml:"issued_at"`
	Expires  string `json:"expires" toml:"expires"`
}
type ConfigSsh struct {
	Name string
	Path string
}
