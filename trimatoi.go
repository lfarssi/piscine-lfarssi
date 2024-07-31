package piscine

func TrimAtoi(s string) int {
	r := []int{}
	bl := true
	hadnumber := false
	counter := 0
	for _, i := range s {
		if i >= '0' && i <= '9' {
			if i == '0' && !hadnumber {
				continue
			}
			r = append(r, int(i-'0'))
			hadnumber = true
			counter++
		} else if i == '-' && !hadnumber {
			bl = false
		}
	}
	op := 1
	res := 0
	for j := counter - 1; j >= 0; j-- {
		res += r[j] * op
		op *= 10
	}
	if !bl {
		res *= -1
	}
	return res
}
