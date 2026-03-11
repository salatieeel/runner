# Referências

Links úteis para o desenvolvimento do Sistema Runner.

## Especificações FHIR (parâmetros de assinatura)

- [Caso de Uso: Criar Assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-criar-assinatura.html)
- [Caso de Uso: Validar Assinatura](https://fhir.saude.go.gov.br/r4/seguranca/caso-de-uso-validar-assinatura.html)

## Arquitetura (C4 e diagramas)

- [C4 Model](https://c4model.com/) — Diagrama de Contexto (nível 1) e de Contêineres (nível 2)
- [PlantUML](https://plantuml-documentation.readthedocs.io/en/latest/) — Diagramas no repositório: [problema/diagramas/](../problema/diagramas/)

## Assinatura e integridade de artefatos

- [Cosign](https://docs.sigstore.dev/cosign/overview/) — Assinatura de artefatos
- [Sigstore](https://www.sigstore.dev/) — Ecossistema (OIDC, transparency log)

## Versionamento e releases

- [Semantic Versioning (SemVer)](https://semver.org/)
- GitHub Releases — Distribuição dos binários

## Caminho de implementação (estrutura de código)

- [way.md](../problema/way.md) — Estrutura sugerida para **assinador** (Java/Maven) e **assinatura** (CLI, ex.: Go/Java)

## Boas práticas

- CLI: mensagens claras, tratamento de erros, help integrado
- BDD / critérios de aceitação: [Cucumber](https://cucumber.io/docs/)
- Requisitos: [SMART](https://thebaguide.com/blog/a-good-requirement-is-a-smart-requirement/), [INVEST](https://www.boost.co.nz/blog/2021/10/invest-criteria/)
