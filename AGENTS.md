# Regras para Agentes de IA

Este diretorio centraliza regras para uso com agentes de IA em tarefas reais de analise, alteracao e validacao de codigo.

## Objetivo

Use estas instrucoes para manter consistencia, seguranca e qualidade ao trabalhar com codigo, configuracao, validacao e evolucao de sistemas.

## Arquitetura: monorepo

O projeto aparenta ser um monorepo, com dois servicos independentes (`address` e `partner`) sob a mesma raiz e compartilhando um unico `go.mod`. Cada servico possui Dockerfile proprio, manifests k8s independentes e entry point separado. A comunicacao entre eles e via gRPC. A governanca deve preservar fronteiras entre servicos e validar apenas os componentes afetados.

Stack detectada: Go.
Frameworks detectados: gRPC, gqlgen (GraphQL).

## Estrutura de Pastas

```
.
.docs
.docs/Go-Challenge.png
.docs/gochallenge.puml
.github
.github/agents
.github/agents/bugfix.agent.md
.github/agents/prd-writer.agent.md
.github/agents/project-analyzer.agent.md
.github/agents/refactorer.agent.md
.github/agents/reviewer.agent.md
.github/agents/task-executor.agent.md
.github/agents/task-planner.agent.md
.github/agents/technical-specification-writer.agent.md
.github/skills
.github/skills/agent-governance
.github/skills/analyze-project
.github/skills/bugfix
.github/skills/create-prd
.github/skills/create-tasks
.github/skills/create-technical-specification
.github/skills/execute-task
.github/skills/go-implementation
.github/skills/object-calisthenics-go
.github/skills/refactor
.github/skills/review
.github/workflows
.github/workflows/ci-cd.yaml
.gitignore
.k8s
.k8s/certmanager
.k8s/certmanager/clusterissuer.yml
.k8s/deployments
.k8s/deployments/address-api-dp.yaml
.k8s/deployments/partner-api-dp.yaml
.k8s/ingress
.k8s/ingress/ingress.yaml
.k8s/namespaces
.k8s/namespaces/partner-ns.yaml
.k8s/services
.k8s/services/address-api-svc.yaml
.k8s/services/partner-api-svc.yaml
.vscode
.vscode/launch.json
address
address/Dockerfile
address/address.go
address/enviroments
address/enviroments/enviroment.go
address/facades
address/facades/geocode.go
address/facades/via_cep.go
address/pb
address/pb/cep.pb.go
address/pb/cep_grpc.pb.go
address/protofiles
address/protofiles/cep.proto
address/services
address/services/cep.go
address/settings.json
address/structs
address/structs/geocode.go
address/structs/via_cep.go
address/validators
address/validators/cep_validator.go
docker-compose.yaml
go.mod
go.sum
partner
partner/Dockerfile
partner/domain
partner/domain/entities
partner/domain/entities/address.go
partner/domain/entities/partner.go
partner/domain/mappings
partner/domain/mappings/partner.go
partner/facades
partner/facades/address.go
partner/gqlgen.yml
partner/graph
partner/graph/generated
partner/graph/generated/generated.go
partner/graph/model
partner/graph/model/models_gen.go
partner/graph/resolver.go
partner/graph/schema.graphqls
partner/graph/schema.resolvers.go
partner/infrastructure
partner/infrastructure/database
partner/infrastructure/database/mongo.go
partner/infrastructure/environments
partner/infrastructure/environments/environment.go
partner/infrastructure/repositories
partner/infrastructure/repositories/partner.go
partner/server.go
partner/settings.Development.json
partner/settings.Production.json
readme.md
terraform
terraform/acr.tf
terraform/aks.tf
terraform/cosmos_db.tf
terraform/main.tf
terraform/output.tf
terraform/variables.tf
```

## Padrao Arquitetural

Arquitetura em Camadas com separacao de dominio. O servico `partner` apresenta camadas claras: `domain/entities` (modelos de dominio), `domain/mappings` (DTOs), `infrastructure/` (banco e repositorios), `facades/` (anti-corruption layer para servicos externos via gRPC), e `graph/` (transporte GraphQL). O servico `address` segue estrutura mais simples: `services/` (implementacao gRPC), `facades/` (clientes HTTP externos — ViaCEP, Google Geocode), `structs/` (modelos de resposta externa), `validators/` (validacao de entrada), e `pb/` (codigo gerado protobuf).

### Servicos

