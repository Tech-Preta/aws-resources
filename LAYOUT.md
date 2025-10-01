# Layout Padrão do Projeto Go

Este projeto segue o [Golang Standards Project Layout](https://github.com/golang-standards/project-layout), proporcionando uma estrutura organizada e padronizada para aplicações Go.

## Estrutura de Diretórios

### Diretórios Go

#### `/cmd`
Aplicações principais para este projeto.

O nome do diretório para cada aplicação deve corresponder ao nome do executável que você deseja ter (por exemplo, `/cmd/myapp`).

Não coloque muito código no diretório da aplicação. Se você achar que o código pode ser importado e usado em outros projetos, então ele deve estar no diretório `/pkg`. Se o código não é reutilizável ou se você não quer que outros o reutilizem, coloque esse código no diretório `/internal`. Você ficará surpreso com o que outros farão, então seja explícito sobre suas intenções!

É comum ter uma pequena função `main` que importa e invoca o código dos diretórios `/internal` e `/pkg` e nada mais.

Veja o diretório [`/cmd`](cmd/) para exemplos.

#### `/internal`
Código privado da aplicação e biblioteca. Este é o código que você não quer que outros importem em suas aplicações ou bibliotecas. Note que este padrão de layout é imposto pelo próprio compilador Go. Veja as [notas de lançamento Go 1.4](https://golang.org/doc/go1.4#internalpackages) para mais detalhes. Note que você não está limitado ao diretório `internal` de nível superior. Você pode ter mais de um diretório `internal` em qualquer nível da sua árvore de projeto.

Você pode opcionalmente adicionar um pouco de estrutura extra aos seus pacotes internos para separar seu código interno compartilhado e não compartilhado. Não é obrigatório (especialmente para projetos menores), mas é bom ter pistas visuais mostrando o uso pretendido do pacote. Seu código de aplicação atual pode ir no diretório `/internal/app` (por exemplo, `/internal/app/myapp`) e o código compartilhado por essas aplicações no diretório `/internal/pkg` (por exemplo, `/internal/pkg/myprivatelib`).

#### `/pkg`
Código de biblioteca que está ok para usar por aplicações externas (por exemplo, `/pkg/mypubliclib`). Outros projetos irão importar essas bibliotecas e esperar que elas funcionem, então pense duas vezes antes de colocar algo aqui :-) Note que o diretório `internal` é uma maneira melhor de garantir que seus pacotes privados não sejam importáveis porque é imposto pelo Go. O diretório `/pkg` ainda é uma boa maneira de comunicar explicitamente que o código nesse diretório é seguro para uso por outros. O post [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) por Travis Jeffery fornece uma boa visão geral dos diretórios `pkg` e `internal` e quando pode fazer sentido usá-los.

É também uma maneira de agrupar código Go em um local quando seu diretório raiz contém muitos componentes e diretórios não-Go, tornando mais fácil executar várias ferramentas Go (como mencionado nessas palestras: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) de GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) e [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

Veja o diretório [`/pkg`](pkg/) se você quiser ver quais repositórios Go populares usam esse padrão de layout de projeto. Este é um padrão de layout comum, mas não é universalmente aceito e alguns na comunidade Go não o recomendam.

Está ok não usá-lo se seu projeto de aplicação for realmente pequeno e onde um nível extra de aninhamento não agrega muito valor (a menos que você realmente queira :-)). Pense nisso quando estiver ficando grande o suficiente e seu diretório raiz ficar bastante ocupado (especialmente se você tiver muitos componentes de aplicação não-Go).

### Diretórios de Aplicação de Serviço

#### `/api`
Definições de API (arquivos OpenAPI/Swagger, arquivos de esquema JSON, arquivos de definição de protocolo).

Veja o diretório [`/api`](api/) para exemplos.

### Diretórios de Aplicação Web

#### `/web`
Componentes específicos de aplicação web: assets web estáticos, templates do lado do servidor e SPAs.

### Diretórios Comuns de Aplicação

#### `/configs`
Modelos de arquivo de configuração ou configurações padrão.

Coloque seus arquivos de modelo `confd` ou `consul-template` aqui.

#### `/init`
Configurações de init de sistema (systemd, upstart, sysv) e gerenciador de processos/supervisor (runit, supervisord).

#### `/scripts`
Scripts para executar várias operações de build, instalação, análise, etc.

