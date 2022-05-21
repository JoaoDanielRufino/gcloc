package utils

func ConvertToMap(arr []string) map[string]bool {
	mp := map[string]bool{}

	for _, x := range arr {
		mp[x] = true
	}

	return mp
}
