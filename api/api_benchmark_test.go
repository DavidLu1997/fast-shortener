package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davidlu1997/fast-shortener/model"
	"github.com/valyala/fasthttp"
)

func BenchmarkPutLink(b *testing.B) {
	api := testingAPI()

	ctxs := buildPutCtxs(b.N)

	b.ResetTimer()

	runCtxs(api, ctxs)
}

func BenchmarkGetLink(b *testing.B) {
	api := testingAPI()
	ctxs := buildPutCtxs(b.N)

	runCtxs(api, ctxs)

	ctxs = buildGetCtxs(b.N)

	b.ResetTimer()

	runCtxs(api, ctxs)
}

func runCtxs(api *API, ctxs []*fasthttp.RequestCtx) {
	for i := range ctxs {
		api.RequestHandler(ctxs[i])
	}
}

func buildPutCtxs(n int) []*fasthttp.RequestCtx {
	ctxs := make([]*fasthttp.RequestCtx, n)
	for i := range ctxs {
		link := model.Link{
			URL:     fmt.Sprintf("http://derp.com/get-%d", i),
			Key:     fmt.Sprintf("key-%d", i),
			Seconds: 30,
		}

		json, _ := json.Marshal(link)

		var req fasthttp.Request
		req.SetRequestURI("/put")
		req.SetBody(json)

		var ctx fasthttp.RequestCtx
		ctx.Init(&req, nil, nil)

		ctxs[i] = &ctx
	}

	return ctxs
}

func buildGetCtxs(n int) []*fasthttp.RequestCtx {
	ctxs := make([]*fasthttp.RequestCtx, n)
	for i := range ctxs {
		uri := fmt.Sprintf("/key-%d", i)
		var req fasthttp.Request
		req.SetRequestURI(uri)

		var ctx fasthttp.RequestCtx
		ctx.Init(&req, nil, nil)

		ctxs[i] = &ctx
	}

	return ctxs
}
