package ascii

import (
	"testing"
)

func TestExtendedASCIIText(t *testing.T) {
	text := " € ÷ ¾ dollar "

	s := ExtendedASCIIText(text)
	for _, c := range s {
		if c < 0x80 {
			t.Error("This string has characters from ASCII ")
		}
	}
}