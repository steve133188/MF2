import Head from 'next/head'
import Image from 'next/image'
import React from "react";
import {LoginPanel} from "../../components/LoginPanel";
import {ForgetPasswordPanel} from "../../components/ForgetPasswordPanel";


export default function Login(){
    return(
        <div className={"login-layout"}>
            <LoginPanel />
        </div>
    )
}