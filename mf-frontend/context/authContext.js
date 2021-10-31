import {useEffect, createContext, Context, useState} from "react";
import {useRouter} from "next/router";
import {redirect} from "next/dist/server/api-utils";
import AuthService from "../services/auth";

export const AuthContext = createContext({
    userInfo:null,
    authReady : false,
})

export const AuthContextProvider = ({children}) =>{
    const [user , setUser] = useState({
        userInfo: {
            name:null ,
            email:null ,
            role:null,
            organization:{},
            authority:[],
            phone:null,
            status:"online",
            notification:0,
        } ,
        authReady : false,
    })

    const router = useRouter()
    useEffect(()=>{
        if(!user){
            router.push("/login")
        }
    },[])
    const login = (email , pwd)=>{
        if(email == "wiva.wei@matrixsense.tech" && pwd == "1234"){
            console.log("login success")
            setUser({userInfo:{name:"Wiva " , email: email , role:"super admin" , organization: {"Matrixsense":"CEO"}}, authReady: true})
            router.push("/dashboard")
        }else if (email == "steve.chak@matrixsense.tech" && pwd =="1234"){
            setUser({userInfo:{name:"Steve.Chak " }, authReady: true})
            router.push("/dashboard")
        }else if(email =="ben.cheng@matrixsense.tech" && pwd == "1234"){
            setUser({userInfo:{name:"Ben.cheng " , email: email , role:"admin" , organization: {"Matrixsense":"Developer"}}, authReady: true})
            router.push("/dashboard")
        }else if(email =="lewis.chan@matrixsense.tech"  && pwd == "1234"){
            setUser({userInfo:{name:"Lewis.chan " , email: email , role:"admin" , organization: {"Matrixsense":"Developer"}}, authReady: true})
            router.push("/dashboard")
        }else{
            console.log("Something went Wrong")
            return "Something went Wrong"
        }
    }
    return(
        <AuthContext.Provider value={{user, login}}>{children}</AuthContext.Provider>
    )
}