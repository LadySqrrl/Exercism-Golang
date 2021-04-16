//return a string in reverse order
package reverse

//reverse the order of the input
func Reverse(s string) string {
	revStr := ""

	for _, c := range s {
		revStr = string(c) + revStr
	}

	return revStr
}
