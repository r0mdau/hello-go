package __iteration

// repeat a string a given number of times
func Repeat(str string, times int) string {
	var repeatedStr string
	for i := 0; i < times; i++ {
		repeatedStr += str
	}
	return repeatedStr
	//return strings.Repeat(str, times)
}
