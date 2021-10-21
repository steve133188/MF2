import '../styles/globals.scss'
import Layout from "../layouts/Layout";
import {AuthContextProvider} from "../context/authContext";


function MyApp({ Component, pageProps }) {
  return(
    <AuthContextProvider>
      <Layout>
      <Component {...pageProps} />
      </Layout>
    </AuthContextProvider>
)
}

export default MyApp
