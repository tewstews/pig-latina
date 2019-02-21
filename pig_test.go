package pig

import "testing"

func TestTranslate(t *testing.T) {
	var trs = &translator {
		VowelsSfx: "hay",
		ConsonantsSfx: "ay",
		IgnoreE: false,
	}

	v := trs.Translate("test")
	if v[0] != "Esttay" {
		t.Error("Expected estay, got ", v)
	}
}
