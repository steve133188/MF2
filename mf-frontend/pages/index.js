import Head from 'next/head'
import Image from 'next/image'
<<<<<<< HEAD
import styles from '../styles/Home.module.css'
import Breadcrumb from '../components/Common/Breadcrumb'
import {CardA} from '../components/Common/Card'

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


          <CardA url="/" />




          {/* Write code here*/}
        </div>
      </main>
=======
import styles from '.././styles/Home.module.css'
import {Card, Card_horizontal, Card_colored, Card_outline, Card_groups} from "../components/Common/Card";


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
                    Get started by editing{''}
                    <code className={styles.code}>pages/index.js</code>
                </p>
>>>>>>> 4a5bdd37ec6f1b622aee72be57823565cc3d9b45

                <div className={styles.grid}>

                    {/* Write code here*/}

                    {/* Write code here*/}

                </div>
            </main>
        </div>
    )
}
