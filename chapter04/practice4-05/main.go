package main

func uniq(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	prev := ""
	i := 1
	for _, s := range slice {
		if s == prev {
			copy(slice[i-1:], slice[i:])
		} else {
			i++
		}
		prev = s
	}
	return slice[:i-1]
}
