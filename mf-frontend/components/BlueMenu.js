import Link from "next/link"
import {BlueMenuDropdown, BlueMenuLink} from "./BlueMenuLink";
import { useState } from 'react';

export function BlueMenu({children,...props}) {
    const [isSelect, setSelect] = useState(false)

    function toggleIsSelect() {
        setSelect(!isSelect);
    }

    return (
        <nav className="blueMenu">
            <BlueMenuGroup>
                <BlueMenuLink link="" title="Dashboard" onClick={toggleIsSelect} isSelect={isSelect} setSelect={setSelect}></BlueMenuLink>
                <BlueMenuDropdown link="" title="Features">
                    <BlueMenuGroup name="blueMenuDropdownGroup" onClick="">
                        <BlueMenuLink link="" title="Pages" onClick={toggleIsSelect}></BlueMenuLink>
                        <BlueMenuLink link="" title="Element" onClick={toggleIsSelect}></BlueMenuLink>
                    </BlueMenuGroup>
                </BlueMenuDropdown>
                <BlueMenuDropdown link="" title="Services">
                    <BlueMenuGroup name="blueMenuDropdownGroup">
                        <BlueMenuLink link="" title="App Design" onClick={toggleIsSelect}></BlueMenuLink>
                        <BlueMenuLink link="" title="Web Design" onClick={toggleIsSelect}></BlueMenuLink>
                    </BlueMenuGroup>
                </BlueMenuDropdown>
                <BlueMenuLink link="" title="Portfolio" onClick={toggleIsSelect}></BlueMenuLink>
                <BlueMenuLink link="" title="Overview" onClick={toggleIsSelect}></BlueMenuLink>
                <BlueMenuLink link="" title="Shortcuts" onClick={toggleIsSelect}></BlueMenuLink>
                <BlueMenuLink link="" title="Feedback" onClick={toggleIsSelect}></BlueMenuLink>
            </BlueMenuGroup>
        </nav>
    )
}

export function BlueMenuGroup({ children,...props}) {
    const {name} = props
    const classname = "blueMenuGroup " + name
    return (
        <ul className={classname}>
            {children}
        </ul>
    )
}
