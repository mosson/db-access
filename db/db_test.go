package db

import (
	"testing"

	"github.com/naoina/genmai"
)

func TestDial(t *testing.T) {
	Close() // 安全に呼び出し可能

	Dial(&genmai.SQLite3Dialect{}, ":memory:")
	defer Close()

	ptr := GetClient()

	if GetClient() == nil {
		t.Errorf("expected initialized, and can be accessed")
	}

	// 何回呼んでも初期化は一回
	Dial(&genmai.SQLite3Dialect{}, ":memory:")
	Dial(&genmai.SQLite3Dialect{}, ":memory:")

	ptr2 := GetClient()

	if ptr != ptr2 {
		t.Errorf("expected be equal")
	}
}
