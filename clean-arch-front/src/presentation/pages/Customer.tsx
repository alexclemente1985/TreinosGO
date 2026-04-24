/* eslint-disable react-hooks/immutability */
import { useEffect, useState } from "react";
import type { CustomerInterface } from "../../domain/interfaces/CustomerInterface";
import { getCustomers } from "../../services/customerService";
import { useNavigate } from "react-router-dom";
import constants from "../../shared/constants/constants";

export default function CustomerPage(){
    const [customers, setCustomers] = useState<CustomerInterface[]>([])
    const nav = useNavigate()
    
    useEffect(() => {
        load()
    }, [])

    async function load() {
        const data = await getCustomers()
        setCustomers(data)
    }

    function addCustomer(){
        nav(`${constants.api.CUSTOMERS_API}/add-customer`)
    }

    

    return (
        <div>
            <h1>Tela de Clientes</h1>
            {customers.map(c => (
                <div key={c.id}>{c.name}</div>
            ))}
            <button onClick={addCustomer}>Adicionar novo cliente</button>
        </div>
    )

    
}