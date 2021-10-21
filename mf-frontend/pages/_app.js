import '../styles/globals.scss'
import Layout from "../layouts/Layout";
import {AuthContextProvider} from "../context/authContext";
import {useEffect} from "react";
import {client} from "../services/websocket";

function MyApp({ Component, pageProps }) {
    useEffect(()=>{
        client.onopen = () => {
            console.log('WebSocket Client Connected');
        };
        client.onmessage = (message) => {
            console.log(message);
        };
    },[])
  return(
    <AuthContextProvider>
      <Layout>
      <Component {...pageProps} />
      </Layout>
    </AuthContextProvider>
)
}

export default MyApp
