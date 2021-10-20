import Head from 'next/head'
import Image from 'next/image'
import {Search2} from "../../components/Input";
import {NormalButton, NormalButton2, TextWithIconButton} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {BroadcastTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";

export default function Broadcast() {
    return (
        <div className="broadcastContainer">
            <div className="topBar">
                <div className="searchBar">
                    <Search2 />
                </div>
                <div className="buttonGrp">
                    <NormalButton>Select</NormalButton>
                    <NormalButton2>+ New Boardcast</NormalButton2>
                </div>
            </div>
            <NavbarPurple />
            <BroadcastTable />
            <PaginationControlled />
        </div>
    )
}