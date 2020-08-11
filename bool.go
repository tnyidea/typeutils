package typeutils

func BoolPtr(b bool) *bool {
	return &b
}

func BoolYNString(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

func BoolYesNoString(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}
