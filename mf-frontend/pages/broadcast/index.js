import Head from 'next/head'
import Image from 'next/image'
import {Search3} from "../../components/Input";
import {NormalButton, NormalButton2, TextWithIconButton, CancelButton} from "../../components/Button";
import {PaginationControlled} from "../../components/Pagination";
import * as React from 'react';
import {LabelSelect, MultipleSelectPlaceholder} from "../../components/Select";
import {Badge} from "../../components/Badge";
import {StatusPill} from "../../components/Pill";

export default function Broadcast() {
    return (
        <div className="broadcast-layout">
            <div className="rightContent">
                <div className="broadcastContainer">
                    <div className="topBar">
                        <div className="searchBar">
                            <Search3 type="search">Search</Search3>
                        </div>
                        <div className="buttonGrp">
                            <NormalButton>Select</NormalButton>
                            <NormalButton2>+ New Broadcast</NormalButton2>
                        </div>
                    </div>
                    <div className="navbarPurple">
                        <div className="selectButtonGroup">
                            <MultipleSelectPlaceholder placeholder={"Agent"} />
                            <MultipleSelectPlaceholder placeholder="Team" />
                            <MultipleSelectPlaceholder placeholder="Tags" />
                            <MultipleSelectPlaceholder placeholder="Channel" />
                        </div>
                    </div>
                    <div className="broadcastTable">
                        <table className="table">
                            <thead>
                            <tr className="headTr">
                                <th className="trID">Name</th>
                                <th>Period</th>
                                <th>Group</th>
                                <th>Status</th>
                                <th>Created By</th>
                                <th>Created Date</th>
                                <th>Description</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr className="bodyTr">
                                <td>Broadcast 1</td>
                                <td><div><div>Sep 30, 2021 7:00 AM</div><span style={{marginLeft: "65px"}}>-</span><div>Oct 30, 2021 7:00 AM</div></div></td>
                                <td><Badge color="gp1">Group1</Badge></td>
                                <td><StatusPill color="statusActive">Active</StatusPill></td>
                                <td>Lorem Ipsum</td>
                                <td>Mary Foster</td>
                                <td>s</td>
                            </tr>
                            <tr>
                                <td>Broadcast 2</td>
                                <td><div><div>Sep 30, 2021 7:00 AM</div><span style={{marginLeft: "65px"}}>-</span><div>Oct 30, 2021 7:00 AM</div></div></td>
                                <td><Badge color="gp2">Group2</Badge></td>
                                <td><StatusPill color="statusPending">Pending</StatusPill></td>
                                <td>s</td>
                                <td>s</td>
                                <td>Mary Foster</td>
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