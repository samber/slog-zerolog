package slogzerolog

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/rs/zerolog"
)

func TestZerologHandler_WithGroupSiblings(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)

	root := slog.New(Option{Logger: &logger, NoTimestamp: true}.NewZerologHandler())
	parent := root.WithGroup("a").WithGroup("b").WithGroup("c")

	first := parent.WithGroup("d")
	second := parent.WithGroup("e")

	first.Info("first", "k", 1)
	second.Info("second", "k", 2)

	dec := json.NewDecoder(&buf)

	var got map[string]any
	if err := dec.Decode(&got); err != nil {
		t.Fatal(err)
	}
	if _, ok := got["a"].(map[string]any)["b"].(map[string]any)["c"].(map[string]any)["d"]; !ok {
		t.Fatalf("first record lost group d: %v", got)
	}

	got = nil
	if err := dec.Decode(&got); err != nil {
		t.Fatal(err)
	}
	if _, ok := got["a"].(map[string]any)["b"].(map[string]any)["c"].(map[string]any)["e"]; !ok {
		t.Fatalf("second record lost group e: %v", got)
	}
}
