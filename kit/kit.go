package kit

func DeleteItemForList(item string, list []string) []string {
	newString := make([]string, 0)
	for _, v := range list {
		if v != item {
			newString = append(newString, v)
		}
	}
	return newString
}
