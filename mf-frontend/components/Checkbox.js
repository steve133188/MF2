import React from "react";
import {Pill} from "./Pill";

export function Checkbox1(props) {
    return (
        <>
            <label className="checkboxContainer">{props.children}
                <input type="checkbox" checked={props.checked}/>
                <span className="checkmark"></span>
            </label>
        </>
    )
}

export function Checkbox2(props) {
    return (
        <>
            <label className="checkboxContainer"><img
                src={props.src} alt=""/>{props.children}
                <input type="checkbox" checked={props.checked}/>
                <span className="checkmark"></span>
            </label>
        </>
    )
}

export function CheckboxPill(props) {
    return (
        <>
            <label className="checkboxContainer">
                <Pill color="lightBlue">VIP</Pill>
                <input type="checkbox" checked={props.checked}/>
                <span className="checkmark"></span>
            </label>
        </>
    )
}

export function SingleBox(props) {
    const classname = "checkmark " + props.fill
    return (
        <label className="testCheckboxContainer">
            <input type="checkbox" name={props.name} value={props.value} onClick={props.onClick} />
            <span className={classname}></span>
        </label>
    )
}
export function ColumnCheckbox(props) {
    return (
    <label className="columnCheckboxContainer"><img
        src="icon-columnControl.svg" alt=""/>{props.children}
        <input type="checkbox" checked={props.checked}/>
        <span className="checkmark"></span>
    </label>
    )
}

{/*
<label className="checkboxContainer">Unread
    <input type="checkbox" checked="checked"/>
    <span className="checkmark"></span>
</label>
<label className="checkboxContainer">Unassigned
    <input type="checkbox"/>
    <span className="checkmark"></span>
</label>
<label className="checkboxContainer">ChatBox
    <input type="checkbox"/>
    <span className="checkmark"></span>
</label>
*/
}