# Requisitos do Sistema Runner

Resumo dos requisitos para o time. Fonte: [problema/contexto.md](../problema/contexto.md) e [problema/especificacao.md](../problema/especificacao.md).

## User stories

### US-01: Invocar Assinador via CLI

**Como** usuário do Sistema Runner  
**Quero** executar comandos de assinatura digital através da linha de comandos  
**Para que** eu possa invocar a aplicação **assinador.jar** (Assinador) sem conhecer os detalhes técnicos de configuração Java

**Critérios de aceitação:**
- O CLI deve aceitar comandos para criação e validação de assinatura
- O CLI deve invocar o Assinador com os parâmetros fornecidos
- O CLI deve suportar invocação direta do Assinador (modo local/CLI)
- O CLI deve suportar invocação do Assinador via HTTP (modo servidor)
- O CLI deve exibir o resultado da operação de forma legível

**Modos de invocação:** direta (cold start) ou via HTTP (warm start).

---

### US-02: Simular Assinatura Digital com Validação de Parâmetros

**Como** usuário do Sistema Runner  
**Quero** que o Assinador valide rigorosamente os parâmetros de entrada antes de simular uma operação de assinatura digital  
**Para que** eu receba feedback imediato sobre erros de parâmetros

**Critérios de aceitação:**
- O Assinador deve validar todos os parâmetros conforme especificações FHIR
- O Assinador deve simular criação de assinatura retornando resposta pré-construída quando parâmetros válidos
- O Assinador deve simular validação de assinatura retornando resultado pré-determinado
- O Assinador deve suportar interação com dispositivo criptográfico (token/smart card) via interface PKCS#11
- O Assinador deve retornar mensagens de erro claras quando parâmetros forem inválidos

---

### US-03: Gerenciar Ciclo de Vida do Simulador do HubSaúde

**Como** usuário do Sistema Runner  
**Quero** iniciar, parar e monitorar o Simulador do HubSaúde (**simulador.jar**) através do CLI  
**Para que** eu possa gerenciar o ciclo de vida do Simulador sem conhecer os comandos Java subjacentes

**Critérios de aceitação:**
- O CLI deve permitir iniciar o Simulador
- O CLI deve verificar se as portas necessárias para o Simulador estão disponíveis antes de iniciar
- O CLI deve permitir parar o Simulador
- O CLI deve exibir o status atual do Simulador (ou que não está em execução)
- O Simulador (simulador.jar) não faz parte do escopo de desenvolvimento
- O Simulador deve ser obtido dinamicamente pelo CLI (baixar versão mais recente do repositório da disciplina)
- O CLI não deve baixar o Simulador se a versão mais recente já estiver disponível localmente

---

### US-04: Provisionar JDK Automaticamente

**Como** usuário do Sistema Runner  
**Quero** que o sistema baixe e configure automaticamente o JDK necessário quando este não estiver disponível  
**Para que** eu possa utilizar o Assinador e o Simulador sem instalar ou configurar o Java manualmente

**Critérios de aceitação:**
- O sistema deve detectar se o JDK está presente na máquina (na versão exigida)
- O sistema deve baixar o JDK compatível quando ausente
- O sistema deve disponibilizar o JDK baixado para uso pelo Assinador e Simulador
- O download deve funcionar nas três plataformas (Windows, Linux, macOS)

---

### US-05: Disponibilizar binários multiplataforma

**Como** usuário do Sistema Runner  
**Quero** baixar uma versão pré-compilada do CLI para minha plataforma (Windows, Linux ou macOS)  
**Para que** eu possa utilizar o sistema imediatamente sem necessidade de compilação

**Critérios de aceitação:**
- Disponibilizar binário para Windows (amd64)
- Disponibilizar binário para Linux (amd64)
- Disponibilizar binário para macOS (amd64)
- Distribuir via GitHub Releases
- Incluir checksums SHA256 para verificação de integridade
- Utilizar versionamento semântico (SemVer)

---

## Parâmetros de entrada (FHIR)

- **Criar assinatura:** [Caso de Uso: Criar Assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-criar-assinatura.html)
- **Validar assinatura:** [Caso de Uso: Validar Assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-validar-assinatura.html)

## Assinatura de artefatos (Cosign/Sigstore)

Todos os artefatos distribuídos nas releases **devem** ser assinados com **Cosign** (Sigstore), usando identidade OIDC. Para cada artefato: publicar `<artefato>`, `<artefato>.sig` e `<artefato>.pem`. A assinatura deve ser feita automaticamente pelo pipeline de CI/CD.
