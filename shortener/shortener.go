package shortener

import "github.com/davidlu1997/fast-shortener/model"

type Shortener interface {
	Get(key string) *model.Link
	Put(link *model.Link) error
}
