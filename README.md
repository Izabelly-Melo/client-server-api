# Client-Server API - Cotação do Dólar

Desafio da pós GoExpert (Full Cycle) que junta HTTP, Context, banco de dados e manipulação de arquivos em Go. São dois programas separados: um server que busca a cotação do dólar e salva no banco, e um client que consulta esse server e grava o resultado em um arquivo.

## Como funciona

O `server` recebe uma requisição em `/cotacao`, chama a API da AwesomeAPI pra pegar a cotação atual do USD-BRL, salva essa cotação num banco SQLite e devolve o valor pro client em JSON.

O `client` faz uma requisição pro server, pega o valor recebido e escreve num arquivo `cotacao.txt` no formato:

```
Dólar: {valor}
```

Os dois se comunicam com timeouts bem curtos de propósito (200ms pra API externa, 10ms pro banco, 300ms pro client aguardar o server), então é esperado ver algum log de timeout de vez em quando — principalmente na hora de salvar no banco. Mais detalhes sobre isso lá embaixo.

## Pré-requisitos

- Go 1.25 ou superior instalado

## Rodando o projeto

Primeiro, baixe as dependências (GORM + driver SQLite):

```bash
go mod tidy
```

Depois, abra **dois terminais**.

No primeiro, suba o server:

```bash
go run cmd/server/server.go
```

Ele vai ficar escutando na porta `8080`. Deixe esse terminal aberto.

No segundo terminal, rode o client:

```bash
go run cmd/client/client.go
```

Se tudo der certo, vai aparecer um arquivo `cotacao.txt` na raiz do projeto com a cotação do dia, tipo:

```
Dólar: 5.4321
```

## Sobre os timeouts

Os prazos usados aqui são curtos de propósito, pra forçar o cenário de estouro de tempo que o desafio pede:

- **200ms** pra chamar a API externa de câmbio.
- **10ms** pra gravar a cotação no banco SQLite (esse é bem apertado mesmo).
- **300ms** pro client esperar a resposta do server.

Por causa do timeout de 10ms no banco, é normal o console do **server** mostrar um log de timeout de vez em quando ao tentar salvar a cotação — não é bug, é o comportamento esperado pelo desafio.

Se o timeout do client (300ms) estourar, o console do client vai logar isso e o `cotacao.txt` não vai ser gerado nessa execução. Nesse caso, é só rodar o client de novo.

## Estrutura do projeto

```
.
├── cmd/
│   ├── server/
│   │   └── server.go
│   ├── client/
│   │   └── client.go
│   ├── database/
│   │   └── database.go
│   └── models/
│       └── currency.go
├── go.mod
├── go.sum
└── README.md
```