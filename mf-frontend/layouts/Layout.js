import {navItems} from "./nav-item";
import {useRouter} from "next/router";
import {useEffect, useContext, useState} from "react";
import {AuthContext} from "../context/authContext"
import SideBar from "./SideBar";
import {MultipleSelectPlaceholder, SingleSelect, SingleSelect2} from "../components/Select";
import * as React from "react";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";

export default function Layout({children}) {
    const [userSelect , setUserSelect] = useState("")
    const router = useRouter()
    const {user} = useContext(AuthContext)
    const { stateAuth, dispatchAuth } = useContext(AuthContext);
    let layout = (
        <div className={"layout"}><SideBar navItems={navItems} />
            <div className={"layout-main"}><div className={'navbar container-fluid'}>
                <div className={"page-title"}> {router.pathname.charAt(1).toUpperCase()+router.pathname.substring(2)}</div>
                <div className={'d-flex user-session'}>
                    <div className={'notificationDropdownSet'}>
                        {/*{user.userInfo.notification!=0? <div className={'badge'}> 10</div>: null}*/}
                        {user}
                        <div className="notificationDropdown">
                            <SingleSelect2/>
                        </div>
                    </div>
                    <div className="loggingStatusDropdown">
                        <FormControl sx={{m: 1, minWidth: 120}}>
                            <Select
                                value={userSelect}
                                // onChange={handleChange}
                                displayEmpty
                                label={`<div className={'selectStatusOnline'}></div>
                                    <span>{user.userInfo.username}</span></MenuItem>`}
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
                                <MenuItem value={""}>
                                    <div className={'selectStatusOnline'}></div>
                                    <span>{user.userInfo.username}</span></MenuItem>
                                <MenuItem value={"offine"}>
                                    <div className={'selectStatusOffline'}></div>
                                    <span>{user.userInfo.username}</span></MenuItem>
                                <MenuItem value={"other"}>
                                    <div className={''}></div>
                                    <span>User preference</span></MenuItem>
                                <MenuItem value={"other"}>
                                    <div className={''}></div>
                                    <span>Logout</span>
                                </MenuItem>
                            </Select>
                        </FormControl>
                    </div>
                </div>
            </div>{children}</div>
        </div>
    )

    let unAuth = (<div className={"unauth"}>{children}</div>)
    useEffect(()=>{
        if(!user.authReady){
            console.log("please log in")
            router.push("/login")
            layout = (<div className={"unauth"}>{children}</div>)
        }
    },[])
    return (
        user["authReady"] ? layout : unAuth
    )
}
