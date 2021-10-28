import {BlueMenu} from "../../components/BlueMenu";
import {BlueMenuDropdown, BlueMenuLink} from "../../components/BlueMenuLink";
import {Search3} from "../../components/Input";
import {NormalButton, NormalButton2} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {BroadcastTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";
import {Badge} from "../../components/Badge";
import {StatusPill} from "../../components/Pill";
import * as React from "react";

export default function Organization() {
    return (
        <div className="organization-layout">
            <BlueMenu></BlueMenu>
            <div className="rightContent">
                <div className="broadcastContainer">
                    <div className="topBar">
                        <div className="searchBar">
                            <Search3 type="search">Search</Search3>
                        </div>
                        <div className="buttonGrp">
                            <NormalButton>Select</NormalButton>
                            <NormalButton2>+ New Team</NormalButton2>
                            <NormalButton2>+ New Division</NormalButton2>
                        </div>
                    </div>
                    <div className="navbarPurple">
                        <div>Team2   ( 8 )</div>
                        <NormalButton2>+ New Agent</NormalButton2>
                    </div>
                    <div className="broadcastTable">
                        <table className="table">
                            <thead>
                            <tr className="headTr">
                                <th className="trID">Name</th>
                                <th>Role</th>
                                <th>Email</th>
                                <th>Phone</th>
                                <th>No. Of Leads</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr className="bodyTr">
                                <td style={{display: "flex"}}>
                                    <div className="selectStatusOnline"></div>Harry Stewart
                                </td>
                                <td>Admin</td>
                                <td>Harry.stewart@gmail.com</td>
                                <td>+852 9765 0348</td>
                                <td>7</td>
                            </tr>
                            <tr>
                                <td style={{display: "flex"}}>
                                    <div className="selectStatusOffline"></div><span>Jasmine Miller</span></td>
                                <td>Agent</td>
                                <td>jasmine.miller@gmail.com</td>
                                <td>+852 9765 0348</td>
                                <td>6</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    <PaginationControlled/>
                </div>
            </div>
        </div>
    )
}