import {useEffect, createContext, Context, useState} from "react";
import {useRouter} from "next/router";
import {redirect} from "next/dist/server/api-utils";

export const AuthContext = createContext({
    user:null,
    authReady : false
})

export const AuthContextProvider = ({children}) =>{
    const [user , setUser] = useState({user:null , authReady : false})


    const router = useRouter()
    useEffect(()=>{
        if(!user){
            router.push("/login")
        }
    },[])
    const login = (email , pwd)=>{
        console.log(email)
        console.log(pwd)
        if(email == "wiva.wei@matrixsense.tech" && pwd == "1234"){
            console.log("login success")
            setUser({user:"Wiva", authReady: true})
            router.push("/testing")
        }else{
            console.log("Something went Wrong")
            return "Something went Wrong"
        }
    }
    return(
        <AuthContext.Provider value={{user, login}}>{children}</AuthContext.Provider>
    )
}