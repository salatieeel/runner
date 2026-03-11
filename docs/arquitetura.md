# Arquitetura do Sistema Runner

Resumo da arquitetura. Diagramas e detalhes em [problema/desing.md](../problema/desing.md) e [problema/contexto.md](../problema/contexto.md).

## Diagrama de contexto

O **Sistema Runner** interage com:

- **Usuário:** interage via CLI (comandos de assinatura e gerenciamento do simulador)
- **Dispositivo de Assinatura Digital:** sistema externo (token/smart card) para operações criptográficas via PKCS#11
- **Simulador do HubSaúde:** aplicação externa (simulador.jar) orquestrada pelo Runner

Diagramas em PlantUML: [problema/diagramas/contexto.puml](../problema/diagramas/contexto.puml).

## Diagrama de contêineres

Dentro do Sistema Runner:

| Contêiner | Descrição |
|-----------|-----------|
| **assinador CLI** | CLI multiplataforma; recebe comandos de assinatura do usuário e coordena |
| **assinador.jar** | Aplicação Java; valida parâmetros e simula assinaturas; pode interagir com dispositivo criptográfico |
| **simulador CLI** | CLI multiplataforma; gerencia ciclo de vida do simulador e expõe comandos ao usuário |

Comunicação: Usuário → CLIs; assinador CLI → assinador.jar (CLI direto ou HTTP); assinador.jar → Dispositivo (PKCS#11); simulador CLI → Simulador HubSaúde (HTTP/CLI).

Diagramas em PlantUML: [problema/diagramas/conteineres.puml](../problema/diagramas/conteineres.puml).

## Fluxo de criação de assinatura

1. Usuário executa comando para criar assinatura  
2. assinatura valida entrada do usuário  
3. assinatura invoca assinador.jar (diretamente ou via HTTP)  
4. assinador.jar valida parâmetros  
5. assinador.jar retorna assinatura simulada  
6. assinatura formata resultado e apresenta ao usuário  

## Fluxo de validação de assinatura

1. Usuário executa comando para validar assinatura  
2. assinatura valida entrada do usuário  
3. assinatura invoca assinador.jar (diretamente ou via HTTP)  
4. assinador.jar valida parâmetros  
5. assinador.jar retorna resultado simulado  
6. assinatura formata resultado e apresenta ao usuário  

## Tratamento de erros

Em qualquer ponto do fluxo, erros devem ser:

- Capturados apropriadamente  
- Propagados de forma estruturada  
- Apresentados ao usuário de forma clara  
- Incluir informação suficiente para correção  

## Modos de invocação do Assinador

- **Invocação direta (CLI):** cold start; adequado a execuções esporádicas ou scripts  
- **Invocação via HTTP (servidor):** warm start; menor latência e maior throughput para múltiplas requisições  
