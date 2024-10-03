package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := encode([]string{"neet", "code", "love", "you"})
	fmt.Println(str)
	decoded := decode(str)
	fmt.Println(decoded)
}

func encode(strs []string) string {

	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteRune('#')
		sb.WriteString(str)
	}

	return sb.String()
}

func decode(str string) []string {

	var res []string

	encoded := []rune(str)

	for i := 0; i < len(encoded); i++ {

		var length strings.Builder

		for encoded[i] != '#' {
			length.WriteRune(encoded[i])
			i++
		}

		strLen, _ := strconv.Atoi(length.String())

		res = append(res, string(encoded[i+1:i+strLen+1]))
		i += strLen
	}

	return res
}
