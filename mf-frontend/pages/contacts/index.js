import Head from 'next/head'
import Image from 'next/image'
import {Search3} from "../../components/Input";
import {
    CancelButton,
    ToggleButton,
    SelectButton,
    NormalButton,
    NormalButton2,
    TextWithIconButton
} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {NormalTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";
import {useState} from "react";
import {TableItem} from "../../components/TableItem"
import Avatar from "@mui/material/Avatar";
import {Pill} from "../../components/Pill";
import {Checkbox1, SingleBox} from "../../components/Checkbox"
import {EditColumnPopper} from "../../components/EditColumnPopper";
import {MultipleSelectPlaceholder} from "../../components/Select";
import {AddPopper} from "../../components/AddPopper";
import {DeletePopper} from "../../components/DeletePopper";
import * as React from "react";

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
        <div className="contacts-layout">


            <div className="rightContent">
                <div className="contactsContainer">
                    <div className="topBar">
                        <div className="searchBar">
                            <Search3 type="search">Search</Search3>
                        </div>
                        <div className="buttonGrp">
                            {isSelectRow ? (
                                <span onClick={toggleSelectRow}><CancelButton/></span>
                            ) : (
                                <span onClick={toggleSelectRow}><SelectButton/></span>
                            )}
                            <EditColumnPopper/>
                            <NormalButton>Import</NormalButton>
                            <NormalButton2>+ New Contact</NormalButton2>
                        </div>
                    </div>
                    <div className="navbarPurple">
                        <div className="selectButtonGroup">
                            <MultipleSelectPlaceholder/>
                            <MultipleSelectPlaceholder/>
                            <MultipleSelectPlaceholder/>
                            <MultipleSelectPlaceholder/>
                        </div>

                        <div className="tagsButtonGroup">
                            <AddPopper/>

                            <div>
                                <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="#2198fa" cursor="pointer"
                                     className="bi bi-upload" viewBox="0 0 16 16">
                                    <path
                                        d="M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5z"/>
                                    <path
                                        d="M7.646 1.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708L8.5 2.707V11.5a.5.5 0 0 1-1 0V2.707L5.354 4.854a.5.5 0 1 1-.708-.708l3-3z"/>
                                </svg>
                            </div>

                            <DeletePopper/>
                        </div>
                    </div>
                    <NormalTable classname={isSelectRow ? null : "checkBox"}>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                    </NormalTable>
                    <PaginationControlled/>
                </div>
            </div>

        </div>
    )
}