package shortener

import "github.com/davidlu1997/fast-shortener/model"

// Shortener defines Get and Put operations
type Shortener interface {
	// Gets a link given a key, returns nil or a valid link
	Get(key string) *model.Link

	// Puts a link
	// returns error if link is invalid or already present
	Put(link *model.Link) error
}
