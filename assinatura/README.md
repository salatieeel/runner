# assinatura (CLI)

Módulo da aplicação **CLI multiplataforma** do Sistema Runner.

## Objetivo

Fornecer um binário executável (Windows, Linux, macOS) que:

- Recebe comandos do usuário via linha de comandos
- Valida sintaxe básica dos parâmetros de entrada
- Invoca o **assinador.jar** (modo direto ou via HTTP) para criar/validar assinaturas
- Gerencia o ciclo de vida do **Simulador do HubSaúde** (simulador.jar): iniciar, parar, status
- Provisiona JDK automaticamente quando não estiver disponível
- Apresenta resultados e mensagens de erro de forma legível

## Responsabilidades (a implementar)

- Comandos de assinatura: criar e validar (repasse ao assinador.jar)
- Comandos do simulador: iniciar, parar, status; verificação de portas; download do simulador.jar quando necessário
- Detecção e download de JDK nas três plataformas
- Formatação de saída e tratamento de erros para o usuário

## Estrutura sugerida

A organização de diretórios e componentes recomendada está em **[problema/way.md](../problema/way.md)** (seção 3), com exemplos em Go e em Java, incluindo:

- `cmd` / ponto de entrada
- Pacotes ou pastas para: CLI (comandos e parser), integração com assinador (CLI e HTTP), gerenciamento do simulador, JDK (detector, download, configurador), UI (saída e erros)

A linguagem do CLI não é fixada no enunciado (Go, Rust, Java com GraalVM, etc.); a escolha fica a critério do time, desde que se gerem binários para Windows, Linux e macOS.

## Documentação do projeto

- [docs/](docs/) — Convenção de comandos, guia de instalação, exemplos de uso específicos do CLI.
