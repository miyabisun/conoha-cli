package images

func Show(tokenId string, name string, item *Image) error {
	items := &[]Image{}
	err := Get(tokenId, items)
	if err != nil {
		return err
	}

	for _, image := range *items {
		if image.Name == name {
			*item = image
			break
		}
	}

	return nil
}
