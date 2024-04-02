# Controle de Transação

[![Test](https://github.com/maiaaraujo5/controle-de-transacao/actions/workflows/tests.yaml/badge.svg)](https://github.com/maiaaraujo5/controle-de-transacao/actions/workflows/tests.yaml)

Esta aplicação visa simular operações em uma conta.

## Tecnologias Utilizadas

- **Go 1.22.1**
- **Banco de Dados Postgres**
- **Gihub Actions**
- **Docker**

## Bibliotecas utilizadas

- [Echo](https://github.com/labstack/echo)
- [Bun ORM](https://github.com/uptrace/bun)
- [Validator](https://github.com/go-playground/validator)
- [Viper](https://github.com/spf13/viper)
- [FX](https://github.com/uber-go/fx)
- [Testify](https://github.com/stretchr/testify)
- [Monkey](https://github.com/bouk/monkey)
- [Swag](https://github.com/swaggo/swag)
- [Echo-Swagger](https://github.com/swaggo/echo-swagger)
- [Testcontainers-Go](https://github.com/testcontainers/testcontainers-go)


## Documentação
- Quando a aplicação estiver rodando
  <a href="http://localhost:8080/swagger/index.html" target="_blank">Clique aqui</a>
  para acessar a documentação

- [Clique Aqui](Controle%20de%20transa%C3%A7%C3%A3o%20-%20Lucas%20Maia.postman_collection.json) para pegar a collection do postman

## Checklist de Funcionalidades

- [x] Criar contas
- [x] Recuperar contas
- [x] Criar transações
- [x] Salvar operações de **COMPRA A VISTA**, **COMPRA PARCELADA** e **SAQUE** como negativo no banco de dados
- [x] Verificar a existência da conta antes de salvar a transação
- [x] Erros personalizados
- [x] Validação do request body

## Como rodar a aplicação

### *Requisitos necessários*
Os requisitos necessários para rodar a aplicação são:

* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/install/)
* [Go](https://go.dev/dl/)

**Clique nos links acima para instalar as dependências caso ainda não tenha no seu computador**

### make docker-compose-up

Este comando é responsável por subir o banco de dados e a aplicação em container docker deixando a aplicação pronta para receber requests. Neste comando fazemos o build da imagem docker da
aplicação.

### make run-go
Este comando é responsável por rodar a aplicação utilizando o go e subindo o banco de dados deixando a aplicação pronta para receber requests


## Comandos no Makefile

### make test
Este comando é responsável por rodar todos os testes unitários da aplicação.

### make e2e_test
Este comando é responsável por rodar os testes end to end da aplicação.

### make lint
Este comando é responsável por rodar os linters.

### make docker-compose-up
Este comando é responsável por subir o banco de dados e a aplicação em container docker. Neste comando fazemos o build da imagem docker da
aplicação.

### make docker-compose-up-dependencies
Este comando é responsável por subir somente o banco de dados presente no docker-compose.yaml

### make docker-compose-down
Este comando é responsável por parar os containers que estão rodando e fazem parte do docker-compose.yaml

### make docker-build
Este comando é responsável por criar imagem docker da aplicação

### make generate-swagger
Este comando é responsável por gerar a documentação swagger da aplicação

### make run-go
Este comando é responsável por rodar a aplicação utilizando o go e subindo o banco de dados deixando a aplicação pronta para receber requests

## Contato

- [Linkedin](https://www.linkedin.com/in/lucasmaiamelo/)
- Email: maia.araujo7@gmail.com