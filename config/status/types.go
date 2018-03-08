package status

type Config struct {
	Id       string
	Name     string
	KeyName  string
	IpAddr   string
	ImageId  string
	FlavorId string
}

type SshConfig struct {
	Name         string
	HostName     string
	User         string
	KeyName      string
	IdentityFile string
}
