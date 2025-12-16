package asciiart

import (
	_ "embed"

	"github.com/romance-dev/ascii-art"
)

//go:embed font.flf
var font []byte

func init() {
	asciiart.RegisterFont("emboss-2", font)
}
