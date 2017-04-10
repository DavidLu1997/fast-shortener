package api

import (
	"encoding/json"
	"testing"

	"fmt"

	"github.com/davidlu1997/fast-shortener/config"
	"github.com/davidlu1997/fast-shortener/model"
	"github.com/valyala/fasthttp"
)

const configPath = "../config/config.json"
const numLinks = 10

func testingAPI() *API {
	config, err := config.GetConfig(configPath)
	if err != nil {
		return nil
	}

	return InitAPI(config)
}

func TestPutLink(t *testing.T) {
	api := testingAPI()
	if api == nil {
		t.Fatal("Failed to initialize testing API")
	}

	for i := 0; i < numLinks; i++ {
		go func(a int) {
			link := model.Link{
				URL:     fmt.Sprintf("http://derp.com/put-%d", a),
				Key:     fmt.Sprintf("key-%d", a),
				Seconds: 30,
			}

			json, err := json.Marshal(link)
			if err != nil {
				t.Fatalf("Error marshalling json: %s", err)
			}

			var req fasthttp.Request
			req.SetRequestURI("/put")
			req.SetBody(json)

			var ctx fasthttp.RequestCtx
			ctx.Init(&req, nil, nil)

			api.RequestHandler(&ctx)

			if status := ctx.Response.StatusCode(); status != fasthttp.StatusOK {
				t.Fatalf("Expected %d status got %d status", fasthttp.StatusOK, status)
			}
		}(i)
	}
}

func TestPutInvalidJson(t *testing.T) {
	api := testingAPI()
	if api == nil {
		t.Fatal("Failed to initialize testing API")
	}

	json := []byte{'{', 'a'}

	var req fasthttp.Request
	req.SetRequestURI("/put")
	req.SetBody(json)

	var ctx fasthttp.RequestCtx
	ctx.Init(&req, nil, nil)

	api.RequestHandler(&ctx)

	if status := ctx.Response.StatusCode(); status != fasthttp.StatusBadRequest {
		t.Fatalf("Expected %d status got %d status", fasthttp.StatusBadRequest, status)
	}
}

func TestPutInvalidLink(t *testing.T) {
	api := testingAPI()
	if api == nil {
		t.Fatal("Failed to initialize testing API")
	}

	link := model.Link{
		URL:     "abcd",
		Key:     "a",
		Seconds: 30,
	}

	json, err := json.Marshal(link)
	if err != nil {
		t.Fatalf("Error marshalling json: %s", err)
	}

	var req fasthttp.Request
	req.SetRequestURI("/put")
	req.SetBody(json)

	var ctx fasthttp.RequestCtx
	ctx.Init(&req, nil, nil)

	api.RequestHandler(&ctx)

	if status := ctx.Response.StatusCode(); status != fasthttp.StatusBadRequest {
		t.Fatalf("Expected %d status got %d status", fasthttp.StatusBadRequest, status)
	}
}

func TestGetLink(t *testing.T) {
	api := testingAPI()
	if api == nil {
		t.Fatal("Failed to initialize testing API")
	}

	successCh := make(chan int)
	for i := 0; i < numLinks; i++ {
		go func(a int) {
			link := model.Link{
				URL:     fmt.Sprintf("http://derp.com/get-%d", a),
				Key:     fmt.Sprintf("key-%d", a),
				Seconds: 30,
			}

			json, err := json.Marshal(link)
			if err != nil {
				t.Fatalf("Error marshalling json: %s", err)
			}

			var req fasthttp.Request
			req.SetRequestURI("/put")
			req.SetBody(json)

			var ctx fasthttp.RequestCtx
			ctx.Init(&req, nil, nil)

			api.RequestHandler(&ctx)

			if ctx.Response.StatusCode() != fasthttp.StatusOK {
				t.Fatal("wrong status")
			}

			successCh <- a
		}(i)
	}

	for i := 0; i < numLinks; i++ {
		a := <-successCh
		uri := fmt.Sprintf("/key-%d", a)
		var req fasthttp.Request
		req.SetRequestURI(uri)

		var ctx fasthttp.RequestCtx
		ctx.Init(&req, nil, nil)

		api.RequestHandler(&ctx)

		if status := ctx.Response.StatusCode(); status != fasthttp.StatusFound {
			t.Fatalf("Expected %d status got %d status", fasthttp.StatusFound, status)
		}

		url := fmt.Sprintf("http://derp.com/get-%d", a)
		if location := ctx.Response.Header.Peek("Location"); string(location) != url {
			t.Fatalf("Expected %s redirect, got %s", url, location)
		}
	}
}

func TestGetInvalidLink(t *testing.T) {
	api := testingAPI()
	if api == nil {
		t.Fatal("Failed to initialize testing API")
	}

	var req fasthttp.Request
	req.SetRequestURI("/does-not-exist")

	var ctx fasthttp.RequestCtx
	ctx.Init(&req, nil, nil)

	api.RequestHandler(&ctx)

	if status := ctx.Response.StatusCode(); status != fasthttp.StatusNotFound {
		t.Fatalf("Expected %d status got %d status", fasthttp.StatusOK, status)
	}

	if body := ctx.Response.Body(); len(body) > 0 {
		t.Fatal("Expected body to be empty")
	}
}

func TestHealth(t *testing.T) {
	api := testingAPI()
	if api == nil {
		t.Fatal("Failed to initialize testing API")
	}

	var req fasthttp.Request
	var ctx fasthttp.RequestCtx
	req.SetRequestURI("/ok")
	ctx.Init(&req, nil, nil)

	api.RequestHandler(&ctx)

	if status := ctx.Response.StatusCode(); status != fasthttp.StatusOK {
		t.Fatalf("Expected %d status got %d status", fasthttp.StatusOK, status)
	}
}
