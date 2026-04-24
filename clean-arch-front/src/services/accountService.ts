import type { AccountInterface } from "../domain/interfaces/AccountInterface";
import type { AccountTransactionInterface } from "../domain/interfaces/AccountTransactionInterface";

import constants from "../shared/constants/constants";

export function createAccount(data: AccountInterface){
    return fetch(`${constants.api.API}${constants.api.ACCOUNTS_API}`, 
        {
            method: constants.http.POST,
            body: JSON.stringify(data),
            headers: {
                "Content-Type": constants.headers.contentTypes.JSON
            }
        })
}

export function accountDeposit(data: AccountTransactionInterface){
    return fetch(`${constants.api.API}${constants.api.ACCOUNTS_API}/${data.accountID}/deposit`, 
        {
            method: constants.http.POST,
            body: JSON.stringify(data),
            headers: {
                "Content-Type": constants.headers.contentTypes.JSON
            }
        })
}

export function accountWithdraw(data: AccountTransactionInterface){
    return fetch(`${constants.api.API}${constants.api.ACCOUNTS_API}/${data.accountID}/withdraw`, 
        {
            method: constants.http.POST,
            body: JSON.stringify(data),
            headers: {
                "Content-Type": constants.headers.contentTypes.JSON
            }
        })
}

export function accountBalance(data: AccountTransactionInterface){
    return fetch(`${constants.api.API}${constants.api.ACCOUNTS_API}/${data.accountID}/balance`).then(r => r.json())
}