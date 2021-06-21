package strings

import (
	"bytes"
	string "kazdream-test-task/string"
)

func ToLower(s string.String) string.String {
	var buf bytes.Buffer
	for _, b := range s.Bytes {
		if b >= 65 && b <= 90 {
			buf.WriteByte(b + 32)
		} else {
			buf.WriteByte(b)
		}
	}
	s.Bytes = buf.Bytes()
	return s
}

func SplitWords(s string.String) []string.String {
	var arr []string.String
	var buf bytes.Buffer
	for _, b := range s.Bytes {
		if b >= 97 && b <= 122 {
			buf.WriteByte(b)
		} else {
			word := buf.Bytes()
			if len(word) != 0 {
				arr = append(arr, string.String{
					Bytes: word,
				})
				buf = bytes.Buffer{}
			}
		}
	}
	return arr
}

func Equal(s1 string.String, s2 string.String) bool {
	return bytes.Equal(s1.Bytes, s2.Bytes)
}
