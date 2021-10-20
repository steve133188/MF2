import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>New Matrixforce</title>
        <meta name="description" content="Matrixforce 2.0 beta" />
        {/*<link rel="icon" href="/favicon.ico" />*/}
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          UI Test page
        </h1>

        <p className={styles.description}>
          Get started by editing{' '}
          <code className={styles.code}>pages/index.js</code>
        </p>

        <div className={styles.grid}>
        {/* Write code here*/}

          {/* Write code here*/}
        </div>
      </main>


    </div>
  )
}
