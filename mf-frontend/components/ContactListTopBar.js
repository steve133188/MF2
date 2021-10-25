import {Search2} from "./Input";
import {LabelSelect} from "./Select";
import React from "react";

export function ContactListTopBar() {
    return(
        <div className="contactListTopBar">
            <div className="contactListFilterBar">

                <div className="contactListBtns">
                    <LabelSelect/>
                    <span>
                        <img className="type"
                               src="data:image/png;base64,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"
                               alt=""/>
                    </span>
                </div>
            </div>
        </div>
    )
}

function contactInput(){

    return(
        <div>
            <input type="text" className={'chat-filter-input'}/>
        </div>
    )
}