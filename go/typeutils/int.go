package typeutils

func IntDefault(i int, defaultValue int) int {
	if i == 0 {
		return defaultValue
	}
	return i
}
func IntPtr(i int) *int {
	return &i
}
