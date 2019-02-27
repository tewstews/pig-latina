package pig

import (
	"testing"
)


// todo fix out of range panic how called "string [punctuation][punctuation] string" case
// todo fix case "word..ee[punctuation]", "word..e[punctuation]"

func TestTranslate(t *testing.T) {
	var trs = &translator{
		VowelsSfx:     "hay",
		ConsonantsSfx: "ay",
		IgnoreE:       false,
	}

	v := trs.Translate("test")
	if v[0] != "esttay" {
		t.Error("Expected estay, got ", v)
	}

	c := trs.Translate("Test")
	if c[0] != "Esttay" {
		t.Error("Expected Esttay, got ", c)
	}


	sentence := trs.Translate("This is test sentence: with!     Wrong punctuation!")
	if sentence[0] != "Isthay ishay esttay entencesay: ithway! Ongwray unctuationpay!" {
		t.Error("Expected [Isthay ishay estsay entenceshay: ithway! Ongwray unctuationpay!], got ", sentence)
	}

	// test ignore "e" letter

	trs.IgnoreE = true

	sentence = trs.Translate("This is test sentence: with!     Wrong punctuation!")
	if sentence[0] != "Isthay ishay esttay entencsay: ithway! Ongwray unctuationpay!" {
		t.Error("Expected [Isthay ishay estsay entencsay: ithway! Ongwray unctuationpay!], got ", sentence)
	}
}