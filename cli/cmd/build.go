package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runBuild(args []string) error {
	target := "all"
	if len(args) > 0 {
		target = args[0]
	}

	switch target {
	case "all", "assinador":
		if err := buildAssinadorMaven(); err != nil {
			return err
		}
		if target == "assinador" {
			return nil
		}
		fallthrough
	case "cli":
		return buildCLI()
	default:
		return fmt.Errorf("alvo de build desconhecido: %s (use: assinador, cli ou all)", target)
	}
}

func projectRoot() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(filepath.Join(filepath.Dir(ex), ".."))
}

func buildAssinadorMaven() error {
	root, err := projectRoot()
	if err != nil {
		return err
	}

	fmt.Println("Compilando assinador com Maven...")
	cmd := exec.Command("mvn", "-B", "-q", "package", "-f", filepath.Join(root, "assinador", "pom.xml"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = root
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao compilar assinador (requer Maven e JDK 21+): %w", err)
	}

	jar := filepath.Join(root, "assinador", "target", "assinador.jar")
	fmt.Printf("✅ Artefato gerado em %s\n", jar)
	return nil
}

func buildCLI() error {
	root, err := projectRoot()
	if err != nil {
		return err
	}

	cliDir := filepath.Join(root, "cli")
	outputName := "assinatura"
	if os.Getenv("GOOS") == "windows" {
		outputName = "assinatura.exe"
	}

	fmt.Println("Compilando CLI em Go...")
	cmd := exec.Command("go", "build", "-o", outputName, ".")
	cmd.Dir = cliDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao compilar CLI: %w", err)
	}

	fmt.Printf("✅ CLI compilado em %s\n", filepath.Join(cliDir, outputName))
	return nil
}
