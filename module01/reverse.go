package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
//First solution
// func Reverse(word string) string {
// 	rev := ""
// 	for i := range word {
// 		rev = rev + string(word[len(word)-1-i])
// 	}
// 	return rev
// }

//Reverse function for string
func Reverse(word string) string {
	rev := ""
	for _, i := range word {
		rev = string(i) + rev
	}
	return rev
}
