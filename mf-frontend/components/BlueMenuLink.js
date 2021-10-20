import Link from "next/link";

export function BlueMenuLink (props) {
    return (
        <Link href={props.link}><a><li className="blueMenuLink">{props.children}</li></a></Link>
    )
}

export function BlueMenuDropdown (props) {
    return (
        <Link href={props.link}><a><li className="blueMenuLink blueMenuLinkActive">{props.children}</li></a></Link>
    )
}