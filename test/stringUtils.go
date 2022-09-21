package main

func parseTabe(tab []string) string {
	if len(tab) <= 0 {
		return ""
	}
	return tab[0] + map[bool]string{true: " ", false: ""}[len(tab) > 1] + parseTabe(tab[1:])
}

func StrIndex(s, find string) int {
	s = ToLower(s)
	find = ToLower(find)
	SizeS := len([]rune(s))
	SizeF := len([]rune(find))
	for i := 0; i <= SizeS-SizeF; i++ {
		if s[i:i+SizeF] == find {
			return i
		}
	}
	return -1
}

func ToLower(s string) string {
	str := ""
	for _, c := range s {
		if c >= 65 && c <= 90 {
			str += string(c + 32)
		} else {
			str += string(c)
		}
	}
	return str
}

func Replace(target, torep, rep string) string {
	tab := Split(target, torep)
	str := ""
	for i, s := range tab {
		if i < len(tab)-1 {
			str += s + torep
		} else {
			str += s
		}
	}
	return str
}
