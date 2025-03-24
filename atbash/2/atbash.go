package atbash

import "strings"

func Encode(s string) string {
	cipher := make(map[rune]rune, 26)
	for in, out := 'a', 'z'; in <= 'z' && out >= 'a'; in, out = in+1, out-1 {
		cipher[in] = out
	}

	var output []rune
	for _, value := range strings.ToLower(s) {
		switch {
		case '0' <= value && value <= '9':
			output = append(output, value)
		case 'a' <= value && value <= 'z':
			output = append(output, cipher[value])
		default:
		}
	}

	builder := strings.Builder{}

	for idx, value := range output {
		if idx%5 == 0 && idx != 0 {
			builder.WriteRune(' ')
		}
		builder.WriteRune(value)
	}

	return builder.String()
}
