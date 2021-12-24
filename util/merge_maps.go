package util

func MergeMaps(maps ...map[string]string) map[string]string {
	result := map[string]string{}

	for _, singleMap := range maps {
		for key, value := range singleMap {
			result[key] = value
		}
	}

	return result
}
