import Head from 'next/head'
import Image from 'next/image'
import {Alert, AlertTitle} from "@mui/material";
import {Search2} from "../../components/Input";
import {Switch} from "../../components/Switch";
import Link from "next/link";
import {NormalButton2} from "../../components/Button";
import {ForgetPasswordPanel} from "../../components/ForgetPasswordPanel";
import {SuccessPanel} from "../../components/SuccessPanel"
import {useState} from "react";

export default function Recovery() {
    const [email , setEmail] = useState("")
    const [isSubmit , setIsSubmit] = useState(false)
    function validateForm() {
        return email.length > 0;
    }
    let p = <SuccessPanel/>
    let pp = (<div className="forgetPasswordPanel">
        <div className="companyLogo">
            <img src="MS_logo-square.svg" alt=""/>
        </div>
        <div className="mainContent">
            <div className="welcomeMessage">
                <h1>Forget Password</h1>
                <p>Enter the email address associated with <br/>your account and we’ll send you a link <br/> to reset your password.</p>
            </div>
            <div className="inputSet">
                <Search2 handleChange={(e)=> {
                    setEmail(e.target.value);
                    console.log(email)
                }}  type="text">Email</Search2>
            </div>
            <NormalButton2 onClick={()=>setIsSubmit(true)} disabled={!validateForm()}>
                Request Reset
            </NormalButton2>
        </div>
    </div>)
    const wrapper =(text)=>{
        return<div><h1>{text}</h1></div>
    }

    return (
        <div className={"login-layout "}>
            <div className={"container"}>
                {isSubmit ? p : pp}
            </div>
        </div>
    )
}