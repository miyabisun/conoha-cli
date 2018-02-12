package servers

type VmList struct {
	Servers []Server
}
type Server struct {
	Id        string
	Name      string
	Status    string
	KeyName   string `json:"key_name"`
	Created   string
	Updated   string
	Addresses map[string][]Address
	Image     struct {
		Id string
	}
	Flavor struct {
		Id string
	}
	Metadata struct {
		Instance_name_tag string
	}
}
type Address struct {
	Version int
	Addr    string
	Type    string `json:"OS-EXT-IPS:type"`
	MacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
}
