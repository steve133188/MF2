
import {useRouter} from "next/router";
import {useEffect , useContext} from "react";
import {AuthContext} from "../context/authContext"
export  default function Layout ({children}){
    const router = useRouter()
    const {user , login} = useContext(AuthContext)
    let layout =(
        <div className={"layout"}>
            {children}

        </div>
    )
    useEffect(()=>{
        if(!user["authReady"]){
            router.push("/login")
            layout = (<div className={"unauth"}>{children}</div>)
        }
    },[])
    return(
        layout
    )
}