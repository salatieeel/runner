# assinador (assinador.jar)

Módulo que produz a aplicação Java **assinador.jar**, parte do Sistema Runner.

## Objetivo

Gerar um JAR executável que:

- Valida rigorosamente os parâmetros de entrada conforme especificações FHIR
- Simula a **criação** de assinatura digital (retorna resposta pré-construída quando parâmetros são válidos)
- Simula a **validação** de assinatura digital (retorna resultado pré-determinado)
- Suporta interação com dispositivo criptográfico (token/smart card) via interface PKCS#11
- Retorna mensagens de erro claras quando parâmetros forem inválidos

## Modos de execução

- **Modo local (CLI):** invocado diretamente via linha de comando (ex.: `java -jar assinador.jar ...`); cada execução é cold start.
- **Modo servidor (HTTP):** sobe um servidor HTTP e permanece em execução; o CLI assinatura envia requisições HTTP (warm start).

## Responsabilidades (a implementar)

- Validação de parâmetros (criar e validar assinatura) conforme FHIR
- Simulação de criação de assinatura
- Simulação de validação de assinatura
- Exposição via CLI e via HTTP (endpoints para criar/validar)
- Tratamento de erros e respostas estruturadas

## Estrutura sugerida

A organização de pacotes e diretórios recomendada está em **[problema/way.md](../problema/way.md)** (seção 2), incluindo:

- Ponto de entrada (`Main`), decisão modo CLI vs HTTP
- Pacotes para CLI, HTTP, core (regras de negócio/simulação), model (request/response/erro), validação FHIR
- Testes unitários para serviços de assinatura e validação

A implementação será feita nesta pasta seguindo essa estrutura (Maven ou Gradle, conforme escolha do time).

## Documentação do projeto

- [docs/](docs/) — Decisões e documentação específica do assinador (API HTTP, formato de erros, etc.)
