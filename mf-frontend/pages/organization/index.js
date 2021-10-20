import {BlueMenu} from "../../components/BlueMenu";
import {BlueMenuDropdown, BlueMenuLink} from "../../components/BlueMenuLink";

export default function Organization() {
    return (
        <BlueMenu>
            <BlueMenuLink link="#">Account Setting</BlueMenuLink>
            <BlueMenuLink link="#">Reset Password</BlueMenuLink>
            <BlueMenuDropdown link="#">Dropdown</BlueMenuDropdown>
        </BlueMenu>

    )
}