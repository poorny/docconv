package docconv

import (
	"os"
	"reflect"
	"testing"
)

func TestConvertPDF(t *testing.T) {
	f, err := os.Open("testdata/pdf-test.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	_, meta, err := ConvertPDF(f)
	want := map[string]string{
		"Suspects":       "no",
		"Form":           "none",
		"File size":      "20597 bytes",
		"Page rot":       "0",
		"Title":          "PDF Test Page",
		"Producer":       "Acrobat Distiller 7.0.5 (Windows)",
		"Tagged":         "yes",
		"JavaScript":     "no",
		"Pages":          "1",
		"Author":         "Yukon Department of Education",
		"Creator":        "Acrobat PDFMaker 7.0.7 for Word",
		"PDF version":    "1.6",
		"UserProperties": "no",
		"Encrypted":      "yes (print:yes copy:no change:no addNotes:no algorithm:RC4)",
		"Page size":      "612 x 792 pts (letter)",
		"Optimized":      "yes",
		"CreationDate":   "D:20080604154400-07'00'",
		"ModDate":        "D:20080604154736-07'00'",
	}
	if !reflect.DeepEqual(meta, want) {
		t.Errorf("ConvertPDF() meta = %#v; not %#v", meta, want)
	}
}
