package funcs

// MultiArgsSum MultiArgs 不定参数
func MultiArgsSum(a int, args ...int) (sum int) {
	sum = a
	for _, v := range args {
		sum += v
	}
	return
}
