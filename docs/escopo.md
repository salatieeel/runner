# Escopo do Sistema Runner

Fonte: [problema/especificacao.md](../problema/especificacao.md) (seções 4.1 e 4.2).

## O que ESTÁ no escopo

- Desenvolvimento da aplicação **assinatura** (CLI multiplataforma)
- Desenvolvimento da aplicação **assinador.jar** (Java)
- Integração entre as duas aplicações
- Validação rigorosa de parâmetros pelo assinador.jar
- Simulação de criação de assinatura (assinador.jar)
- Simulação de validação de assinatura (assinador.jar)
- Tratamento de erros dos parâmetros e exceções (assinador.jar)
- Testes
- Documentação de uso

## O que NÃO ESTÁ no escopo

- Implementação real de assinatura digital criptográfica
- Implementação real de validação de assinatura digital criptográfica
- Integração com autoridades certificadoras
- Armazenamento persistente de assinaturas
- Interface gráfica (GUI)
- Autenticação de usuários
- Geração de certificados digitais
