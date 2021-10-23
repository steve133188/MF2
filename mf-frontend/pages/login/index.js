import Head from 'next/head'
import Image from 'next/image'
import React, {useContext, useEffect} from "react";
import {LoginPanel} from "../../components/LoginPanel";
import {ForgetPasswordPanel} from "../../components/ForgetPasswordPanel";

import {useRouter} from "next/router";
import {AuthContext} from "../../context/authContext";


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