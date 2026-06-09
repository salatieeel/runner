package assinador

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const demoCertPEM = `-----BEGIN CERTIFICATE-----
SIMULATED-DEMO-CERTIFICATE-FOR-EDUCATIONAL-USE-ONLY
-----END CERTIFICATE-----
`

// SignResponse representa a resposta JSON do assinador Java.
type SignResponse struct {
	Operation       string `json:"operation"`
	Status          string `json:"status"`
	Signature       string `json:"signature"`
	SignatureFormat string `json:"signatureFormat"`
	Message         string `json:"message"`
}

// EnsureDemoCertificate cria ou reutiliza um certificado simulado no diretório informado.
func EnsureDemoCertificate(dir string) (string, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	path := filepath.Join(dir, "certificado-demo.pem")
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	if err := os.WriteFile(path, []byte(demoCertPEM), 0644); err != nil {
		return "", err
	}
	return path, nil
}

// WriteSignOutputs grava os três arquivos prometidos no README após assinatura simulada.
func WriteSignOutputs(outputDir, baseName, inputPath, jsonResult string) error {
	resp, err := ParseSignResponse(jsonResult)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("assinatura não concluída: %s", jsonResult)
	}

	assinadoPath := filepath.Join(outputDir, baseName+".assinado.json")
	assinadoContent, _ := json.MarshalIndent(map[string]interface{}{
		"inputFile":       inputPath,
		"signature":       resp.Signature,
		"signatureFormat": resp.SignatureFormat,
		"signedAt":        time.Now().Format(time.RFC3339),
		"simulated":       true,
	}, "", "  ")
	if err := os.WriteFile(assinadoPath, assinadoContent, 0644); err != nil {
		return err
	}

	transcricao := fmt.Sprintf(`Transcrição assinada (simulada)
Documento: %s
Assinatura: %s
Formato: %s
Data: %s
`, inputPath, resp.Signature, resp.SignatureFormat, time.Now().Format("02/01/2006 15:04:05"))
	transcricaoPath := filepath.Join(outputDir, baseName+"-transcricao-assinada.txt")
	if err := os.WriteFile(transcricaoPath, []byte(transcricao), 0644); err != nil {
		return err
	}

	comprovante := fmt.Sprintf(`Comprovante de assinatura simulada
%s
Arquivo original: %s
Hash de assinatura: %s
Este comprovante é apenas para fins educacionais e de demonstração.
`, resp.Message, inputPath, resp.Signature)
	comprovantePath := filepath.Join(outputDir, baseName+"-comprovante-assinatura.txt")
	return os.WriteFile(comprovantePath, []byte(comprovante), 0644)
}
