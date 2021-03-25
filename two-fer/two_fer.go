//Package twofer takes a name and prints out "One for ____, one for me."
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
