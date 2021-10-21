// Cannot resize Switches

import React, {useState} from "react"

export function Input(props) {
    return (
        <div className="col-md-10">
            <input
                className="form-control"
                type={props.type}
                placeholder={props.value}
            />
        </div>
    )
}

export function Input2(props) {
    return (
        <div className="inputField">
            <span>{props.title}</span>
            <input type="text" placeholder={props.children} disabled={props.disabled}/>
        </div>
    )
}


export function Search(props) {
    return (
        <div className="position-relative">
            <input
                type="text"
                className="form-control"
                placeholder={props.placeholder}
            />
            <span className="uil-search"></span>
        </div>
    )
}

export function Search2(...props) {
    return (
        <div className="logoInputContainer">
            <span className="requiredMessage">This field is required.</span>
            <label className="searchSVG">
                <input type={props.type} placeholder={props.children} onChange={props.handleChange}/>
            </label>
        </div>
    )
}

export function Search3(props) {
    return (
        <div className="logoInputContainer2">
            <span className="requiredMessage">This field is required.</span>
            <label className="searchSVG">
                <input type={props.type} placeholder={props.children}/>
            </label>
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

export function TextRadio() {
    return (
        <ul className="textRadio">
            <li>
                <input type='radio' value='1' name='textRadio' id='textRadio1'/>
                <label htmlFor='textRadio1'>Value 1</label>
            </li>
            <li>
                <input type='radio' value='2' name='textRadio' id='textRadio2'/>
                <label htmlFor='textRadio2'>Value 2</label>
            </li>
            <li>
                <input type='radio' value='3' name='textRadio' id='textRadio3'/>
                <label htmlFor='textRadio3'>Value 3</label>
            </li>
        </ul>
    )
}

export function TextRadio2() {
    return (
        <ul className="textRadio2">
            <li>
                <input type='radio' value='1' name='textRadio' id='textRadio1'/>
                <label htmlFor='textRadio1'>Info</label>
            </li>
            <li>
                <input type='radio' value='2' name='textRadio' id='textRadio2'/>
                <label htmlFor='textRadio2'>Note</label>
            </li>
            <li>
                <input type='radio' value='3' name='textRadio' id='textRadio3'/>
                <label htmlFor='textRadio3'>File</label>
            </li>
        </ul>
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
                {label: "Mustard", value: "Mustard"},
                {label: "Ketchup", value: "Ketchup"},
                {label: "Relish", value: "Relish"}
            ]
        },
        {
            label: "Camping",
            options: [
                {label: "Tent", value: "Tent"},
                {label: "Flashlight", value: "Flashlight"},
                {label: "Toilet Paper", value: "Toilet Paper"}
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

export function ValidationInput() {
    return (
        <form className="row g-3 needs-validation" noValidate>
            <div className="col-md-8">
                <label htmlFor="validationCustom01" className="form-label">First name</label>
                <input type="text" className="form-control" id="validationCustom01" value="Mark" required/>
                <div className="valid-feedback">
                    Looks good!
                </div>
            </div>
            <div className="col-md-8">
                <label htmlFor="validationCustom03" className="form-label">City</label>
                <input type="text" className="form-control" id="validationCustom03" required/>
                <div className="invalid-feedback">
                    Please provide a valid city.
                </div>
            </div>

            <div className="col-12">
                <button className="btn btn-primary" type="submit">Submit form</button>
            </div>
        </form>
    )
}

export function CheckboxGroup1(props) {
    return (
        <div className="checkboxGroup1">
            <h1>{props.title}</h1>
            <div className="checkboxGrp">
                {props.children}
            </div>
        </div>
    )
}

export function CheckboxGroup2(props) {
    return (
        <div className="checkboxGroup2">
            <p>{props.title}</p>
            <div className="checkboxGrp">
                {props.children}
            </div>
        </div>
    )
}

export function Checkbox3() {
    return (
        <div className="checkbox3">
            <label className="checkboxContainer">
                <input type="checkbox" checked="checked"/>
                <span className="checkmark"></span>
            </label>
        </div>
    )
}

export function Checkbox() {
    return (
        <div></div>
    )

}

export function LogoTextField(props) {
    return (
        <div>
            <label className="LogoTextField">
                <input type="text" placeholder={props.children}/>
            </label>
        </div>
    )
}

export function LogInputField(props) {
    return (
        <label className="logInputField">
            <input type="text" className="logInput" placeholder={props.children}/>
            <input type="submit" className="logSubmit" value="Log"/>
        </label>
    )
}

export function LogoInputField(props) {
    return (
        <label className="logoInputField">
            <input type="text" className="logoInput" placeholder={props.children}/>
            <input type="submit" className="logoSubmit" value="Log"/>
        </label>
    )
}