# 🏦 GoGiota API (ou insira seu nome escolhido aqui)

> ⚠️ **Aviso:** Este projeto é um exercício prático do curso de Golang da **Rocketseat**. É uma API bancária com proteção CSRF, porque até em projeto de estudo a gente precisa proteger nossos zero reais de hackers mal-intencionados.

## 🚀 Sobre o Projeto

Neste projeto, desenvolvemos um servidor HTTP em **Golang** para gerenciar operações bancárias simples de Pessoas Físicas e Jurídicas. O foco principal é entender como criar e estruturar rotas HTTP, conectar com um banco de dados **PostgreSQL** usando migrations e, o mais importante, blindar a aplicação contra ataques implementando um middleware de validação de tokens **CSRF** em operações que alteram o estado dos dados.

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Banco de Dados:** PostgreSQL
- **Conceitos abordados:** API REST, Middlewares, CSRF Protection, Migrations, Regras de Negócio.

## 🗄️ Estrutura do Banco de Dados

O banco utiliza duas tabelas principais para gerenciar os calo... digo, os clientes:

### Tabela: `pessoa_fisica`
| Campo | Tipo | Descrição |
| :--- | :--- | :--- |
| `id` | `PK` | Identificador único (autoincrementável) |
| `renda_mensal` | `Decimal` | Renda mensal do indivíduo |
| `idade` | `Inteiro` | Idade em anos |
| `nome_completo`| `Varchar(255)` | Nome da pessoa |
| `celular` | `Varchar(20)` | Telefone de contato |
| `email` | `Varchar(255)` | E-mail de contato |
| `categoria` | `Varchar(50)` | Categoria do cliente |
| `saldo` | `Decimal` | Saldo disponível para gastar (ou chorar) |

### Tabela: `pessoa_juridica`
| Campo | Tipo | Descrição |
| :--- | :--- | :--- |
| `id` | `PK` | Identificador único (autoincrementável) |
| `faturamento` | `Decimal` | Faturamento anual da empresa |
| `idade` | `Inteiro` | Tempo de existência da empresa (anos) |
| `nome_fantasia`| `Varchar(255)` | Nome comercial |
| `celular` | `Varchar(20)` | Telefone de contato |
| `email_corporativo`| `Varchar(255)`| E-mail de contato |
| `categoria` | `Varchar(50)` | Categoria da empresa |
| `saldo` | `Decimal` | Saldo disponível em caixa |

## 🛣️ Rotas da API

A lógica de negócios é livre, mas a API respeita os seguintes endpoints (todos retornando JSON):

- `POST /conta` - Cria uma nova conta (PF ou PJ).
- `GET /conta/{id}/saldo` - Consulta o saldo (preparando o psicológico).
- `POST /conta/{id}/deposito` - Deposita dinheiro na conta.
- `POST /conta/{id}/saque` - Saca dinheiro da conta.
- `POST /conta/transferencia` - Transfere dinheiro entre contas.
- `DELETE /conta/{id}` - Fecha a conta (para fugir das dívidas).

## 🛡️ Proteção CSRF

Para garantir que ninguém faça transferências fantasmas no seu nome, o projeto conta com um **Middleware CSRF**. 
Todas as rotas que modificam dados (POST, DELETE) exigem a presença e validação de um token CSRF no header da requisição. Se o token não for válido, a API barra a transação na hora.

## 💻 Como rodar este projeto

1. Clone este repositório:
   ```bash
   git clone [https://github.com/luanrem/Go-Giota-API.git](https://github.com/luanrem/Go-Giota-API.git)