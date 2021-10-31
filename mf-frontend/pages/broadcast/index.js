import Head from 'next/head'
import Image from 'next/image'
import {Search3} from "../../components/Input";
import {NormalButton, NormalButton2, TextWithIconButton} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {BroadcastTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";

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
                    <NavbarPurple/>
                    <BroadcastTable/>
                    <PaginationControlled/>
                </div>
            </div>
        </div>
    )
}