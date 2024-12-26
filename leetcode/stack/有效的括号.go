package main

func main() {

}
func isValid(s string) bool {
	hash := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stake := make([]byte, 0)
	for i := range s {

		if value, ok := hash[s[i]]; ok {
			if len(stake) == 0 || stake[len(stake)-1] != value {
				return false
			}
			stake = stake[:len(stake)-1]
		} else {
			stake = append(stake, s[i])
		}
	}
	if len(stake) > 0 {
		return false
	}
	return true
}
