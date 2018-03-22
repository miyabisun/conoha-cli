package conoha

func SshPath(name string) string {
	config := &Config{}
	if err := Read(config); err != nil {
		return ""
	}

	keypair := &ConfigSsh{}
	for _, it := range config.Ssh {
		if it.Name == name {
			*keypair = it
		}
	}

	return keypair.Path
}
