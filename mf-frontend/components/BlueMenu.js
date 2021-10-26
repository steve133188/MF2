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

        // <nav className="blueMenu">
        //     <BlueMenuGroup>
        //         <BlueMenuLink link="" title="Dashboard" onClick={toggleIsSelect} isSelect={isSelect}
        //                       setSelect={setSelect}></BlueMenuLink>
        //         <BlueMenuDropdown link="" title="Features">
        //             <BlueMenuGroup name="blueMenuDropdownGroup">
        //                 {isShow ? (
        //                         <BlueMenuLink link="" title="Pages" onClick={toggleIsSelect}></BlueMenuLink>)
        //                     : null
        //                 }
        //                 {isShow ?
        //                     (<BlueMenuLink link="" title="Element" onClick={toggleIsSelect}></BlueMenuLink>)
        //                     : null
        //                 }
        //             </BlueMenuGroup>
        //         </BlueMenuDropdown>
        //         <BlueMenuDropdown link="" title="Services">
        //             <BlueMenuGroup name="blueMenuDropdownGroup">
        //                 <BlueMenuLink link="" title="App Design" onClick={toggleIsSelect}></BlueMenuLink>
        //                 <BlueMenuLink link="" title="Web Design" onClick={toggleIsSelect}></BlueMenuLink>
        //             </BlueMenuGroup>
        //         </BlueMenuDropdown>
        //         <BlueMenuLink link="" title="Portfolio" onClick={toggleIsSelect}></BlueMenuLink>
        //         <BlueMenuLink link="" title="Overview" onClick={toggleIsSelect}></BlueMenuLink>
        //         <BlueMenuLink link="" title="Shortcuts" onClick={toggleIsSelect}></BlueMenuLink>
        //         <BlueMenuLink link="" title="Feedback" onClick={toggleIsSelect}></BlueMenuLink>
        //     </BlueMenuGroup>
        // </nav>
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

export function BlueMenuGroup({children, ...props}) {
    const {name, onClick} = props
    const classname = "blueMenuGroup " + name
    return (
        <ul className={classname} onClick={onClick}>
            {children}
        </ul>
    )
}
