import {LabelSelect, LabelSelect2, TeamFilterSelect} from "./Select";
import {Contacts} from "./Contacts";
import {Checkbox3, CheckboxGroup1, CheckboxGroup2, Search2, Search3} from "./Input"
import {ContactListTopBar} from "./ContactListTopBar"
import {useState} from "react";
import {ContactFilterList} from "./ContactFilterList";
import {Pill, PillInIcon} from "../components/Pill"
import {ContactType} from "./ContactType";
import {IconWithPill} from "./Icon";
import {Draggable} from "react-beautiful-dnd";
import * as React from "react";
import Select from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import {Checkbox1, Checkbox2} from "./Checkbox";
import {CancelButton, NormalButton2} from "./Button";

export function ContactList() {
    const [isFavourite, setFavourite] = useState(false)

    function addToFavourite() {
        setFavourite(!isFavourite);
    }

    const [isFilter, setIsFilter] = useState(false);
    const contacts = [
        {
            avatar: "https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg",
            messageAmount: 10,
            contactName: "Debra Patel1",
            team: "Team A",
            lastOnlineTime: "03:45 PM",
            contactExtraTime: "00:00:00",
            contactType: "/Group 4965.svg",
            pillStyle: "pill teamA"
        },
        {
            avatar: "https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg",
            messageAmount: 2,
            contactName: "Ddsfdsf",
            team: "Team B",
            lastOnlineTime: "01:55 PM",
            contactExtraTime: "00:00:00",
            contactType: "/Group 4965.svg",
            pillStyle: "pill teamB"
        },
        {
            avatar: "https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg",
            messageAmount: 10,
            contactName: "Debra Patel2",
            team: "Team C",
            lastOnlineTime: "03:45 PM",
            contactExtraTime: "00:00:00",
            contactType: "/Group 4965.svg",
            pillStyle: "pill teamC"
        },
        {
            avatar: "https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg",
            messageAmount: 10,
            contactName: "Debra Patel3",
            team: "Team D",
            lastOnlineTime: "03:45 PM",
            contactExtraTime: "00:00:00",
            contactType: "/Group 4965.svg",
            pillStyle: "pill teamD"
        },
        {
            avatar: "https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg",
            messageAmount: 10,
            contactName: "Debra Patel4",
            team: "Team E",
            lastOnlineTime: "03:45 PM",
            contactType: "/Group 4965.svg",
            contactExtraTime: "00:00:00",
            pillStyle: "pill teamE"
        },
    ];
    const filterToggle = () => {
        setIsFilter(!isFilter)
    }
    const [age, setAge] = React.useState('');
    const handleFilterChange = (event) => {
        setAge(event.target.value);
        console.log(event.target.value);
    };
    return (
        <div className="contactList">
            <div className="contactListTopBar">
                <div className="logoInputContainer2">
                    <span className="requiredMessage">This field is required.</span>
                    <label className="searchSVG">
                        <input placeholder="Search"/>
                    </label>
                </div>
                <div className="contactListFilterBar">
                    <div className="contactListBtns">
                        <div className="teamFilterSelect">

                            <FormControl sx={{m: 1, minWidth: 120}}>
                                <Select
                                    value={age}
                                    onChange={handleFilterChange}
                                    displayEmpty
                                    inputProps={{'aria-label': 'Without label'}}
                                    sx={{
                                        width: 160,
                                        height: 31,
                                        borderRadius: 19,
                                        background: "#D0E9FF",
                                        border: "none",
                                        textAlign: "center",
                                        color: "black",
                                        justifyContent: "space-between",
                                    }}
                                >
                                    <MenuItem value="" sx={{
                                        color: "#2198FA",
                                        justifyContent: "space-between",
                                        marginLeft: "0px"
                                    }}>
                                        <span>Mary Foster</span>
                                        <div className={"smallPill"}>20</div>
                                    </MenuItem>
                                    <MenuItem value={10} sx={{
                                        color: "#2198FA",
                                        justifyContent: "space-between",
                                        marginLeft: "0px"
                                    }}>

                                        <span>Ten</span>
                                        <div className={"smallPill"}>20</div>
                                    </MenuItem>
                                    <MenuItem value={20} sx={{
                                        color: "#2198FA",
                                        justifyContent: "space-between",
                                        marginLeft: "0px"
                                    }}>

                                        <span>Twenty</span>
                                        <div className={"smallPill"}>20</div>
                                    </MenuItem>
                                    <MenuItem value={30} sx={{
                                        color: "#2198FA",
                                        justifyContent: "space-between",
                                        marginLeft: "0px"
                                    }}>

                                        <span>Thirty</span>
                                        <div className={"smallPill"}>20</div>
                                    </MenuItem>
                                </Select>
                            </FormControl>
                            {/*    */}
                        </div>
                        <span className={isFilter ? "filterIcon filterIconActive" : "filterIcon"}
                              onClick={filterToggle}>
                        <img src="icon-filter.svg" width="18px" height="18px" alt=""/>
                    </span>
                    </div>
                </div>
            </div>


            {isFilter ?

                (<div className="filterArea">
                    <CheckboxGroup1 title="Filter">
                        <Checkbox1 checked="checked">Unread</Checkbox1>
                        <Checkbox1>Unassigned</Checkbox1>
                        <Checkbox1>ChatBot</Checkbox1>
                    </CheckboxGroup1>
                    <CheckboxGroup2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            All Channel
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png"
                                   checked="checked">
                            WhatsApp
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            WhatsApp Business API
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            Messager
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            WeChat
                        </Checkbox2>
                    </CheckboxGroup2>
                    <div className="agentFilter">
                        <p>Agent</p>
                        <LabelSelect2/>
                        <div className="agentGroup">
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightBlue" size="size30">MF</Pill>
                        </div>
                    </div>
                    <div className="tagFilter">
                        <p>Tag</p>
                        <div className="tagGroup">
                            <Pill color="vip">VIP</Pill>
                            <Checkbox3/>
                        </div>
                    </div>
                    <div className="buttonGrp">
                        <NormalButton2>Confirm</ NormalButton2>
                        <CancelButton>Cancel</ CancelButton>
                    </div>
                </div>)
                : (<div className="contactContainerGrp">

                    {contacts.map(({
                                       avatar,
                                       messageAmount,
                                       contactName,
                                       team,
                                       lastOnlineTime,
                                       contactExtraTime,
                                       contactType,
                                       pillStyle
                                   }) => {
                        return (
                            <div className="contactContainer" key={contactName}>
                                <div className="contact">
                                    {isFavourite ? (
                                        <svg id={"fillStarColor"} onClick={addToFavourite}
                                             xmlns="http://www.w3.org/2000/svg" width="16"
                                             height="16" fill="currentColor"
                                             className="bi bi-star-fill" viewBox="0 0 16 16">
                                            <path
                                                d="M3.612 15.443c-.386.198-.824-.149-.746-.592l.83-4.73L.173 6.765c-.329-.314-.158-.888.283-.95l4.898-.696L7.538.792c.197-.39.73-.39.927 0l2.184 4.327 4.898.696c.441.062.612.636.282.95l-3.522 3.356.83 4.73c.078.443-.36.79-.746.592L8 13.187l-4.389 2.256z"/>
                                        </svg>) : (
                                        <svg id={"addToFavouriteBtn"} onClick={addToFavourite}
                                             xmlns="http://www.w3.org/2000/svg" width="16"
                                             height="16"
                                             fill="currentColor"
                                             className="bi bi-star" viewBox="0 0 16 16">
                                            <path
                                                d="M2.866 14.85c-.078.444.36.791.746.593l4.39-2.256 4.389 2.256c.386.198.824-.149.746-.592l-.83-4.73 3.522-3.356c.33-.314.16-.888-.282-.95l-4.898-.696L8.465.792a.513.513 0 0 0-.927 0L5.354 5.12l-4.898.696c-.441.062-.612.636-.283.95l3.523 3.356-.83 4.73zm4.905-2.767-3.686 1.894.694-3.957a.565.565 0 0 0-.163-.505L1.71 6.745l4.052-.576a.525.525 0 0 0 .393-.288L8 2.223l1.847 3.658a.525.525 0 0 0 .393.288l4.052.575-2.906 2.77a.565.565 0 0 0-.163.506l.694 3.957-3.686-1.894a.503.503 0 0 0-.461 0z"/>
                                        </svg>)}

                                    <div className="contactInfoAndTime">
                                        <div className="contactInfoGrp">
                                            <span className="iconPillContainer">
                                                <img className="iconWithPill" src={avatar} alt=""/>
                                                <span className="pillInIcon">{messageAmount}</span>
                                            </span>
                                            <div className="contactChatInfo">
                                                <div className="contactName">{contactName}
                                                    <img className="contactType" src={contactType} alt=""/>
                                                </div>
                                                <span className="pillContainer">
                                                <span className={pillStyle}>{team}</span>
                                            </span>
                                            </div>
                                        </div>
                                        <div className="msgTimeGrp">
                                            <div className="lastOnlineTime">{lastOnlineTime}</div>
                                            <div className="contactExtraTime">{contactExtraTime}</div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        );
                    })}
                </div>)
            }
        </div>
    )
}