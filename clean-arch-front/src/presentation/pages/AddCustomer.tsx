import { useState } from "react";
import { createCustomer } from "../../services/customerService";
import type { CustomerInterface } from "../../domain/interfaces/CustomerInterface";
import { useNavigate } from "react-router-dom";
import constants from "../../shared/constants/constants";

export function AddCustomerPage() {
    const [name, setName] = useState<string>("")
    const [email, setEmail] = useState<string>("")
    const [phone, setPhone] = useState<string>("")
    const [mobilePhone, setMobilePhone] = useState<string>("")
    const nav = useNavigate()


    async function addCustomer() {
        const customer: CustomerInterface = {
            name,
            email,
            phone,
            mobilePhone
        }
        await createCustomer(customer)

        nav(constants.api.CUSTOMERS_API)
    }

    return (
        <div>
            <h1>Criação de cliente</h1>
            <input type="text" onChange={e => setName(e.target.value)} placeholder="Nome" />
            <input type="text" onChange={e => setEmail(e.target.value)} placeholder="Email" />
            <input type="text" onChange={e => setPhone(e.target.value)} placeholder="Telefone" />
            <input type="text" onChange={e => setMobilePhone(e.target.value)} placeholder="Celular" />

            <button onClick={addCustomer}>Adicionar Cliente</button>
        </div>
    )
}