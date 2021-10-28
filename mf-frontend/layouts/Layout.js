import {navItems} from "./nav-item";
import {useRouter} from "next/router";
import {useEffect, useContext, useState} from "react";
import {AuthContext} from "../context/authContext"
import SideBar from "./SideBar";
import {MultipleSelectPlaceholder, SingleSelect, SingleSelect2} from "../components/Select";
import * as React from "react";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";

export default function Layout({children}) {
    const router = useRouter()
    const {user, login} = useContext(AuthContext)


    let layout = (
        <div className={"layout"}><SideBar navItems={navItems} />
            <div className={"layout-main"}><LayoutTop page_title={router.pathname.charAt(1).toUpperCase()+router.pathname.substring(2)}/>{children}</div>
        </div>
    )

    let unAuth = (<div className={"unauth"}>{children}</div>)
    useEffect(()=>{
        if(!user["authReady"]){
            router.push("/login")
            layout = (<div className={"unauth"}>{children}</div>)
        }
    },[])
    return (
        user["authReady"] ? layout : unAuth
    )
}

function LayoutTop(props) {
    const {user} = useContext(AuthContext)
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
                    <UserStatusSelect username={user.user}/>
                </div>
            </div>
        </div>
    )
}

function UserStatusSelect({username}) {
    const [age, setAge] = React.useState('');

    const handleChange = (event) => {
        setAge(event.target.value);
        console.log(event.target.value);
    };

    return (
        <FormControl sx={{m: 1, minWidth: 120}}>
            <Select
                value={age}
                onChange={handleChange}
                displayEmpty
                inputProps={{'aria-label': 'Without label'}}
                sx={{
                    width: 160,
                    height: 31,
                    borderRadius: 19,
                    background: "#F5F6F8",
                    border: "none",
                    textAlign: "center"
                }}
            >
                <MenuItem value="">
                    <div className={'selectStatusOnline'}></div>
                    <span>{username}</span></MenuItem>
                <MenuItem value={10}>
                    <div className={'selectStatusOnline'}></div>
                    <span>{username}</span></MenuItem>
                <MenuItem value={20}>
                    <div className={'selectStatusOffline'}></div>
                    <span>Twenty</span></MenuItem>
                <MenuItem value={30}>
                    <div className={'selectStatusOffline'}></div>
                    <span>Thirty</span></MenuItem>
            </Select>
        </FormControl>
    )
}