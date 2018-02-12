package status

import s "github.com/miyabisun/conoha-cli/endpoints/servers"

func Save(server *s.Server) error {
	config := &Config{}
	config.Id = server.Id
	config.Name = server.Metadata.Instance_name_tag
	config.KeyName = server.KeyName
	config.ImageId = server.Image.Id
	config.FlavorId = server.Flavor.Id

	for _, addrList := range server.Addresses {
		for _, addr := range addrList {
			if addr.Version == 4 {
				config.IpAddr = addr.Addr
				break
			}
		}
		if config.IpAddr != "" {
			break
		}
	}
	return Write(config)
}
