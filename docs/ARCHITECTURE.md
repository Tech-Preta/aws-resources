# Arquitetura do AWS Resources

## Visão Geral

O AWS Resources é uma ferramenta de linha de comando moderna construída em Go que fornece uma interface de usuário terminal (TUI) interativa para criar e gerenciar recursos AWS.

## Estrutura do Projeto

### Diretórios Principais

- **`/cmd`**: Aplicações principais
  - `/cmd/aws-resources`: Ponto de entrada principal da aplicação
  
- **`/pkg`**: Código de biblioteca público
  - `/pkg/cli`: Interface TUI usando Bubble Tea
  - `/pkg/services`: Implementações dos serviços AWS
  
- **`/internal`**: Código privado da aplicação
  - `/internal/app`: Lógica específica da aplicação
  - `/internal/pkg`: Bibliotecas privadas compartilhadas

### Padrões Arquiteturais

#### 1. Padrão Service Layer
- Cada serviço AWS (S3, EC2) tem sua própria implementação
- Interface comum `ResourceService` para todos os serviços
- BaseService fornece funcionalidade comum

#### 2. Padrão MVC com Bubble Tea
- **Model**: Estado da aplicação e dados
- **View**: Renderização da interface do usuário
- **Update**: Manipulação de eventos e atualizações de estado

#### 3. Dependency Injection
- Serviços são injetados via construtores
- Configuração centralizada via AWS SDK

## Fluxo de Dados

```
Usuário -> TUI (Bubble Tea) -> CLI Handler -> Service Layer -> AWS SDK -> AWS API
```

1. Usuário interage com a TUI
2. Eventos são processados pelo Bubble Tea
3. Comandos são enviados para o Service Layer
4. Serviços fazem chamadas para APIs AWS
5. Resultados são retornados através da cadeia

## Componentes Principais

### CLI Application (`/pkg/cli`)
- Interface TUI usando Bubble Tea framework
- Gerenciamento de estado da aplicação
- Navegação entre telas e formulários

### Service Layer (`/pkg/services`)
- Abstrações para serviços AWS
- Tratamento de erros unificado
- Validação de parâmetros

### AWS Integration
- Uso do aws-sdk-go-v2
- Configuração automática de credenciais
- Suporte a múltiplas regiões

## Extensibilidade

### Adicionando Novos Serviços AWS

1. Implementar a interface `ResourceService`
2. Criar arquivo no diretório `/pkg/services`
3. Adicionar opções de menu na TUI
4. Implementar formulários específicos

### Personalizando a TUI

1. Modificar estilos em `/pkg/cli/app.go`
2. Adicionar novas telas seguindo o padrão existente
3. Implementar novos tipos de formulário

## Segurança

- Credenciais AWS são gerenciadas pelo AWS SDK
- Não há armazenamento de credenciais no código
- Validação de entrada rigorosa
- Princípio do menor privilégio