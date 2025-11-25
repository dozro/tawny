package embed

func truncateString(str string, num int) string {
	truncated := ""
	count := 0
	for _, char := range str {
		truncated += string(char)
		count++
		if count >= num {
			break
		}
	}
	return truncated
}
