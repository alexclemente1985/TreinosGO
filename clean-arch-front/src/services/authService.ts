import constants from "../shared/constants/constants"

export async function loginService (email: string, password: string): Promise<number> {
    return await fetch(`${constants.api.API}${constants.api.AUTH_API}/login`,{
        method: constants.http.POST,
        body: JSON.stringify({email, password}),
        headers: {"Content-Type": constants.headers.contentTypes.JSON}
    }).then(
        async (r: Response)=> {
            if (r.status == constants.http.status.OK){
                const {accessToken, refreshToken} = await r.json()
                
                localStorage.setItem(constants.localStorage.ACCESS_TOKEN, accessToken)
                localStorage.setItem(constants.localStorage.REFRESH_TOKEN, refreshToken)
            }
            return r.status
        }
    )
}