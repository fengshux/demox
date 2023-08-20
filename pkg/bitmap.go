package pkg

import (
	"github.com/fengshux/bitmap"
)

var (
	Bitmap *bitmap.BitMap
)

func init() {

	Bitmap = bitmap.New(uint64(20000))
}
