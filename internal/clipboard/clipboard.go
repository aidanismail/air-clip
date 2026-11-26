// Package clipboard defines the boundary between transfers and clipboard storage.
package clipboard

import (
	"context"
	"errors"
)

type Kind string

const (
	Text  Kind = "text"
	Image Kind = "image"
)

// Item holds validated UTF-8 text or one PNG. The caller owns the returned data
// from a read, and an adapter must not retain the input slice from a write.
type Item struct {
	Kind Kind
	Data []byte
}

var (
	ErrChanged         = errors.New("clipboard changed")
	ErrNoSupportedItem = errors.New("no supported clipboard item")
)

// Clipboard exposes on-demand access. CompareAndWrite compares and writes
// under the same exclusive acquisition. Its caller validates the item first.
type Clipboard interface {
	ReadSnapshot(context.Context) (Item, error)
	Sequence(context.Context) (uint32, error)
	CompareAndWrite(context.Context, uint32, Item) error
}
