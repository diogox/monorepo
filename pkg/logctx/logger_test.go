package logctx_test

import (
	"context"
	"testing"

	"github.com/diogox/monorepo/pkg/logctx"
	"github.com/diogox/monorepo/pkg/logctx/with"
)

// TODO

func TestInfo(t *testing.T) {
	ctx := context.Background()

	logctx.Info(ctx, "log message",
		with.Bool("bool", true),
		with.Int("int", 1),
		with.Float32("float32", 1.1),
		with.Float64("float32", 1.2),
		with.String("string", "value"),
		with.Stringer("stringer", nil),
		with.Bytes("bytes", []byte("value")),
	)
}
