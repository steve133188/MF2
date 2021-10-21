import  {AccountSettingPanel} from "../components/AccountSettingPanel"
import {BlueMenu} from "../components/BlueMenu";
import {BlueMenuLink} from "../components/BlueMenuLink";

export default function Layout_live_chat (){

    return(
        <div className="layout_setting">
            <AccountSettingPanel/>
            <BlueMenu>
                <BlueMenuLink link="">Account Setting</BlueMenuLink>
                <BlueMenuLink link="">Reset Password</BlueMenuLink>
            </BlueMenu>
        </div>
    )
}