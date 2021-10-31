package node

func (n Node) TextContent_() string {

	s, err := n.TextContent()

	if err != nil {

		n.Debug(err.Error())
	}

	return s

}
