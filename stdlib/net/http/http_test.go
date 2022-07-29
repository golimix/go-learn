package stdlib_net_http

import "testing"

func TestHttpServe(t *testing.T) {
	http.handlerfunc("", indexhandler)
}
