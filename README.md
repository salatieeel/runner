# Assistente de Assinatura de Documentos

Sistema em terminal para assinar documentos de forma simulada.

A ideia é simples: a pessoa abre o sistema, escolhe uma opção pelo número do menu e segue as instruções na tela.

## Pré-requisitos

- **Node.js** 18 ou superior (para os scripts `npm`)
- **Go** 1.21 ou superior
- **Java JDK** 21 ou superior (`java` no PATH)
- **Maven** 3.9 ou superior (`mvn` no PATH)

## Parte 1: comandos para rodar o sistema

### 1. Clonar o projeto

Abra o terminal e execute o comando abaixo para baixar o projeto para sua máquina:

```bash
git clone git@github.com:salatieeel/runner.git
```

## 2. Entrar na pasta do projeto

Após concluir o download, acesse a pasta do projeto:

```bash
cd runner
```

> 💡 **O que esses comandos fazem?**
>
> - `git clone`: baixa uma cópia do projeto do GitHub para o seu computador.
> - `cd runner`: entra na pasta do projeto para que você possa executar os próximos comandos.

### 2. Rodar o sistema

Use este comando:

```bash
npm start
```

Esse comando compila o projeto e abre o menu principal.

### 3. Se quiser apenas compilar

Use:

```bash
npm run build
```

### 4. Se quiser limpar arquivos gerados

Use:

```bash

```

Depois de limpar, rode novamente:

```bash
npm start
```

## Parte 2: como usar o terminal

Quando o sistema abrir, ele mostra um menu parecido com este:

```text
1 - Assinar um documento agora
2 - Ver onde os arquivos assinados foram salvos
3 - Abrir a pasta dos arquivos assinados
9 - Opções avançadas
0 - Sair
```

Para escolher uma opção, digite o número e pressione `ENTER`.

### Botão 1: assinar um documento

Use o botão `1` para assinar um documento.

O sistema vai pedir o caminho completo do arquivo. Exemplo:

```text
/home/usuario/Documentos/contrato.pdf
```

Depois de informar o caminho, o sistema envia o documento para o assinador Java e grava três arquivos simulados na pasta `DocumentosAssinados` (por data).

O documento original não é alterado.

### Botão 2: ver onde os arquivos ficam

Use o botão `2` para ver a pasta padrão dos documentos assinados.

O sistema mostra uma pasta parecida com:

```text
/home/usuario/DocumentosAssinados
```

### Botão 3: abrir a pasta dos documentos

Use o botão `3` para abrir a pasta `DocumentosAssinados`.

Se a pasta ainda não existir, o sistema cria a pasta.

Se o computador conseguir abrir o gerenciador de arquivos, a pasta aparece na tela. Se não conseguir, o próprio terminal mostra o conteúdo da pasta.

### Botão 9: opções avançadas

Use o botão `9` apenas se você estiver desenvolvendo ou testando o projeto.

Ele mostra comandos técnicos, como:

```bash
npm run build
npm run build:assinador
npm run build:cli
```

Para uso normal, não precisa entrar nessa opção.

### Botão 0: sair

Use o botão `0` para fechar o sistema.

## Observação importante

Este projeto faz assinatura simulada. Ele é indicado para estudo, demonstração e testes.

Ele não deve ser usado como assinatura digital oficial sem certificado válido e validação jurídica.
