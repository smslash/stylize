package pkg

import (
	"strings"
)

func Logic(banner, input string, file []byte) string {
	text := ""

	for i := 0; i < len(input); i++ {
		if input[i] != 13 {
			text += string(input[i])
		}
	}

	if len(text) == 0 {
		return ""
	}

	replaced := strings.ReplaceAll(string(file), "\n\n", "\n")
	splited := strings.Split(replaced, "\n")

	ascii := make(map[byte]int)
	var q byte
	for q = 32; q <= 126; q++ {
		ascii[q] = (int(q)-32)*8 + 1
	}

	var newLine []int
	counter := 0
	for i := 0; i < len(text); i++ {
		if text[i] == 10 {
			newLine = append(newLine, i-1)
			counter++
		}
	}

	var result []string
	if counter == 0 {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(text); j++ {
				result = append(result, splited[ascii[text[j]]+i])
			}
			result = append(result, "\n")
		}
	} else {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(text); j++ {
				if text[j] == 10 {
					break
				}
				result = append(result, splited[ascii[text[j]]+i])
			}
			result = append(result, "\n")
		}
		for k := 0; k < len(newLine); k++ {
			for i := 0; i < 8; i++ {
				for j := newLine[k] + 2; j < len(text); j++ {
					if text[j] == 10 {
						break
					}
					result = append(result, splited[ascii[text[j]]+i])
				}
				result = append(result, "\n")
			}
		}
	}

	res := ""
	for i := 0; i < len(result); i++ {
		res += result[i]
	}

	answer := strings.ReplaceAll(res, "\n\n\n\n\n\n\n\n", "\n")

	// counter_new := 0
	// for i := 0; i < len(text); i++ {
	// 	if i+1 < len(text) && text[i] == 10 {
	// 		counter_new++
	// 		i++
	// 	} else {
	// 		counter_new = 0
	// 		break
	// 	}
	// }

	// new_line_string := ""
	// if counter_new == 0 {
	// 	return answer
	// } else if counter_new == len(text)/2 {
	// 	for i := 0; i < counter_new; i++ {
	// 		new_line_string += "\n"
	// 	}
	// 	return new_line_string
	// }

	return answer
}
