package jdk

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetJavaPath retorna o caminho do executável Java 21+ instalado no sistema.
func GetJavaPath() (string, error) {
	path, err := exec.LookPath("java")
	if err != nil {
		return "", fmt.Errorf("Java não encontrado no PATH. Instale JDK 21 ou superior e tente novamente")
	}

	out, err := exec.Command(path, "-version").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("falha ao verificar versão do Java: %w", err)
	}

	versionText := string(out)
	if !strings.Contains(versionText, "21") &&
		!strings.Contains(versionText, "22") &&
		!strings.Contains(versionText, "23") &&
		!strings.Contains(versionText, "24") &&
		!strings.Contains(versionText, "25") {
		return "", fmt.Errorf("Java 21+ é obrigatório. Versão detectada:\n%s", versionText)
	}

	return path, nil
}
