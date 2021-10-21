import Head from 'next/head'
import Image from 'next/image'
import React, {useContext, useEffect} from "react";
import {LoginPanel} from "../../components/LoginPanel";
<<<<<<< HEAD
import {ForgetPasswordPanel} from "../../components/ForgetPasswordPanel";
=======
import {useRouter} from "next/router";
import {AuthContext} from "../../context/authContext";
>>>>>>> 863d9e42ff766350b88638972936682852a0b635


export default function Login(){
    const router = useRouter()
    const {user , login} = useContext(AuthContext)
    useEffect(()=>{
        if(user["authReady"]){
            router.back()
        }
    },[])
    return(
        <div className={"login-layout"}>
            <LoginPanel />
        </div>
    )
}