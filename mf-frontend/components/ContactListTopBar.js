import {Search2, Search3} from "./Input";
import {LabelSelect, TeamFilterSelect} from "./Select";
import React from "react";
import {useState} from "react";

export function ContactListTopBar({ isFilter , setIsFilter}) {
    // const [isFilter, setIsFilter] = useState(false);

    const filterToggle = () => {
        setIsFilter( !isFilter )
    }
    return (
        <div className="contactListTopBar">
            <Search3>Search</Search3>
            <div className="contactListFilterBar">
                <div className="contactListBtns">
                    <div className="teamFilterSelect">
                        <TeamFilterSelect/>
                    </div>
                    <span className={isFilter?"filterIcon filterIconActive":"filterIcon"} onClick={filterToggle}>
                        <img src="icon-filter.svg" width="18px" height="18px" alt=""/>
                    </span>
                </div>
            </div>
        </div>
    )
}

function contactInput() {

    return (
        <div>
            <input type="text" className={'chat-filter-input'}/>
        </div>
    )
}