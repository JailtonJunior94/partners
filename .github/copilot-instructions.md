# GitHub Copilot CLI

Use `AGENTS.md` como instrucao principal deste repositorio.

## Instrucoes

1. Ler `AGENTS.md` no inicio da sessao.
2. `.agents/skills/` e a fonte de verdade dos fluxos procedurais.
3. `.github/agents/` sao wrappers leves que apontam para a habilidade correta.
4. Carregar referencias adicionais apenas quando a tarefa exigir.
5. Preservar estilo, arquitetura e fronteiras existentes antes de propor mudancas.
6. Validar mudancas com comandos proporcionais ao risco.

## Stack

- Projeto com contexto Go detectado: carregar `.agents/skills/go-implementation/SKILL.md` ao alterar codigo Go.
- Validar a versao declarada em `go.mod` antes de introduzir APIs da linguagem ou novas dependencias.
