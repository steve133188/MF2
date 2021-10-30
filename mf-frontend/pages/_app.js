import '../styles/globals.scss'
import Layout from "../layouts/Layout";
import {AuthContextProvider} from "../context/authContext";
import {useEffect} from "react";
import {client} from "../services/websocket";
import Head from "next/head";

function MyApp({Component, pageProps}) {
    useEffect(() => {
        client.onopen = () => {
            console.log('WebSocket Client Connected');
        };
        client.onmessage = (message) => {
            console.log(message);
        };
    }, [])
    return (
        <>
            <Head>
                <title>MatrixForce 2.0</title>
                <meta name="description" content="The best social commerce solution"/>
                <link rel="icon" href="/MS_logo-square.svg"/>
                <link href='https://fonts.googleapis.com/css?family=Manrope' rel='stylesheet'/>
            </Head>
            <AuthContextProvider>
                <Layout>
                    <Component {...pageProps} />
                </Layout>
            </AuthContextProvider>
        </>
    )
}

export default MyApp
