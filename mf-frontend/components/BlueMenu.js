import Link from "next/link"
import {BlueMenuDropdown, BlueMenuLink} from "./BlueMenuLink";
import {useState} from 'react';
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";

export function BlueMenu({children, ...props}) {
    const [isSelect, setSelect] = useState(false)

    function toggleIsSelect() {
        setSelect(!isSelect);
    }

    const [isShow, setShow] = useState(false)

    function toggleIsShow() {
        setShow(!isShow);
    }

    const [isShow2, setShow2] = useState(false)

    function toggleIsShow2() {
        setShow2(!isShow2);
    }

    return (
        <nav className="blueMenu">
            <ul className="blueMenuGroup">
                <li className="blueMenuLink blueLinkActive"><Link href=""><a className={"blueLink"}>Dashboard</a></Link></li>
                <li className="blueMenuLink"><span className="blueLink clickableSpan" onClick={toggleIsShow}>Features<KeyboardArrowDownIcon/></span>
                    {isShow ? (<ul className="blueMenuGroup blueMenuDropdownGroup">
                        <li className="blueMenuLink"><Link href=""><a className="blueLink">Page</a></Link></li>
                        <li className="blueMenuLink"><Link href=""><a className="blueLink">Element</a></Link></li>
                    </ul>):null}
                </li>
                <li className="blueMenuLink "><span className="blueLink clickableSpan" onClick={toggleIsShow2}>Services<KeyboardArrowDownIcon/></span>
                    {isShow2 ? ( <ul className="blueMenuGroup blueMenuDropdownGroup">
                        <li className="blueMenuLink"><Link href=""><a className="blueLink">App Design</a></Link></li>
                        <li className="blueMenuLink"><Link href=""><a className="blueLink">Web Design</a></Link></li>
                    </ul>):null}
                </li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Portfolio</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Overview</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Shortcuts</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Feedback</a></Link></li>
            </ul>
        </nav>
    )
}

export function BlueMenu2({children, ...props}) {
    const [isSelect, setSelect] = useState(false)

    function toggleIsSelect() {
        setSelect(!isSelect);
    }

    const [isShow, setShow] = useState(false)

    function toggleIsShow() {
        setShow(!isShow);
    }

    const [isShow2, setShow2] = useState(false)

    function toggleIsShow2() {
        setShow2(!isShow2);
    }

    return (
        <nav className="blueMenu">
            <ul className="blueMenuGroup">
                <li className="blueMenuLink"><span className="blueLink clickableSpan" onClick={toggleIsShow}>Agent & Terms<KeyboardArrowDownIcon/></span>
                    {isShow ? (<ul className="blueMenuGroup blueMenuDropdownGroup">
                        <li className="blueMenuLink"><Link href=""><a className="blueLink">Agent</a></Link></li>
                        <li className="blueMenuLink"><Link href=""><a className="blueLink">Role</a></Link></li>
                    </ul>):null}
                </li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Contact Groups</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Standard Reply</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Tags</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Assignment Role</a></Link></li>
                <li className="blueMenuLink"><Link href=""><a className={"blueLink"}>Message API</a></Link></li>
            </ul>
        </nav>
    )
}