import type { CustomerInterface } from "../domain/interfaces/CustomerInterface";
import constants from "../shared/constants/constants";

export async function getCustomers(){
    const token = localStorage.getItem(constants.localStorage.ACCESS_TOKEN)
    
    return fetch(`${constants.api.API}${constants.api.CUSTOMERS_API}/`,{
        headers: {
            Authorization: constants.headers.authorization.BEARER_TOKEN.replace("%s",token??"")
        }
    }).then(r => r.json())
}

export async function createCustomer(data: CustomerInterface){
    const token = localStorage.getItem(constants.localStorage.ACCESS_TOKEN)
    console.log("TOKEN ---- ",token)
    return fetch(`${constants.api.API}${constants.api.CUSTOMERS_API}/`, 
        {
            method: constants.http.POST,
            body: JSON.stringify(data),
            headers: {
                "Content-Type": constants.headers.contentTypes.JSON,
                "Authorization": constants.headers.authorization.BEARER_TOKEN.replace("%s",token??"")
            }
        }).then(r => r.json())
}