Esses scripts mantêm o Makefile de nível raiz pequeno e simples (por exemplo, [https://github.com/hashicorp/terraform/blob/master/Makefile](https://github.com/hashicorp/terraform/blob/master/Makefile)).

Veja o diretório [`/scripts`](scripts/) para exemplos.

#### `/build`
Empacotamento e Integração Contínua.

Coloque suas configurações e scripts de empacotamento de cloud (AMI), contêiner (Docker), OS (deb, rpm, pkg) no diretório `/build/package`.

Coloque suas configurações e scripts de CI (travis, circle, drone) no diretório `/build/ci`. Note que algumas das ferramentas de CI (por exemplo, Travis CI) são muito exigentes sobre a localização dos seus arquivos de configuração. Tente colocar os arquivos de configuração no diretório `/build/ci` linkando eles para o local onde as ferramentas de CI esperam (quando possível).

#### `/deployments`
Configurações de implantação IaaS, PaaS, sistema e orquestração de contêineres e modelos (docker-compose, kubernetes/helm, mesos, terraform, bosh). Note que em alguns repositórios (especialmente aplicações implantadas com kubernetes) esse diretório é chamado `/deploy`.

#### `/test`
Aplicações de teste externas adicionais e dados de teste. Sinta-se livre para estruturar o diretório `/test` da maneira que quiser. Para projetos maiores, faz sentido ter um subdiretório de dados. Por exemplo, você pode ter `/test/data` ou `/test/testdata` se precisar que o Go ignore o que está nesse diretório. Note que o Go também ignorará diretórios ou arquivos que começam com "." ou "_", então você tem mais flexibilidade em termos de como nomear seu diretório de dados de teste.

Veja o diretório [`/test`](test/) para exemplos.

### Outros Diretórios

#### `/docs`
Documentos de design e usuário (além da documentação godoc gerada automaticamente).

Veja o diretório [`/docs`](docs/) para exemplos.

#### `/examples`
Exemplos para suas aplicações e/ou bibliotecas públicas.

Veja o diretório [`/examples`](examples/) para exemplos.

#### `/assets`
Outros assets que acompanham o repositório (imagens, logos, etc).

## Diretórios que você não deveria ter

### `/src`
Alguns projetos Go têm uma pasta `src`, mas isso geralmente acontece quando os desenvolvedores vieram do mundo Java onde é um padrão comum. Se você pode se ajudar, tente não adotar esse padrão Java. Você realmente não quer que seu código Go ou projetos Go se pareçam com Java :-)

Não confunda o diretório `/src` do nível do projeto com o diretório `/src` que o Go usa para seus workspaces conforme descrito em [`How to Write Go Code`](https://golang.org/doc/code.html). A variável de ambiente `$GOPATH` aponta para seu workspace atual (por padrão ela aponta para `$HOME/go` em sistemas não-Windows). Este workspace inclui os diretórios de nível superior `/pkg`, `/bin` e `/src`. Seu projeto atual acaba sendo um subdiretório sob `/src`, então se você tiver o diretório `/src` em seu projeto, o caminho do projeto parecerá com isso: `/some/path/to/workspace/src/your_project/src/your_code.go`. Note que com Go 1.11 é possível ter seu projeto fora do seu `GOPATH`, mas isso ainda não significa que seja uma boa ideia usar esse padrão de layout.

## Badges

* [Go Report Card](https://goreportcard.com/) - Ele escaneará seu código com `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` e `misspell`. Substitua `github.com/golang-standards/project-layout` pela referência do seu projeto.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - Ele fornecerá uma versão online da documentação gerada pelo GoDoc. Mude o link para apontar para seu projeto.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev é um novo destino para documentação Go. Você pode criar um badge usando a [ferramenta de geração de badge](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - Ele mostrará o número da versão mais recente para seu projeto. Mude o link do github para apontar para seu projeto.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)

## Notas

Um layout de projeto mais opinativo está disponível junto com exemplos, configurações reutilizáveis, scripts e código no repositório [`golang-templates/seed`](https://github.com/golang-templates/seed).

## Estrutura Atual do Projeto

```
.
├── api/                    # Definições de API
├── assets/                 # Assets do repositório
├── build/                  # Configurações de build e CI
│   ├── ci/                # Configurações de CI/CD
│   └── package/           # Scripts de empacotamento
├── cmd/                   # Aplicações principais
│   └── aws-resources/     # Aplicação principal
├── configs/               # Configurações e templates
├── deployments/           # Configurações de deployment
│   ├── docker/           # Configurações Docker
│   └── kubernetes/       # Configurações Kubernetes
├── docs/                  # Documentação
├── examples/              # Exemplos de uso
│   └── basic-usage/      # Exemplo básico
├── internal/              # Código privado
│   ├── app/              # Código da aplicação privada
│   └── pkg/              # Bibliotecas privadas
├── pkg/                   # Código público
│   ├── cli/              # Interface TUI
│   └── services/         # Serviços AWS
├── scripts/               # Scripts de build e desenvolvimento
└── test/                  # Testes adicionais e dados de teste
```