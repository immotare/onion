import { useState } from 'react'
import reactLogo from './assets/react.svg'
import './App.css'
import { BrowserRouter, Route, Router, Routes } from 'react-router-dom'
import { Home } from './pages/Home';
import { SignIn } from './pages/SignIn';
import { CreateAccount } from './pages/CreateAccount';
import { AuthContextProvider } from './firebase_utils/authContext';

function App() {
  return (
    <AuthContextProvider>
      <BrowserRouter>
        <Routes>
            <Route path='/' element={<Home/>}></Route>
            <Route path='/signin' element={<SignIn/>}></Route>
            <Route path='/create_account' element={<CreateAccount/>}></Route>
        </Routes>
      </BrowserRouter>
    </AuthContextProvider>
  );
}

export default App
