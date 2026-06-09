package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kyriosdata/assinatura/internal/assinador"
)

func runSign(args []string) error {
	fs := flag.NewFlagSet("sign", flag.ContinueOnError)
	inputFile := fs.String("input-file", "", "Arquivo a ser assinado")
	certificate := fs.String("certificate", "", "Certificado digital")
	pkcs11Lib := fs.String("pkcs11-lib", "", "Biblioteca PKCS#11")
	pin := fs.String("pin", "", "PIN do dispositivo")
	alias := fs.String("alias", "", "Alias do certificado")
	sigFormat := fs.String("signature-format", "CMS", "Formato da assinatura")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if *inputFile == "" {
		return fmt.Errorf("--input-file é obrigatório")
	}

	certPath := *certificate
	if certPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		certPath, err = assinador.EnsureDemoCertificate(
			filepath.Join(home, "DocumentosAssinados"),
		)
		if err != nil {
			return err
		}
	}

	req := assinador.SignRequest{
		InputFile:       *inputFile,
		CertificateFile: certPath,
		Pkcs11Library:   *pkcs11Lib,
		Pin:             *pin,
		Alias:           *alias,
		SignatureFormat: *sigFormat,
	}

	result, err := assinador.ExecuteSign(req)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func runValidate(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	inputFile := fs.String("input-file", "", "Arquivo assinado")
	signatureFile := fs.String("signature-file", "", "Arquivo de assinatura")
	certificate := fs.String("certificate", "", "Certificado digital")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if *inputFile == "" || *signatureFile == "" || *certificate == "" {
		return fmt.Errorf("--input-file, --signature-file e --certificate são obrigatórios")
	}

	req := assinador.ValidateRequest{
		InputFile:       *inputFile,
		CertificateFile: *certificate,
		SignatureFile:   *signatureFile,
	}

	result, err := assinador.ExecuteValidate(req)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
