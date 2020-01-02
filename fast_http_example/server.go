package fast_http_example

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	foobar string
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	_, _ = fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func StartServer() {
	// pass bound struct method to fasthttp
	myHandler := &MyHandler{
		foobar: "foobar",
	}
	go fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)

	// pass plain function to fasthttp
	go  fasthttp.ListenAndServe(":8081", fastHTTPHandler)


}