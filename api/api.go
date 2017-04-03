package api

import (
	"net/http"

	"github.com/davidlu1997/fast-shortener/shortener"
	"github.com/valyala/fasthttp"
)

type API struct {
	shortener shortener.Shortener
}

func InitAPI() *API {
	return &API{
		shortener: shortener.Shortener{},
	}
}

func (a *API) putLinkHandler(ctx *fasthttp.RequestCtx) {

}

func (a *API) getLinkHandler(ctx *fasthttp.RequestCtx) {

}

func (a *API) okHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(http.StatusOK)
}

func (a *API) RequestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/put":
		a.putLinkHandler(ctx)
	case "/ok":
		a.okHandler(ctx)
	default:
		a.getLinkHandler(ctx)
	}
}