package parse

import "strings"

// Unslugify takes a slug and returns a space separated string.
func Unslugify(orig string) string {
	retval := strings.Replace(orig, "-", " ", -1)
	retval = strings.Replace(retval, "+", " ", -1)
	return retval
}

// Slugify takes a space separated string and returns a slug.
func Slugify(orig string) string {
	retval := strings.Replace(orig, " ", "-", -1)
	return retval
}
