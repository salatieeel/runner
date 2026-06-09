package assinador

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kyriosdata/assinatura/internal/jdk"
)

type SignRequest struct {
	InputFile       string `json:"inputFile"`
	CertificateFile string `json:"certificateFile"`
	Pkcs11Library   string `json:"pkcs11Library,omitempty"`
	Pin             string `json:"pin,omitempty"`
	Alias           string `json:"alias,omitempty"`
	SignatureFormat string `json:"signatureFormat,omitempty"`
}

type ValidateRequest struct {
	InputFile       string `json:"inputFile"`
	CertificateFile string `json:"certificateFile"`
	SignatureFile   string `json:"signatureFile"`
}

func getAssinadorJarPath() (string, error) {
	var searchDirs []string

	if cwd, err := os.Getwd(); err == nil {
		searchDirs = append(searchDirs, cwd)
	}

	if ex, err := os.Executable(); err == nil {
		searchDirs = append(searchDirs, filepath.Dir(ex))
		searchDirs = append(searchDirs, filepath.Join(filepath.Dir(ex), ".."))
	}

	seen := make(map[string]bool)
	for _, dir := range searchDirs {
		dir, err := filepath.Abs(dir)
		if err != nil || seen[dir] {
			continue
		}
		seen[dir] = true

		for _, sub := range []string{
			filepath.Join(dir, "assinador", "target"),
			filepath.Join(dir, "assinador", "build"),
		} {
			if jar, err := findAssinadorJar(sub); err == nil {
				return jar, nil
			}
		}
	}

	return "", fmt.Errorf("assinador.jar não encontrado. Execute: npm run build:assinador")
}

func findAssinadorJar(dir string) (string, error) {
	exact := filepath.Join(dir, "assinador.jar")
	if _, err := os.Stat(exact); err == nil {
		return filepath.Abs(exact)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".jar") {
			continue
		}
		if strings.HasPrefix(name, "assinador") && !strings.Contains(name, "sources") && !strings.Contains(name, "javadoc") {
			return filepath.Abs(filepath.Join(dir, name))
		}
	}

	return "", fmt.Errorf("jar não encontrado em %s", dir)
}

func ExecuteSign(req SignRequest) (string, error) {
	return executeLocal("sign", buildSignArgs(req))
}

func ExecuteValidate(req ValidateRequest) (string, error) {
	return executeLocal("validate", buildValidateArgs(req))
}

func buildSignArgs(req SignRequest) []string {
	args := []string{"--input-file", req.InputFile, "--certificate", req.CertificateFile}
	if req.Pkcs11Library != "" {
		args = append(args, "--pkcs11-lib", req.Pkcs11Library)
	}
	if req.Pin != "" {
		args = append(args, "--pin", req.Pin)
	}
	if req.Alias != "" {
		args = append(args, "--alias", req.Alias)
	}
	if req.SignatureFormat != "" {
		args = append(args, "--signature-format", req.SignatureFormat)
	}
	return args
}

func buildValidateArgs(req ValidateRequest) []string {
	return []string{
		"--input-file", req.InputFile,
		"--certificate", req.CertificateFile,
		"--signature-file", req.SignatureFile,
	}
}

func executeLocal(operation string, opArgs []string) (string, error) {
	javaPath, err := jdk.GetJavaPath()
	if err != nil {
		return "", err
	}

	jarPath, err := getAssinadorJarPath()
	if err != nil {
		return "", err
	}

	args := append([]string{"-jar", jarPath, operation}, opArgs...)
	cmd := exec.Command(javaPath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("falha ao executar assinador: %v\n%s", err, string(out))
	}

	return string(out), nil
}

// ParseSignResponse decodifica a resposta JSON do assinador.
func ParseSignResponse(jsonResult string) (SignResponse, error) {
	var resp SignResponse
	if err := json.Unmarshal([]byte(jsonResult), &resp); err != nil {
		return resp, fmt.Errorf("resposta inválida do assinador: %w", err)
	}
	return resp, nil
}
