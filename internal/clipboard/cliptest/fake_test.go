package cliptest

import (
	"context"
	"errors"
	"testing"

	"github.com/aidanismail/air-clip/internal/clipboard"
)

func TestSnapshotOwnsData(t *testing.T) {
	ctx := t.Context()
	f := New(clipboard.Item{})
	if _, err := f.ReadSnapshot(ctx); !errors.Is(err, clipboard.ErrNoSupportedItem) {
		t.Fatalf("empty snapshot error = %v", err)
	}
	seq, err := f.Sequence(ctx)
	if err != nil {
		t.Fatal(err)
	}
	item := clipboard.Item{Kind: clipboard.Text, Data: []byte("fixture")}
	if err := f.CompareAndWrite(ctx, seq, item); err != nil {
		t.Fatal(err)
	}
	item.Data[0] = '?'
	first, err := f.ReadSnapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	first.Data[0] = '?'
	second, err := f.ReadSnapshot(ctx)
	if err != nil || second.Kind != clipboard.Text || string(second.Data) != "fixture" {
		t.Fatal("snapshot was modified through another slice or could not be read")
	}
}

func TestCompareAndWriteAllowsOneCommit(t *testing.T) {
	ctx := t.Context()
	f := New(clipboard.Item{Kind: clipboard.Text, Data: []byte("initial")})
	seq, err := f.Sequence(ctx)
	if err != nil {
		t.Fatal(err)
	}
	item := clipboard.Item{Kind: clipboard.Text, Data: []byte("replacement")}
	start := make(chan struct{})
	results := make(chan error, 2)
	for range 2 {
		go func() {
			<-start
			results <- f.CompareAndWrite(ctx, seq, item)
		}()
	}
	close(start)
	var applied, conflicts int
	for range 2 {
		err := <-results
		switch {
		case err == nil:
			applied++
		case errors.Is(err, clipboard.ErrChanged):
			conflicts++
		default:
			t.Fatalf("unexpected write error: %v", err)
		}
	}
	if applied != 1 || conflicts != 1 || f.Calls() != (Calls{Sequences: 1, WriteAttempts: 2, Writes: 1}) {
		t.Fatal("the same sequence did not yield one commit and one conflict")
	}
	newSequence, err := f.Sequence(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.CompareAndWrite(ctx, newSequence, item); err != nil {
		t.Fatalf("intentional identical content with a fresh sequence: %v", err)
	}
	if got := f.Calls().Writes; got != 2 {
		t.Fatalf("writes = %d, want 2", got)
	}
}

func TestCanceledWriteDoesNotCommit(t *testing.T) {
	f := New(clipboard.Item{})
	seq, err := f.Sequence(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if err := f.CompareAndWrite(ctx, seq, clipboard.Item{Kind: clipboard.Text, Data: []byte("fixture")}); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled write error = %v", err)
	}
	if got := f.Calls().Writes; got != 0 {
		t.Fatalf("committed %d canceled writes", got)
	}
}
