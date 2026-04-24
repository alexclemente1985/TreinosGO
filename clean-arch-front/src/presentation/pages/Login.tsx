import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { loginService } from "../../services/authService";
import constants from "../../shared/constants/constants";

export default function LoginPage(){
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [status, setStatus] = useState(0)

    const nav = useNavigate()

    async function login(){
       const status: number = await loginService(email, password)

       if (status == constants.http.status.OK){
        nav(constants.api.CUSTOMERS_API)
       }
       setStatus(status)
    }

    function errorMsgHandler(){
        if(status != 0 && status != constants.http.status.OK){
            return <h2 style={{"color": "red"}}>Falha no Login</h2>
        }
    }

    return (
        <div>
            <h1>Tela de Login</h1>
            <input type="text" onChange={e => setEmail(e.target.value)} placeholder="Email" />
            <input type="password" onChange={e => setPassword(e.target.value)} placeholder="Senha" />

            <button onClick={login}>Login</button>
            {errorMsgHandler()}
        </div>
    )
}