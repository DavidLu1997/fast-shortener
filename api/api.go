package api

import (
	"encoding/json"

	"bytes"

	"github.com/davidlu1997/fast-shortener/config"
	"github.com/davidlu1997/fast-shortener/model"
	"github.com/davidlu1997/fast-shortener/shortener"
	"github.com/valyala/fasthttp"
)

type API struct {
	shortener shortener.Shortener
	config    *config.Configuration
}

func InitAPI(config *config.Configuration) *API {
	return &API{
		shortener: shortener.InitCacheShortener(config),
		config:    config,
	}
}

func (a *API) putLinkHandler(ctx *fasthttp.RequestCtx) {
	var link model.Link
	if err := json.Unmarshal(ctx.PostBody(), &link); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	}

	if err := a.shortener.Put(&link); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (a *API) getLinkHandler(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()
	key := path[bytes.LastIndexByte(path, '/')+1:]
	link := a.shortener.Get(string(key))
	if link == nil {
		ctx.Error("", fasthttp.StatusNotFound)
		return
	}

	json, err := json.Marshal(link)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadGateway)
		return
	}

	ctx.SetBody(json)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (a *API) okHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
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
