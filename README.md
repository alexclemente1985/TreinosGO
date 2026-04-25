# Projeto JWT Secret Generator

## Inicialização projeto Wails + React ts

1) `wails init -n <nome_projeto> -t react-ts`
2) Apague a pasta `frontend` gerada
3) Dentro do projeto, no terminal: `npm create vite@latest frontend -- --template react-ts`
4) No arquivo `wails.json`:
- Acrescentar `"frontend:dir": "frontend",`
- Atualizar: `"frontend:dev:serverUrl": "http://localhost:5173",`
- Atualizar os parâmetros de `author` para os desejados