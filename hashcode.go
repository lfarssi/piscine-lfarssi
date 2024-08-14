package piscine

func HashCode(dec string) string {
	var result string
	for _, ch := range dec {
		if ch >= 32 && ch <= 127 {
			result += string(int(ch) + len(dec)%127)
		} else if ch < 32 {
			result += string(int(ch) + len(dec)%127 + 33)
		}
	}
	return result
}
