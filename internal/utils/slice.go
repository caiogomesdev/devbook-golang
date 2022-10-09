package utils

func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func SliceMap[E interface{}, T interface{}](slice []E, funcMap func(item E) T) []T {
  var result []T
  for _, item := range slice {
    result = append(result, funcMap(item))
  }
  return result
}
