// Cannot resize Switches

import React, { useState } from "react"

export function Input(props) {
    return (
        <div className="col-md-10">
            <input
                className="form-control"
                type={props.type}
                defaultValue={props.value}
            />
        </div>
    )
}

export function Check() {
    return (
        <div>
            <div className="form-check mb-3">
                <input
                    className="form-check-input"
                    type="checkbox"
                    value=""
                    id="defaultCheck1"
                />
                <label
                    className="form-check-label"
                    htmlFor="defaultCheck1"
                >
                    Form Checkbox
                </label>
            </div>
            <div className="form-check form-check-end">
                <input
                    className="form-check-input"
                    type="checkbox"
                    value=""
                    id="defaultCheck2"
                    defaultChecked
                />
                <label
                    className="form-check-label"
                    htmlFor="defaultCheck2"
                >
                    Form Checkbox checked
                </label>
            </div>
        </div>
    )
}

export function Radio() {
    return (
        <>
            <div>
                <div className="form-check mb-3">
                    <input
                        className="form-check-input"
                        type="radio"
                        name="exampleRadios"
                        id="exampleRadios1"
                        value="option1"
                        defaultChecked
                    />
                    <label
                        className="form-check-label"
                        htmlFor="exampleRadios1"
                    >
                        Form Radio
                    </label>
                </div>
                <div className="form-check">
                    <input
                        className="form-check-input"
                        type="radio"
                        name="exampleRadios"
                        id="exampleRadios2"
                        value="option2"
                    />
                    <label
                        className="form-check-label"
                        htmlFor="exampleRadios2"
                    >
                        Form Radio checked
                    </label>
                </div>
            </div>
        </>
    )
}

export function Switches(props) {
    const [toggleSwitch, settoggleSwitch] = useState(true)
    // Size: sm md lg
    const size = "form-check form-switch form-switch-" + props.size + " mb-3";

    return (
        <div className={size}>
            <input
                type="checkbox"
                className="form-check-input"
                id="customSwitchsize"
                onClick={e => {
                    settoggleSwitch(!toggleSwitch)
                }}
            />
        </div>
    )
}
export function Select() {
    const optionGroup = [
        {
            label: "Picnic",
            options: [
                { label: "Mustard", value: "Mustard" },
                { label: "Ketchup", value: "Ketchup" },
                { label: "Relish", value: "Relish" }
            ]
        },
        {
            label: "Camping",
            options: [
                { label: "Tent", value: "Tent" },
                { label: "Flashlight", value: "Flashlight" },
                { label: "Toilet Paper", value: "Toilet Paper" }
            ]
        }
    ]
    const [selectedGroup, setselectedGroup] = useState(null)
    function handleSelectGroup(selectedGroup) {
        setselectedGroup(selectedGroup)
    }

    return (
        <div className="mb-3">
            <label>Single Select</label>
            <select className="form-select" aria-label="Default select example">
                <option selected>Open this select menu</option>
                <option value="1">One</option>
                <option value="2">Two</option>
                <option value="3">Three</option>
            </select>
        </div>
    )
}