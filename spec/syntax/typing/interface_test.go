package typing

import "testing"

func TestTypingInterfaceSpec(t *testing.T) {
	println(`
		InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
		InterfaceElem  = MethodElem | TypeElem .
		MethodElem     = MethodName Signature .
		MethodName     = identifier .
		TypeElem       = TypeTerm { "|" TypeTerm } .
		TypeTerm       = Type | UnderlyingType .
		UnderlyingType = "~" Type .
	`)
}
