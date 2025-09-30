# AGENT.md: Guia de Boas Práticas para Agentes em Projetos Go

Este documento serve como referência para criação e manutenção de agentes (services, workers, handlers, etc.) em projetos Go, seguindo os padrões do [Standard Go Project Layout](https://github.com/golang-standards/project-layout) e as recomendações do [Uber Go Style Guide PT-BR](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br).

---

## Estrutura de Projeto

- Siga o [Standard Go Project Layout](https://github.com/golang-standards/project-layout):
  - Separe pacotes internos (`internal/`), aplicações (`cmd/`), bibliotecas e módulos reutilizáveis (`pkg/`).
  - Utilize o arquivo `Makefile` para automação de tarefas.

## Organização do Código

- Agrupe funcionalidades relacionadas em pacotes.
- Evite pacotes grandes e genéricos; prefira nomes claros e objetivos.
- Mantenha o princípio da responsabilidade única para cada agente.

## Convenções de Código

- Siga o [Uber Go Style Guide PT-BR](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br):
  - Nomeie variáveis, funções e tipos de forma explícita.
  - Prefira nomes curtos para variáveis de escopo reduzido.
  - Use `err` para variáveis de erro.
  - Evite abreviações exceto as já consagradas pela comunidade Go.

## Boas Práticas para Agentes

- Cada agente deve ser implementado como uma estrutura (`struct`) com métodos bem definidos.
- Documente os métodos públicos utilizando comentários GoDoc.
- Implemente interfaces para facilitar testes e extensibilidade.
- Prefira injeção de dependências via construtores.

## Gerenciamento de Erros

- Sempre verifique e trate erros.
- Retorne erros contextuais e evite mascarar detalhes relevantes.
- Use pacotes como `errors` ou `fmt.Errorf` para encadeamento e formatação de erros.

## Concurrency & Goroutines

- Utilize goroutines de forma consciente, evitando vazamentos (leaks).
- Sincronize concorrência com canais (`chan`) e mutexes conforme necessário.
- Sempre trate o encerramento correto de goroutines com `context.Context`.

## Testes

- Mantenha testes unitários em `*_test.go` próximos aos arquivos de implementação.
- Use mocks para dependências externas.
- Execute testes via `go test` e inclua comandos relevantes no `Makefile`.

## Ferramentas de Qualidade

- Utilize ferramentas como `golint`, `go vet`, `staticcheck` e `gofmt`.
- Configure o `Makefile` para rodar lint e testes automaticamente.

## Exemplos de Estrutura de Agente

```go
type Worker struct {
    logger *zap.Logger
    repo   Repository
}

func NewWorker(logger *zap.Logger, repo Repository) *Worker {
    return &Worker{logger: logger, repo: repo}
}

func (w *Worker) Run(ctx context.Context) error {
    // Implementação do agente
}
```

---

## Referências

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Uber Go Style Guide PT-BR](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br)

---
