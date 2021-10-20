import Head from 'next/head'
import Image from 'next/image'
import {Search3} from "../../components/Input";
import {CancelButton, ToggleButton, SelectButton, NormalButton, NormalButton2, TextWithIconButton} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {NormalTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";
import {useState} from "react";
import {TableItem} from "../../components/TableItem"
import Avatar from "@mui/material/Avatar";
import {Pill} from "../../components/Pill";
import {Checkbox1, SingleBox} from "../../components/Checkbox"

export default function Contacts() {
    const [isSelectRow, setSelectRow] = useState(false);
    function toggleSelectRow() {
        setSelectRow(!isSelectRow);
    }
    const [isFillCheckbox, setFillCheckbox] = useState(false);
    function toggleFill() {
        setFillCheckbox(!isFillCheckbox);
    }


    return (

            <div className="contactsContainer">
                <div className="topBar">
                    <div className="searchBar">
                        <Search3 type="search">Search</Search3>
                    </div>
                    <div className="buttonGrp">
                        <span onClick={toggleSelectRow}><SelectButton/></span>
                        <TextWithIconButton>Edit Column</TextWithIconButton>
                        <NormalButton>Import</NormalButton>
                        <NormalButton2>+ New Contact</NormalButton2>
                    </div>
                </div>
                <NavbarPurple/>
                <NormalTable classname={isSelectRow ? null : "checkBox"}>
                    <tr>
                        <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                        <TableItem>0000001</TableItem>
                        <TableItem><div className="nameGroup"><Avatar alt="Remy Sharp" src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" />Debra Patel</div></TableItem>
                        <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                        <TableItem><div className="channel"><img src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png" alt=""/></div></TableItem>
                        <TableItem><div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New Customer</Pill></div></TableItem>
                        <TableItem><div className="assigneeGroup">
                            <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                            <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                            <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                            <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                            <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                        </div></TableItem>
                    </tr>
                    <tr>
                        <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                        <TableItem>0000001</TableItem>
                        <TableItem><div className="nameGroup"><Avatar alt="Remy Sharp" src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" />Debra Patel</div></TableItem>
                        <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                        <TableItem><div className="channel"><img src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png" alt=""/></div></TableItem>
                        <TableItem><div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New Customer</Pill></div></TableItem>
                        <TableItem><div className="assigneeGroup">
                            <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                            <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                            <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                            <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                            <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                        </div></TableItem>
                    </tr>

                </NormalTable>
                <PaginationControlled/>
            </div>

    )
}