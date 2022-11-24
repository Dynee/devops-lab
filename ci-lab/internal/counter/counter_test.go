package counter_test

import (
	"errors"
	"log"
	"testing"

	"dynee/cilab/internal/counter"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name    string
		items   []any
		wantLen int
	}{
		{
			"ten integers",
			[]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
		},
		{
			"5 strings",
			[]any{"one", "two", "three", "four", "five"},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantLen != counter.Count(tt.items) {
				log.Fatal(errors.New("unexpected count"))
			}
		})
	}
}
