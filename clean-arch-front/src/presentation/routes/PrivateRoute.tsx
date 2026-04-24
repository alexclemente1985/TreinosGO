import { Navigate } from "react-router-dom";
import constants from "../../shared/constants/constants";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function PrivateRoute({children}: any){
    const token = localStorage.getItem(constants.localStorage.ACCESS_TOKEN)

    if (!token) return <Navigate to="/"/>
    
    return children
}