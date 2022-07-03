package literal

import "testing"

func TestLiteralStringSpec(t *testing.T) {
	println(`
		string_lit             = raw_string_lit | interpreted_string_lit .
		raw_string_lit         = "` + "`" + `" { unicode_char | newline } "` + "`" + `" .
		interpreted_string_lit = "` + "`" + `" { unicode_value | byte_value } "` + "`" + `" .
	`)
}

func TestLiteralRawString(t *testing.T) {
	println(`raw spec "" nothing is impossible`)
}

func TestLiteralInterpreted(t *testing.T) {
	println("what can i do for \tyou \" \t ")
}
