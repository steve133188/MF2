import {navItems} from "./nav-item";
import {useRouter} from "next/router";
import {useEffect, useContext, useState} from "react";
import {AuthContext} from "../context/authContext"
import SideBar from "./SideBar";
export  default function Layout ({children}){
    const router = useRouter()
    const {user , login} = useContext(AuthContext)
    let layout =(
        <div className={"layout"}><SideBar navItems={navItems} /><div className={"layout-main"}><LayoutTop page_title={"Dashboard"} />{children}</div></div>
    )
    let unAuth = (<div className={"unauth"}>{children}</div>)
    // useEffect(()=>{
    //     if(!user["authReady"]){
    //         router.push("/login")
    //         layout = (<div className={"unauth"}>{children}</div>)
    //     }
    // },[])
    return(
        !user["authReady"] ?layout:unAuth
    )
}

function LayoutTop(props){
    const [isMenuShow , setIsMenuShow] = useState(false)
    const [isNotificationShow , setIsNotificationShow] = useState(false)
    const {page_title} = props

    return(
        <div className={'navbar container-fluid'}>
            <div className={"page-title"}> {page_title}</div>
            <div className={'d-flex user-session'}>
                <button className={"notification m-1"} onClick={()=>{setIsNotificationShow(!isNotificationShow);setIsMenuShow(false);console.log(isNotificationShow)}}>
                    <div className={'badge'}> 10</div>
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                         className="bi bi-bell" viewBox="0 0 16 16">
                        <path
                            d="M8 16a2 2 0 0 0 2-2H6a2 2 0 0 0 2 2zM8 1.918l-.797.161A4.002 4.002 0 0 0 4 6c0 .628-.134 2.197-.459 3.742-.16.767-.376 1.566-.663 2.258h10.244c-.287-.692-.502-1.49-.663-2.258C12.134 8.197 12 6.628 12 6a4.002 4.002 0 0 0-3.203-3.92L8 1.917zM14.22 12c.223.447.481.801.78 1H1c.299-.199.557-.553.78-1C2.68 10.2 3 6.88 3 6c0-2.42 1.72-4.44 4.005-4.901a1 1 0 1 1 1.99 0A5.002 5.002 0 0 1 13 6c0 .88.32 4.2 1.22 6z"/>
                    </svg>
                    {isNotificationShow ? (<div className="dropdown-menu notification-menu">
                       <button className={" dropdown-item "}>notifications</button>
                    </div>):null }
                </button>

                <div className={"user-dropdown m-1"}>
                    <button className={'dropdown-btn'} onClick={()=>{setIsMenuShow(!isMenuShow);setIsNotificationShow(false);console.log(isMenuShow)}}> <div className={'online'}></div>  username</button>
                {isMenuShow ? (<div className="dropdown-menu user-menu">
                    <button className={"dropdown-item"}><div className={"d-flex"}><div className={'offline'}></div> offline </div></button>
                    <button className={"dropdown-item"}><div className={"d-flex"}><div className={'offline'}></div> offline </div></button>
                    <button className={"dropdown-item"}><div className={"d-flex"}><div className={'offline'}></div> offline </div></button>
                    <button className={"dropdown-item"}><div className={"d-flex"}><div className={'offline'}></div> offline </div></button>
                </div>):null }
            </div>
            </div>
        </div>
    )
}