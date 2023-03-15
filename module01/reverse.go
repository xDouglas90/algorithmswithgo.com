package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) (reversed string) {
	reversed = ""
	for _, v := range word {
		reversed = string(v) + reversed
	}
	return
}
