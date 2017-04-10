# Fast Shortener
[![CircleCI](https://circleci.com/gh/DavidLu1997/fast-shortener/tree/master.svg?style=svg&circle-token=913bd3241b1007590e4c743144cd8c598ce59220)](https://circleci.com/gh/DavidLu1997/fast-shortener/tree/master)
[![codecov](https://codecov.io/gh/DavidLu1997/fast-shortener/branch/master/graph/badge.svg?token=DN5spsSZGE)](https://codecov.io/gh/DavidLu1997/fast-shortener)
[![GoDoc](https://godoc.org/github.com/DavidLu1997/fast-shortener?status.svg)](https://godoc.org/github.com/DavidLu1997/fast-shortener)

A non-persistent fast URL-shortening service in Golang using [valaya/fasthttp](https://github.com/valyala/fasthttp) and [patrickmn/go-cache](https://github.com/patrickmn/go-cache).

## Why use this?

- Links only persist as long as needed
- Speed is good
- You don't like databases

## Usage

### PUT Link

To shorten https://google.com to `hostname/google` for 60 seconds, make a request to `hostname/put`
```json
{
  "url": "https://google.com",
  "key": "google",
  "seconds": 60
}
```

Prohibited keys:
- `put`
- `ok`

### GET Link

To be redirected to https://google.com after PUT-ing the above, make a request to `hostname/google`

### Configuration

See [config.json](https://github.com/DavidLu1997/fast-shortener/blob/master/config/config.json)

## Benchmarks

Ran locally on a mid-2015 15" MBP

```
BenchmarkPutLink-8   	            500000        2912 ns/op
BenchmarkGetLink-8   	           2000000	       848 ns/op
BenchmarkCacheShortenerPut-8   	 2000000	       727 ns/op
BenchmarkCacheShortenerGet-8   	10000000	       230 ns/op
```

## Development

Clone the repo: `go get github.com/davidlu1997/fast-shortener`

### Tests

- Run tests: `make test`
- Run benchmarks: `make benchmark`

### Run

- Run server: `make run`
- Check health: `curl http://localhost:8069/ok`

### Lint

- Lint actionable: `make lint`
- Lint all: `make lint-all`
