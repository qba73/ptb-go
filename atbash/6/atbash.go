package atbash

import "strings"

var cipher = map[rune]rune{
	'a': 'z',
	'b': 'y',
	'c': 'x',
	'd': 'w',
	'e': 'v',
	'f': 'u',
	'g': 't',
	'h': 's',
	'i': 'r',
	'j': 'q',
	'k': 'p',
	'l': 'o',
	'm': 'n',
	'n': 'm',
	'o': 'l',
	'p': 'k',
	'q': 'j',
	'r': 'i',
	's': 'h',
	't': 'g',
	'u': 'f',
	'v': 'e',
	'w': 'd',
	'x': 'c',
	'y': 'b',
	'z': 'a',
}

func Encode(s string) string {
	output := make([]rune, 0, len(cipher))
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
	builder.Grow(len(output))

	for idx, value := range output {
		if idx%5 == 0 && idx != 0 {
			builder.WriteRune(' ')
		}
		builder.WriteRune(value)
	}

	return builder.String()
}
