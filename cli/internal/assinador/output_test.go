package assinador

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSignOutputs(t *testing.T) {
	dir := t.TempDir()
	result := `{"operation":"sign","status":"success","signature":"SIMULATED-abc","signatureFormat":"CMS","message":"ok"}`

	if err := WriteSignOutputs(dir, "doc", "/tmp/doc.pdf", result); err != nil {
		t.Fatal(err)
	}

	for _, name := range []string{"doc.assinado.json", "doc-transcricao-assinada.txt", "doc-comprovante-assinatura.txt"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatalf("arquivo ausente %s: %v", name, err)
		}
	}
}
