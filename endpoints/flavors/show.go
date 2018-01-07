package flavors

func Show(tenantId string, tokenId string, name string, item *Flavor) error {
	items := &[]Flavor{}
	err := Get(tenantId, tokenId, items)
	if err != nil {
		return err
	}

	for _, flavor := range *items {
		if flavor.Name == name {
			*item = flavor
			break
		}
	}

	return nil
}
