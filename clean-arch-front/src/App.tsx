import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import './App.css'
import CustomerPage from './presentation/pages/Customer'
import { AddCustomerPage } from './presentation/pages/AddCustomer'
import LoginPage from './presentation/pages/Login'
import constants from './shared/constants/constants'
import { PrivateRoute } from './presentation/routes/PrivateRoute'

function App() {
  const token = localStorage.getItem(constants.localStorage.ACCESS_TOKEN)
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={token ? <Navigate to={constants.api.CUSTOMERS_API} replace/> : <LoginPage/>}/>
        <Route 
          path={constants.api.CUSTOMERS_API}
          element={
            <PrivateRoute>
              <CustomerPage/>
            </PrivateRoute>
          }
        />
        <Route path={`${constants.api.CUSTOMERS_API}/add-customer` }element={
          <PrivateRoute>
            <AddCustomerPage/>
          </PrivateRoute>          
          }/>
          {/* <Route path={constants.api.ACCOUNTS_API} element={
          <PrivateRoute>
            <AddCustomerPage/>
          </PrivateRoute>          
          }/> */}
      </Routes>
    </BrowserRouter>
  )
}

export default App
