<<<<<<< HEAD
export default function Layout (props){

    return(
        <div>
            {props.children}
=======
import {useRouter} from "next/router";
import {useEffect , useContext} from "react";
import {AuthContext} from "../context/authContext"
export  default function Layout ({children}){
    const router = useRouter()
    const {user , login} = useContext(AuthContext)
    let layout =(
        <div className={"layout"}>
            {children}
>>>>>>> 863d9e42ff766350b88638972936682852a0b635
        </div>
    )
    useEffect(()=>{
        if(!user["authReady"]){
            router.push("/login")
            layout = (<div className={"unauth"}>{children}</div>)
        }
    },[])
    return(
        !user["authReady"] ?<div className={"layout"}>{children}</div>:<div className={"unauth"}>{children}</div>
    )
}