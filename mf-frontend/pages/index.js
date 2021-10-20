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

            <main className={styles.main}>
                <h1 className={styles.title}>
                    Welcome to <a href="https://nextjs.org">Next.js!</a>
                </h1>

                <p className={styles.description}>
                    Get started by editing{' '}
                    <code className={styles.code}>pages/index.js</code>
                </p>


            </main>

            <footer className={styles.footer}>
                <a
                    href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    Powered by{' '}
                    <span className={styles.logo}>
            <Image src="/logo.png" alt=" Logo" width={16} height={16}/>
          </span>
                </a>
            </footer>
        </div>
    )
}
