# Projeto JWT Secret Generator

## Inicialização projeto Wails + React ts

1) `wails init -n <nome_projeto> -t react-ts`
2) Apague a pasta `frontend` gerada
3) Dentro do projeto, no terminal: `npm create vite@latest frontend -- --template react-ts`
4) No arquivo `wails.json`:
- Acrescentar `"frontend:dir": "frontend",`
- Atualizar: `"frontend:dev:serverUrl": "http://localhost:5173",`
- Atualizar os parâmetros de `author` para os desejados

## Preparação do projeto 
- Na pasta do projeto: `wails dev`

## Geração de build
- Na pasta do projeto: `wails build`

## Acesso aos métodos e funções do backend
- `wails dev` irá fazer os bindings em JavaScript, bastando importar a função desejada
-- `import {FUNÇÃO} from './wailsjs/go/main/App'`