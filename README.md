# Controle de Transação

[![Test](https://github.com/maiaaraujo5/controle-de-transacao/actions/workflows/tests.yaml/badge.svg)](https://github.com/maiaaraujo5/controle-de-transacao/actions/workflows/tests.yaml)

Esta aplicação tem como objetivo simular operações em uma conta

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

### make docker-compose-up

Este comando é responsável por subir o banco de dados e a aplicação. Neste comando fazemos o build da imagem docker da
aplicação.

## Contato

- [Linkedin](https://www.linkedin.com/in/lucasmaiamelo/)
- Email: maia.araujo7@gmail.com