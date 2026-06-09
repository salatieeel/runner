package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/kyriosdata/assinatura/internal/assinador"
)

func runInteractiveMenu() {
	if !isInteractiveTerminal() {
		fmt.Fprintln(os.Stderr, "Erro: é necessário um terminal interativo para usar o menu.")
		fmt.Fprintln(os.Stderr, "Abra o terminal do sistema (não use redirecionamento) e execute:")
		fmt.Fprintln(os.Stderr, "  cd runner && npm start")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()

		fmt.Print("Digite sua opção: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			signDocumentInteractive(reader)
		case "2":
			showOutputPath()
		case "3":
			openOutputFolder()
		case "9":
			showAdvancedOptions()
		case "0":
			fmt.Println("\n👋 Até logo!")
			os.Exit(0)
		default:
			fmt.Println("\n❌ Opção inválida. Tente novamente.")
			fmt.Print("Pressione ENTER para continuar...")
			reader.ReadString('\n')
		}
	}
}

func displayMenu() {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║  Assistente de Assinatura de Documentos   ║")
	fmt.Println("╚════════════════════════════════════════════╝\n")
	fmt.Println("Opções disponíveis:\n")
	fmt.Println("  1 – Assinar um documento agora")
	fmt.Println("  2 – Ver onde os arquivos assinados foram salvos")
	fmt.Println("  3 – Abrir a pasta dos arquivos assinados")
	fmt.Println("  9 – Opções avançadas (técnicas)")
	fmt.Println("  0 – Sair\n")
}

func signDocumentInteractive(reader *bufio.Reader) {
	fmt.Print("Cole o caminho completo do documento: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	if filePath == "" {
		fmt.Println("❌ Caminho vazio. Operação cancelada.")
		pause(reader)
		return
	}

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("❌ Arquivo não encontrado: %s\n", filePath)
		pause(reader)
		return
	}

	fmt.Println("\n⏳ Assinando documento...")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Erro ao obter home dir: %v\n", err)
		pause(reader)
		return
	}

	fileName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	today := time.Now().Format("2006-01-02")
	outputDir := filepath.Join(homeDir, "DocumentosAssinados", today)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Erro ao criar diretório: %v\n", err)
		pause(reader)
		return
	}

	certPath, err := assinador.EnsureDemoCertificate(outputDir)
	if err != nil {
		fmt.Printf("❌ Erro ao preparar certificado simulado: %v\n", err)
		pause(reader)
		return
	}

	signReq := assinador.SignRequest{
		InputFile:       filePath,
		CertificateFile: certPath,
	}

	result, err := assinador.ExecuteSign(signReq)
	if err != nil {
		fmt.Printf("❌ Erro ao assinar: %v\n", err)
		pause(reader)
		return
	}

	if err := assinador.WriteSignOutputs(outputDir, fileName, filePath, result); err != nil {
		fmt.Printf("❌ Erro ao salvar arquivos: %v\n", err)
		pause(reader)
		return
	}

	fmt.Println("✅ Documento assinado com sucesso!\n")
	fmt.Printf("📁 Arquivos salvos em: %s\n\n", outputDir)
	fmt.Printf("  • %s.assinado.json\n", fileName)
	fmt.Printf("  • %s-transcricao-assinada.txt\n", fileName)
	fmt.Printf("  • %s-comprovante-assinatura.txt\n\n", fileName)
	fmt.Println(result)
	pause(reader)
}

func showOutputPath() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}

	outputPath := filepath.Join(homeDir, "DocumentosAssinados")
	fmt.Printf("\n📂 Pasta padrão de saída:\n%s\n\n", outputPath)
}

func openOutputFolder() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}

	outputPath := filepath.Join(homeDir, "DocumentosAssinados")

	if err := os.MkdirAll(outputPath, 0755); err != nil {
		fmt.Printf("❌ Erro ao criar diretório: %v\n", err)
		return
	}

	fmt.Printf("✅ Abrindo pasta: %s\n\n", outputPath)

	go func() {
		var commands []string
		switch runtime.GOOS {
		case "linux":
			commands = []string{"xdg-open", "nautilus", "thunar", "dolphin"}
		case "darwin":
			commands = []string{"open"}
		case "windows":
			commands = []string{"explorer"}
		default:
			commands = []string{"xdg-open"}
		}
		for _, cmdName := range commands {
			cmd := exec.Command(cmdName, outputPath)
			if err := cmd.Start(); err == nil {
				return
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)

	entries, err := os.ReadDir(outputPath)
	if err != nil {
		fmt.Printf("❌ Erro ao ler pasta: %v\n", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("📁 Pasta vazia. Nenhum documento assinado ainda.")
	} else {
		fmt.Println("📂 Arquivos assinados:\n")
		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("  📁 %s\n", entry.Name())
				subEntries, _ := os.ReadDir(filepath.Join(outputPath, entry.Name()))
				for _, subEntry := range subEntries {
					if !subEntry.IsDir() {
						info, _ := subEntry.Info()
						fmt.Printf("      • %s (%d bytes)\n", subEntry.Name(), info.Size())
					}
				}
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nPressione ENTER para continuar...")
	reader.ReadString('\n')
}

func showAdvancedOptions() {
	clearScreen()
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║        Opções Avançadas (Técnicas)        ║")
	fmt.Println("╚════════════════════════════════════════════╝\n")
	fmt.Println("Pré-requisitos: Go 1.21+, Java JDK 21+, Maven 3.9+\n")
	fmt.Println("Comandos disponíveis:\n")
	fmt.Println("  npm run build              – Compilar CLI e assinador")
	fmt.Println("  npm run build:assinador     – Compilar serviço Java (Maven)")
	fmt.Println("  npm run build:cli           – Compilar CLI Go")
	fmt.Println("\nUso direto do assinador:\n")
	fmt.Println("  ./cli/assinatura sign --input-file arquivo.pdf --certificate cert.pem")
	fmt.Println("  java -jar assinador/target/assinador.jar sign --input-file arquivo --certificate cert.pem\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pressione ENTER para voltar...")
	reader.ReadString('\n')
}

func pause(reader *bufio.Reader) {
	fmt.Print("Pressione ENTER para continuar...")
	reader.ReadString('\n')
}

func isInteractiveTerminal() bool {
	return isCharDevice(os.Stdin) && isCharDevice(os.Stdout)
}

func isCharDevice(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}

func clearScreen() {
	// Evita \033[2J: em terminais integrados (Cursor/VS Code) pode apagar a saída
	// e dar a impressão de que o menu não abriu.
	fmt.Println()
}
