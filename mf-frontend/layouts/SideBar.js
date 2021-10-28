import Link from "next/link";
import {useRouter} from "next/router";

export default function SideBar(props){
    const router = useRouter()
    const {navItems} =props
    const dropdown = ()=>{

    }
    return(
        <div className={"layout-sidebar"}>
            <div className={"brand-logo"}>
                <img src="/MS_logo-square (1).svg" alt="MatrixForce"/>
            </div>
            <div className={"nav-items"}>
                <div className={router.pathname == "/dashboard" ? "active-side-item" : "side-item "}>
                <Link href={"/dashboard"} >
                    <div className={router.pathname == "/dashboard" ? "active nav-item" : "nav-item "}>Dashboard</div>
                </Link>
                </div>
                <div className={router.pathname == "/livechat" ? "active-side-item" : "side-item "}>
                <Link href={"/livechat"} >
                    <div className={router.pathname == "/livechat" ? "active nav-item" : "nav-item "}>  Live Chat</div>
                </Link>
                </div>
                <div className={router.pathname == "/contacts" ? "active-side-item" : "side-item "}>
                <Link href={"/contacts"} >
                    <div className={router.pathname == "/contacts" ? "active nav-item" : "nav-item "}>  Contacts</div>
                </Link>
                </div>
                <div className={router.pathname == "/broadcast" ? "active-side-item" : "side-item "}>
                <Link href={"/broadcast"} >
                    <div className={router.pathname == "/broadcast" ? "active nav-item" : "nav-item "}> Broadcast</div>
                </Link>
                </div>
                <div className={router.pathname == "/flowbuilder" ? "active-side-item" : "side-item "}>
                <Link href={"/flowbuilder"} >
                    <div className={router.pathname == "/flowbuilder" ? "active nav-item" : "nav-item "}>   Flow Builder</div>
                </Link>
                </div>
                <div className={router.pathname == "/integrations" ?"active-side-item" : "side-item "}>
                    <Link href={"/integrations"} >
                        <div className={router.pathname == "/integrations" ? "active nav-item" : "nav-item "}>  Integrations</div>
                    </Link>
                </div>
                <div className={router.pathname == "/products" ? "active-side-item" : "side-item "}>
                <Link href={"/products"} >
                    <div className={router.pathname == "/products" ? "active nav-item" : "nav-item "}>  Product</div>
                </Link>
                </div>
                <div className={router.pathname == "/organization" ?"active-side-item" : "side-item "}>
                <Link href={"/organization"} >
                    <div className={router.pathname == "/organization" ? "active nav-item" : "nav-item "}>  Organization</div>
                </Link>
                </div>
                <div className={router.pathname == "/admin" ? "active-side-item" : "side-item "}>
                <Link href={"/admin"} >
                    <div className={router.pathname == "/admin" ? "active nav-item" : "nav-item "}> Admin</div>
                </Link>
                </div>
                {/*{navItems.map((i,index)=>{*/}
                {/*    <Link key={index} href={i.url}>*/}
                {/*        <div className={"nav-item"}><img src={i.icon} alt={i.name}/> {i.name} </div>*/}
                {/*    </Link>*/}
                {/*})}*/}
            </div>
        </div>
    )
}

// export async function getStaticProps(context){
//     const res = await fetch(`"../data/nav.json"`)
//     const data = await res.json()
//
//     return{
//         props:{
//             data
//         }
//     }
// }