| Servico | Porta | Transporte | Banco | Descricao |
|---------|-------|------------|-------|-----------|
| `address` | 9000 | gRPC | — | Consulta de CEP via ViaCEP e geocodificacao via Google Geocode |
| `partner` | 8000 | GraphQL | MongoDB (CosmosDB em prod) | CRUD de parceiros com busca por proximidade geografica |

### Fluxo de Dependencias

- `partner` → `address` via gRPC (facade `partner/facades/address.go` consome `address/pb`)
- `partner` → MongoDB via `infrastructure/repositories` e `infrastructure/database`
- `address` → APIs externas (ViaCEP, Google Geocode) via `facades/`
- Transporte e adapters devem depender de casos de uso ou servicos explicitos, nao do contrario.
- Dominio nao deve conhecer detalhes de HTTP, banco, filas, serializacao ou drivers.
- Infraestrutura pode implementar contratos consumidos pela aplicacao, preservando dependencia para dentro.

## Modo de trabalho

1. Entender o contexto antes de editar qualquer arquivo.
2. Preferir a menor mudanca segura que resolva a causa raiz.
3. Preservar arquitetura, convencoes e fronteiras ja existentes no contexto analisado.
4. Nao introduzir abstracoes, camadas ou dependencias sem demanda concreta.
5. Atualizar ou adicionar testes quando houver mudanca de comportamento.
6. Rodar validacoes proporcionais a mudanca.
7. Registrar bloqueios e suposicoes explicitamente quando o contexto estiver incompleto.

## Diretrizes de Estrutura

1. Priorize entendimento do codigo e do contexto atual antes de propor refatoracoes.
2. Respeite padroes existentes de nomenclatura, organizacao e tratamento de erro.
3. Defina estrutura simples, evolutiva e com defaults explicitos.
4. Evite reescritas amplas quando uma alteracao localizada resolver o problema.
5. Estabeleca contratos, testes e comandos de validacao cedo quando eles ainda nao existirem.
6. Considere risco de regressao como restricao principal.
7. Evite overengineering disfarcado de arquitetura futura.

## Regras por Arquitetura

1. Limitar mudancas ao servico afetado (`address/` ou `partner/`).
2. Nao criar dependencias internas cruzadas sem contrato explicito (o unico contrato atual e o protobuf em `address/protofiles/cep.proto`).
3. Validar primeiro apenas o servico impactado antes de ampliar o escopo.
4. Alteracoes no contrato gRPC (`cep.proto`) devem regenerar codigo protobuf e validar ambos os servicos.
5. Nao alterar contratos entre servicos sem deixar o impacto explicito.

## Regras por Linguagem

Para tarefas que alteram codigo, carregar a skill:

- `.agents/skills/agent-governance/SKILL.md`

Para tarefas que alteram codigo Go, carregar tambem:

- `.agents/skills/go-implementation/SKILL.md`

Para tarefas de revisao ou refatoracao incremental de design em Go guiadas por heuristicas de object calisthenics, carregar tambem:

- `.agents/skills/object-calisthenics-go/SKILL.md`

Para tarefas de correcao de bugs com remediacao e teste de regressao, carregar tambem:

- `.agents/skills/bugfix/SKILL.md`

### Composicao Multi-Linguagem

Em projetos com mais de uma linguagem (ex: monorepo Go + Node), carregar apenas a skill da linguagem afetada pela mudanca. Se a tarefa cruzar linguagens, carregar ambas e aplicar a validacao de cada stack nos arquivos correspondentes. Nao misturar convencoes de uma linguagem em arquivos de outra.

## Referencias

Cada skill lista suas proprias referencias em `references/` com gatilhos de carregamento no respectivo `SKILL.md`. Nao duplicar a listagem aqui — consultar o SKILL.md da skill ativa para saber quais referencias carregar e em que condicao.

## Validacao

Antes de concluir uma alteracao:

Seguir Etapa 4 de `.agents/skills/agent-governance/SKILL.md` como base canonica.

Comandos detectados no projeto (Go):
1. Rodar fmt: `gofmt -w .`.
2. Rodar test: `go test ./...`.
3. Rodar lint: `golangci-lint run`.

## Restricoes

1. Nao inventar contexto ausente.
2. Nao assumir versao de linguagem, framework ou runtime sem verificar.
3. Nao alterar comportamento publico sem deixar isso explicito.
4. Nao usar exemplos como copia cega; adaptar ao contexto real.
5. Nao alterar contratos entre servicos sem deixar o impacto explicito.
