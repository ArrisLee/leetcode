package checkonessegment

func checkOnesSegment(s string) bool {
	if len(s) == 1 {
		return true
	}
	exits := false
	for i := range s {
		if exits && s[i] == '1' {
			return false
		}
		if i < len(s)-1 && s[i] == '1' && s[i+1] == '0' {
			exits = true
		}
	}
	return true
}
