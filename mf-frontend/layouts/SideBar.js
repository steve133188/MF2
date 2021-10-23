import Link from "next/link";

export default function SideBar({navItems}){
    const navs =navItems
    const dropdown = ()=>{

    }
    return(
        <div className={"layout-sidebar"}>
            <div className={"brand-logo"}>
                <img src="/MS_logo-square (1).svg" alt="MatrixForce"/>
            </div>
            <div className={"nav-items"}>
                <Link href={"/dashboard"} >
                    <div className={"nav-item active"}><img src="" alt=""/> Dashboard</div>
                </Link>
                <Link href={"/livechat"} >
                    <div className={"nav-item "}><img src="" alt=""/> Live Chat</div>
                </Link>
                <Link href={"/contacts"} >
                    <div className={"nav-item "}><img src="" alt=""/> Contacts</div>
                </Link>
                <Link href={"/broadcast"} >
                    <div className={"nav-item "}><img src="" alt=""/> Broadcast</div>
                </Link>
                <Link href={"/flowbuilder"} >
                    <div className={"nav-item "}><img src="" alt=""/> Flow Builder</div>
                </Link>
                <Link href={"/products"} >
                    <div className={"nav-item "}><img src="" alt=""/> Product</div>
                </Link>
                <Link href={"/organization"} >
                    <div className={"nav-item "}><img src="" alt=""/> Organization</div>
                </Link>
                <Link href={"/admin"} >
                    <div className={"nav-item "}><img src="" alt=""/> Admin</div>
                </Link>
                {/*{navItems.map((i,index)=>{*/}
                {/*    <Link key={index} href={i.url}>*/}
                {/*        <div className={"nav-item"}><img src={i.icon} alt={i.name}/> {i.name} </div>*/}
                {/*    </Link>*/}
                {/*})}*/}
            </div>
        </div>
    )
}