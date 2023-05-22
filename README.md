# Verify My Age - Backend test

Simples CRUD de usuários com o intuito de demonstrar o conhecimento.

## Info

Este projeto segue como base os princípios da Clean Arch de Robert C Martin,
também conhecido como Uncle Bob. Buscando um alto grau de desacoplamento entre os componentes, facilitar a manutenção e 
blindar regras de negócio, para não sofrerem interferência de atores secundários como web frameworks,
ORMs ou banco de dados.

## Stack utilizada
   * Repositório de dados
     * GORM
     * Driver MySQL
   * API
     * Gin web framework
     * Swaggo
   * Criptografia
     * Crypto (bcrypt)
   * Application (regras de negócio)
     * Validator
   * Config
     * Viper
   * Testes
     * Testing
     * Testify

## Rodando o projeto

1. Clonar o repositório:

    ```git clone https://github.com/RafaelRNunes/verify-my-age_backend-test.git```

2. Baixar dependências, rode na raiz do projeto:

    ```go mod tidy```

3. Subir container com mysql, rode na raiz do projeto:

    ```docker-compose up -d```

4. Subir projeto, rode na raiz do projeto:

    ```go run main.go```

5. Para rodar os testes, rode na raiz do projeto:

    ```go test ./test/unit```

## Swagger

1. Para acessar o swagger, insira o seguinte endereço no browser:

    ```http://localhost:8080/swagger/index.html```

2. Para gerar novamente a documentação do swagger rode o comando abaixo na raiz do projeto:

    ```swag init --parseDependency --parseInternal --parseDepth 1```

