# Guia de Desenvolvimento

## Configuração do Ambiente

### Pré-requisitos

- Go 1.21+
- Git
- AWS CLI (para testes com recursos reais)
- Make

### Configuração Inicial

1. Clone o repositório:
   ```bash
   git clone https://github.com/Tech-Preta/aws-resources.git
   cd aws-resources
   ```

2. Instale dependências:
   ```bash
   go mod download
   ```

3. Configure credenciais AWS (opcional, para testes):
   ```bash
   aws configure
   ```

## Scripts de Desenvolvimento

### Build
```bash
make build
# ou
./scripts/build.sh
```

### Testes
```bash
make test
# ou
./scripts/test.sh
```

### Linting
```bash
make check
# ou
./scripts/lint.sh
```

## Estrutura de Desenvolvimento

### Adicionando um Novo Serviço AWS

1. **Criar o serviço**:
   ```bash
   # Crie o arquivo do serviço
   touch pkg/services/novo_servico.go
   ```

2. **Implementar a interface**:
   ```go
   type NovoServico struct {
       *BaseService
       client *novoservico.Client
   }

   func (ns *NovoServico) CreateResource(ctx context.Context, params map[string]interface{}) (*ResourceResult, error) {
       // Implementação
   }
   ```

3. **Adicionar à TUI**:
   - Adicionar menu option em `pkg/cli/app.go`
   - Implementar formulário específico
   - Adicionar tratamento de eventos

4. **Testes**:
   ```bash
   # Criar arquivo de teste
   touch pkg/services/novo_servico_test.go
   ```

### Padrões de Código

#### Convenções de Nomenclatura
- Pacotes: lowercase
- Funções públicas: PascalCase
- Funções privadas: camelCase
- Constantes: UPPER_CASE

#### Tratamento de Erros
```go
if err != nil {
    return &ResourceResult{
        Success: false,
        Error:   "ErrorCode",
        Message: "User-friendly message",
    }, nil
}
```

#### Validação de Parâmetros
```go
if err := s.ValidateRequiredParams(params, []string{"required_param"}); err != nil {
    return &ResourceResult{
        Success: false,
        Error:   "ValidationError",
        Message: err.Error(),
    }, nil
}
```

## Debugging

### Debug da TUI
```bash
# Execute com debug logs
DEBUG=1 ./bin/aws-resources
```

### Debug dos Serviços AWS
```bash
# Execute testes específicos
go test -v ./pkg/services -run TestS3Service
```

### Profiling
```bash
# Build com profiling
go build -race -o bin/aws-resources ./cmd/aws-resources

# Execute com profiling
./bin/aws-resources -cpuprofile=cpu.prof
```

## Contribuindo

### Fluxo de Trabalho

1. **Fork o repositório**
2. **Crie uma branch feature**:
   ```bash
   git checkout -b feature/nova-funcionalidade
   ```
3. **Faça suas alterações**
4. **Execute testes**:
   ```bash
   make check
   ```
5. **Commit suas alterações**:
   ```bash
   git commit -m "feat: adiciona nova funcionalidade"
   ```
6. **Push para sua branch**:
   ```bash
   git push origin feature/nova-funcionalidade
   ```
7. **Crie um Pull Request**

### Mensagens de Commit

Use o padrão Conventional Commits:
- `feat:` nova funcionalidade
- `fix:` correção de bug
- `docs:` documentação
- `style:` formatação
- `refactor:` refatoração
- `test:` testes
- `chore:` manutenção

### Revisão de Código

- Todos os PRs precisam de revisão
- Testes devem passar
- Cobertura de código mantida
- Documentação atualizada