package servers

func Show(tenantId string, tokenId string, name string, server *Server) error {
	servers := &[]Server{}
	err := Get(tenantId, tokenId, servers)
	if err != nil {
		return err
	}

	for _, item := range *servers {
		if item.Metadata.Instance_name_tag == name {
			*server = item
		}
	}

	return nil
}
