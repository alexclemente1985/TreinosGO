import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import heroImg from './assets/hero.png'
import './App.css'
import { GenerateJwtSecret } from '../wailsjs/go/main/App';

function App() {
  const [secret, setSecret] = useState<string>('')

  const handleGenerate = async () => {
    try {
      // 32 bytes = 256 bits (segurança ideal para HS256)
      const result = await GenerateJwtSecret(32)
      setSecret(result) // Atualiza o estado para mostrar na tela
    } catch (err) {
      alert(`Erro ao gerar secret: ${err}`)
    }
  }

  return (
    <>
      <section id="center">
        <div className="hero">
          <img src={heroImg} className="base" width="170" height="179" alt="" />
          <img src={reactLogo} className="framework" alt="React logo" />
          <img src={viteLogo} className="vite" alt="Vite logo" />
        </div>
        <div>
          <h1>JWT Secret Generator</h1>
          <p>
            Gerador simples de <code>JWT Secret</code> em <code>Base64</code>
          </p>
        </div>

        <div className="card">
        <button className='counter' onClick={handleGenerate}>
          Gerar Novo Segredo
        </button>

        {secret && (
          <div style={{ marginTop: '20px' }}>
            <p>Seu Segredo (Base64):</p>
            <code style={{ 
              wordBreak: 'break-all', 
              background: '#333', 
              padding: '10px', 
              display: 'block' 
            }}>
              {secret}
            </code>
            <button 
              onClick={() => navigator.clipboard.writeText(secret)}
              className='counter'
              style={{ marginTop: '10px', fontSize: '0.8em' }}
            >
              Copiar para o Clipboard
            </button>
          </div>
        )}
      </div>

       
      </section>

      <div className="ticks"></div>

    
      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  )
}

export default App
