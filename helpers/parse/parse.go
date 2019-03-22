package parse

import "strings"

// Unslugify takes a slug and returns a space separated string.
func Unslugify(orig string) (retval string) {
	retval = strings.Replace(orig, "-", " ", -1)
	retval = strings.Replace(retval, "+", " ", -1)
	return retval
}
