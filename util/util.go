package util

func Contains(search string, slice []string) bool {
	for _, e := range slice {
		if e == search {
			return true
		}
	}
	return false
}