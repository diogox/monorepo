package with

import (
	"fmt"

	"go.uber.org/zap"
)

type Field func() zap.Field

func String(key string, value string) Field {
	return func() zap.Field {
		return zap.String(key, value)
	}
}

func Stringer(key string, value fmt.Stringer) Field {
	return func() zap.Field {
		return zap.Stringer(key, value)
	}
}

func Bytes(key string, value []byte) Field {
	return func() zap.Field {
		return zap.ByteString(key, value)
	}
}

func Int(key string, value int) Field {
	return func() zap.Field {
		return zap.Int(key, value)
	}
}

func Float32(key string, value float32) Field {
	return func() zap.Field {
		return zap.Float32(key, value)
	}
}

func Float64(key string, value float64) Field {
	return func() zap.Field {
		return zap.Float64(key, value)
	}
}

func Bool(key string, value bool) Field {
	return func() zap.Field {
		return zap.Bool(key, value)
	}
}
