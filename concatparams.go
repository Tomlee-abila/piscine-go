package piscine

func ConcatParams(args []string) string {
	var result string
	result = args[0]
	for i := 1; i < len(args); i++ {
		result += "\n" + args[i]
	}
	return result
}
