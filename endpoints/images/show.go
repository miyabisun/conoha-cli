package images

func Show(tokenId string, name string, item *Image) error {
	items := &[]Image{}
	err := Get(tokenId, items)
	if err != nil {
		return err
	}

	for _, it := range *items {
		if it.Name == name {
			*item = it
			break
		}
	}

	return nil
}
