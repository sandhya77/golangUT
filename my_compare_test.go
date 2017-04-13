package strings


import (

	. "strings"

	"testing"

)


var compareTests = []struct {

	a, b string

	i    int

}{

	{"", "", 0},

	{"a", "", 1},

	{"", "a", -1},

	{"abc", "abc", 0},

	{"ab", "abc", -1},

	{"abc", "ab", 1},

	{"x", "ab", 1},

	{"ab", "x", -1},

	{"x", "a", 1},

	{"b", "x", -1},

	// test runtime·memeq's chunked implementation



