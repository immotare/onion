import React, { Context, createContext, Provider, ReactNode, useEffect, useState } from "react";
import {onAuthStateChanged, User} from "firebase/auth";
import {firebaseAuth} from "./auth";

type FirebaseAuthProperties = {
  user: User | null,
}

const authProperties: FirebaseAuthProperties = {
  user: firebaseAuth.currentUser,
}

export const AuthContext = createContext(authProperties);

export const AuthContextProvider: React.FC<{children: ReactNode}> = ({children}: {children: ReactNode}) => {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    onAuthStateChanged(firebaseAuth, (user) => {
      setUser(user);
    });
  }, []);

  return (
    <AuthContext.Provider value={{user}}>
      {children}
    </AuthContext.Provider>
  )
};
