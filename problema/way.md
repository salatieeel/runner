## Caminho de Implementação do Sistema Runner

Este documento descreve, de forma prática, **como estruturar o código** do Sistema Runner, com foco em:

- A aplicação **Java** `assinador.jar`
- A aplicação **CLI multiplataforma** `assinatura`

Não há obrigatoriedade de uso de frameworks pesados (como Spring). Todo o escopo pode ser atendido com:

- **Java SE** para o `assinador.jar`
- Uma linguagem à sua escolha para o `assinatura` (inclusive Java), desde que gere binários para Windows, Linux e macOS

---

## 1. O que precisa ser feito em Java

Pelo documento de contexto (`contexto.md`), a aplicação `assinador.jar` é:

- Uma **aplicação Java** responsável por:
  - Validar rigorosamente os parâmetros de entrada (conforme FHIR)
  - Simular criação de assinatura
  - Simular validação de assinatura
  - Expor dois modos de execução:
    - **Modo local (CLI)**: chamada direta via linha de comando
    - **Modo servidor (HTTP)**: servidor simples que recebe requisições HTTP

Portanto, **sim**, você necessariamente terá um **projeto em Java** para gerar o arquivo `assinador.jar`.

Não é obrigatório usar frameworks; você pode implementar tudo com **Java puro**, usando:

- `public static void main(String[] args)` para o modo CLI
- `com.sun.net.httpserver.HttpServer` ou similar para o modo HTTP (ou um microframework leve, se desejar)

---

## 2. Estrutura sugerida para o projeto `assinador` (Java)

Abaixo, uma sugestão de estrutura de pastas **bem concreta** usando Maven (poderia ser Gradle com estrutura equivalente).

```text
assinador/
  pom.xml
  src/
    main/
      java/
        br/
          ufg/
            runner/
              assinador/
                Main.java
                cli/
                  AssinadorCli.java
                http/
                  AssinadorHttpServer.java
                core/
                  AssinaturaService.java
                  ValidacaoService.java
                model/
                  CriarAssinaturaRequest.java
                  CriarAssinaturaResponse.java
                  ValidarAssinaturaRequest.java
                  ValidarAssinaturaResponse.java
                  ErroValidacao.java
                fhir/
                  FhirParameterValidator.java
    test/
      java/
        br/
          ufg/
            runner/
              assinador/
                AssinaturaServiceTest.java
                ValidacaoServiceTest.java
```

**Ideia dessa estrutura:**

- **`Main.java`**: ponto de entrada. Lê os argumentos e decide se roda em modo CLI ou inicia o servidor HTTP.
- **`cli/AssinadorCli.java`**: lógica de parse dos parâmetros de linha de comando (criar, validar, etc.).
- **`http/AssinadorHttpServer.java`**: servidor HTTP simples, expondo endpoints para criar/validar assinatura.
- **`core/`**: regras de negócio e simulação de assinatura.
- **`model/`**: classes que representam requisições/respostas e erros.
- **`fhir/FhirParameterValidator.java`**: validação dos parâmetros conforme as especificações FHIR.

Tecnicamente, tudo isso poderia estar em um único arquivo `.java`, mas **é fortemente recomendável** separar em pacotes, para:

- Facilitar testes unitários
- Manter o código legível
- Evoluir a solução sem virar um “arquivo gigante impossível de mexer”

---

## 3. Estrutura sugerida para o projeto `assinatura` (CLI)

O documento especifica que `assinatura` é:

- Uma **aplicação CLI multiplataforma**
- Responsável por:
  - Receber comandos do usuário
  - Validar sintaxe básica dos parâmetros
  - Invocar o `assinador.jar` (via CLI ou HTTP)
  - Apresentar resultados de forma legível

A linguagem dessa aplicação não é fixada no enunciado. Você pode:

- Implementar também em **Java** (e depois gerar binários nativos com ferramentas como GraalVM), ou
- Usar outra linguagem com boa história para CLIs (por exemplo, **Go** ou **Rust**), desde que:
  - Haja binários para Windows, Linux e macOS
  - O CLI saiba encontrar/baixar/executar o `assinador.jar`

Abaixo, uma estrutura de exemplo em **Go** (apenas como referência de organização; adapte para a linguagem escolhida).

```text
assinatura/
  go.mod
  cmd/
    assinatura/
      main.go
  internal/
    cli/
      commands.go        # definição de comandos: criar, validar, simulador start/stop/status
      parser.go          # parse de flags/argumentos
    runner/
      assinador_cli.go   # chamada ao assinador.jar via linha de comando
      assinador_http.go  # chamada ao assinador.jar via HTTP (se usado modo servidor)
      jdk_manager.go     # lógica de detecção/download/configuração do JDK
      simulador.go       # comandos para gerenciar o simulador.jar
    ui/
      output.go          # formatação de saída e mensagens de erro claras
```

Se você preferir fazer o CLI em Java, uma estrutura equivalente usando Maven poderia ser:

```text
assinatura/
  pom.xml
  src/
    main/
      java/
        br/
          ufg/
            runner/
              assinatura/
                Main.java
                cli/
                  CommandParser.java
                  CreateCommand.java
                  ValidateCommand.java
                  SimuladorStartCommand.java
                  SimuladorStopCommand.java
                  SimuladorStatusCommand.java
                integration/
                  AssinadorCliClient.java
                  AssinadorHttpClient.java
                  SimuladorManager.java
                jdk/
                  JdkDetector.java
                  JdkDownloader.java
                  JdkConfigurator.java
                ui/
                  ConsolePrinter.java
                  ErrorFormatter.java
    test/
      java/
        br/
          ufg/
            runner/
              assinatura/
                CommandParserTest.java
                JdkDetectorTest.java
```

---

## 4. Frameworks: são necessários?

Não. Para atender ao enunciado, você pode:

- Usar **Java SE puro** no `assinador` (eventualmente com um microframework HTTP, se quiser simplicidade extra).
- Usar recursos básicos da linguagem escolhida no `assinatura` (biblioteca padrão para parse de argumentos, execução de processos, HTTP e I/O).

Frameworks só seriam úteis para:

- Facilitar a criação do servidor HTTP (no `assinador`)
- Melhorar a ergonomia do CLI (no `assinatura`, com bibliotecas específicas de CLI)

Mas **não são uma exigência** do trabalho; foque primeiro em:

- Atender as user stories (`US-01` a `US-05`)
- Validar parâmetros conforme FHIR
- Simular corretamente criação/validação de assinatura
- Prover boa experiência de uso em linha de comando

Com essa estrutura de pastas, você tem um **“caminho de implementação”** claro para organizar o código sem precisar de frameworks complexos.

