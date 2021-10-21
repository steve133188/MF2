import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {
    return (
        <div className={styles.container}>
            <Head>
                <title>MatrixForce 2.0</title>
                <meta name="description" content="The best social commerce solution"/>
                <link rel="icon" href="/favicon.png"/>
                <link href='https://fonts.googleapis.com/css?family=Manrope' rel='stylesheet'/>
            </Head>
            <main>
                <div id={"dashboard"}>

                </div>
            </main>
        </div>
    )
}
