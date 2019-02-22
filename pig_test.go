package pig

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	var trs = &translator {
		VowelsSfx: "hay",
		ConsonantsSfx: "ay",
		IgnoreE: false,
	}

	v := trs.Translate("test")
	if v[0] != "esttay" {
		t.Error("Expected estay, got ", v)
	}

	c := trs.Translate("Test")
	fmt.Println(c)

	if c[0] != "Esttay" {
		t.Error("Expected Esttay, got ", c)
	}
}
