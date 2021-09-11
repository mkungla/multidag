package multidag

func (pt PortType) MarshalJSON() ([]byte, error) {
	var t string
	switch pt {
	case PortIn:
		t = "\"in\""
	case PortOut:
		t = "\"out\""
	}
	return []byte(t), nil
}
