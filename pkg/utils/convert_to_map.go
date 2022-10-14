package utils

func ConvertToMap(arr []string) map[string]struct{} {
	mp := map[string]struct{}{}

	for _, x := range arr {
		mp[x] = struct{}{}
	}

	return mp
}
