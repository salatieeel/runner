# Entregáveis da disciplina

Lista de entregáveis a serem confeccionados e disponibilizados no repositório (GitHub). Fonte: [problema/especificacao.md](../problema/especificacao.md) (seção 7).

## 1. Código-fonte da aplicação assinatura

- Implementação completa
- Compatível com Windows, Linux e macOS
- Código bem documentado

## 2. Código-fonte da aplicação assinador.jar

- Implementação em Java
- Validação completa de parâmetros
- Simulação das operações

## 3. Testes

- Testes unitários
- Testes de integração
- Casos de teste para cenários de erro
- Testes de aceitação baseados nos critérios definidos

## 4. Documentação

- Manual de usuário para assinatura
- Documentação técnica da integração
- Exemplos de uso
- Guia de instalação

## 5. Especificação

- Contexto e escopo definidos
- Diagramas C4
- Requisitos documentados  
(Referência: documentos em **problema/**)

## 6. Artefatos executáveis

- Binários pré-compilados para as três plataformas:
  - `assinatura-<versão>-windows-amd64.exe` (Windows)
  - `assinatura-<versão>-linux-amd64.AppImage` (Linux)
  - `assinatura-<versão>-macos-amd64.dmg` (macOS)
- Distribuídos via **GitHub Releases**
- Cada release com assinatura dos artefatos (Cosign: `.sig` e `.pem` por artefato)
- Versionamento semântico (SemVer)
