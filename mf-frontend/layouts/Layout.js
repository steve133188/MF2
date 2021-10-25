import {navItems} from "./nav-item";
import {useRouter} from "next/router";
import {useEffect, useContext, useState} from "react";
import {AuthContext} from "../context/authContext"
import SideBar from "./SideBar";
import {MultipleSelectPlaceholder, SingleSelect, SingleSelect2} from "../components/Select";
import * as React from "react";

export default function Layout({children}) {
    const router = useRouter()
    const {user, login} = useContext(AuthContext)
    let layout = (
        <div className={"layout"}><SideBar navItems={navItems}/>
            <div className={"layout-main"}><LayoutTop page_title={"Dashboard"}/>{children}</div>
        </div>
    )
    let unAuth = (<div className={"unauth"}>{children}</div>)
    // useEffect(()=>{
    //     if(!user["authReady"]){
    //         router.push("/login")
    //         layout = (<div className={"unauth"}>{children}</div>)
    //     }
    // },[])
    return (
        !user["authReady"] ? layout : unAuth
    )
}

function LayoutTop(props) {
    const [isMenuShow, setIsMenuShow] = useState(false)
    const [isNotificationShow, setIsNotificationShow] = useState(false)
    const {page_title} = props

    return (
        <div className={'navbar container-fluid'}>
            <div className={"page-title"}> {page_title}</div>
            <div className={'d-flex user-session'}>
                <div className={'notificationDropdownSet'}>
                    <div className={'badge'}> 10</div>
                    <div className="notificationDropdown">
                        <SingleSelect2/>
                    </div>
                </div>
                <div className="loggingStatusDropdown">
                    <SingleSelect/>
                </div>
            </div>
        </div>
    )
}