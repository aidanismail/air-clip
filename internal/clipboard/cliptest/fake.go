// Package cliptest supplies a synthetic clipboard for tests.
package cliptest

import (
	"bytes"
	"context"
	"sync"

	"github.com/aidanismail/air-clip/internal/clipboard"
)

type Calls struct {
	Reads, Sequences, WriteAttempts, Writes int
}

type Fake struct {
	mu       sync.Mutex
	item     clipboard.Item
	sequence uint32
	calls    Calls
}

var _ clipboard.Clipboard = (*Fake)(nil)

func New(initial clipboard.Item) *Fake {
	initial.Data = bytes.Clone(initial.Data)
	return &Fake{item: initial, sequence: 1}
}

func (f *Fake) ReadSnapshot(ctx context.Context) (clipboard.Item, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls.Reads++
	if err := ctx.Err(); err != nil {
		return clipboard.Item{}, err
	}
	if f.item.Kind == "" {
		return clipboard.Item{}, clipboard.ErrNoSupportedItem
	}
	item := f.item
	item.Data = bytes.Clone(item.Data)
	return item, nil
}

func (f *Fake) Sequence(ctx context.Context) (uint32, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls.Sequences++
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	return f.sequence, nil
}

func (f *Fake) CompareAndWrite(ctx context.Context, expected uint32, item clipboard.Item) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls.WriteAttempts++
	if err := ctx.Err(); err != nil {
		return err
	}
	if f.sequence != expected {
		return clipboard.ErrChanged
	}
	item.Data = bytes.Clone(item.Data)
	f.item = item
	f.sequence++
	f.calls.Writes++
	return nil
}

func (f *Fake) Calls() Calls {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls
}
