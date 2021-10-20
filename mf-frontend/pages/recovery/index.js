import Head from 'next/head'
import Image from 'next/image'
import {Alert, AlertTitle} from "@mui/material";
import {Search2} from "../../components/Input";
import {Switch} from "../../components/Switch";
import Link from "next/link";
import {NormalButton2} from "../../components/Button";
import {ForgetPasswordPanel} from "../../components/ForgetPasswordPanel";
import {SuccessPanel} from "../../components/SuccessPanel"

export default function Recovery() {
    return (
        <>
            <ForgetPasswordPanel/>
            <SuccessPanel/>
        </>
    )
}