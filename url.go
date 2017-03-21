package raiponce

import "strings"

func buildURI(a ...string) string {
	return strings.Join(a, "/")
}
