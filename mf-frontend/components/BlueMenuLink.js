import Link from "next/link";
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';

export function BlueMenuLink ({ children ,...props}) {
    const { link ,title, onClick, isSelect, setSelect } = props
    const classname = isSelect?"blueLink menuSelected":"blueLink"
    return (
        <li className="blueMenuLink" onClick={onClick}><Link href={link}><a className={classname}>{title}</a></Link>{children}</li>
    )
}

export function BlueMenuDropdown ({ children ,...props} ) {
    const { link ,title } = props
    return (
        <li className="blueMenuLink "><Link href={link}><a className="blueLink">{title}<KeyboardArrowDownIcon/></a></Link>{children}</li>
    )
}

export function BlueMenuDropdownLink ({ children ,...props} ) {
    const { link ,name } = props
    return (
        <li className="blueMenuLink"><Link href={link}><a className="blueLink">{name}</a></Link>{children}</li>
    )
}