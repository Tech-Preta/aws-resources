# Exemplo de Uso Básico

Este exemplo demonstra como usar a biblioteca aws-resources programaticamente para criar recursos AWS.

## Executando o Exemplo

1. Certifique-se de ter credenciais AWS configuradas:
   ```bash
   aws configure
   ```

2. Execute o exemplo:
   ```bash
   go run main.go
   ```

## O que o Exemplo Faz

- Cria um bucket S3 com nome único
- Cria uma instância EC2 com configurações básicas
- Mostra como tratar os resultados das operações

## Requisitos

- Go 1.21+
- Credenciais AWS configuradas
- Permissões IAM para S3 e EC2