# Backlog inicial do Sistema Runner

Itens derivados das user stories (US-01 a US-05), para priorização e planejamento. Definition of Done será definida por item conforme o time avançar.

## US-01: Invocar Assinador via CLI

- [ ] Definir interface de comandos do CLI (criar, validar)
- [ ] Implementar parse de argumentos e validação sintática no CLI
- [ ] Implementar invocação do assinador.jar em modo direto (linha de comando)
- [ ] Implementar invocação do assinador.jar em modo HTTP (cliente)
- [ ] Formatar e exibir resultado das operações de forma legível
- [ ] Documentar uso dos comandos de assinatura

## US-02: Simular Assinatura Digital com Validação de Parâmetros

- [ ] Implementar validação de parâmetros conforme especificações FHIR (criar assinatura)
- [ ] Implementar validação de parâmetros conforme especificações FHIR (validar assinatura)
- [ ] Implementar simulação de criação de assinatura (resposta pré-construída)
- [ ] Implementar simulação de validação de assinatura (resultado pré-determinado)
- [ ] Suportar interface PKCS#11 para dispositivo criptográfico (token/smart card)
- [ ] Retornar mensagens de erro claras para parâmetros inválidos
- [ ] Testes unitários e de integração para validação e simulação

## US-03: Gerenciar Ciclo de Vida do Simulador do HubSaúde

- [ ] Implementar comando para iniciar o Simulador
- [ ] Verificar disponibilidade de portas antes de iniciar
- [ ] Implementar comando para parar o Simulador
- [ ] Implementar comando/consulta de status do Simulador
- [ ] Obter simulador.jar dinamicamente (download versão mais recente do repositório da disciplina)
- [ ] Evitar novo download se versão mais recente já estiver local
- [ ] Documentar comandos do simulador

## US-04: Provisionar JDK Automaticamente

- [ ] Detectar se JDK (versão exigida) está presente na máquina
- [ ] Implementar download do JDK compatível quando ausente
- [ ] Suportar download nas três plataformas (Windows, Linux, macOS)
- [ ] Configurar/disponibilizar JDK baixado para Assinador e Simulador
- [ ] Documentar comportamento de provisionamento

## US-05: Disponibilizar binários multiplataforma

- [ ] Definir pipeline de build para Windows (amd64)
- [ ] Definir pipeline de build para Linux (amd64)
- [ ] Definir pipeline de build para macOS (amd64)
- [ ] Configurar GitHub Releases e artefatos (SemVer)
- [ ] Incluir checksums SHA256 nos releases
- [ ] Integrar Cosign no CI/CD para assinatura de artefatos (.sig e .pem)
- [ ] Documentar processo de release e verificação de assinatura

## Cross-cutting

- [ ] Documentação: manual de usuário, guia de instalação, exemplos de uso
- [ ] Documentação técnica da integração CLI ↔ assinador.jar
- [ ] Testes de aceitação alinhados aos critérios das user stories
- [ ] Revisão e atualização de diagramas C4 se necessário
