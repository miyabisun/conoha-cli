package conoha

func TokenId() string {
	err := Refresh()
	if err != nil {
		panic(err)
	}

	var config Config
	err = Read(&config)
	if err != nil {
		panic(err)
	}

	return config.Token.Id
}
