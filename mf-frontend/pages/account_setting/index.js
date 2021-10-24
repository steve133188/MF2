import {BlueMenu} from "../../components/BlueMenu";
import {AccountSettingPanel} from "../../components/AccountSettingPanel";
import {BlueMenuLink} from "../../components/BlueMenuLink";

export default function Account_setting() {
    return (
<<<<<<< HEAD

=======
        <>
>>>>>>> caec76d974697c3fd6f4dfbb9fc97ae936d2db12
        <div className="account_setting-layout">
            <div className="leftMenu">MENU</div>

            <div className="rightContent">
                <div className="navbar">NAVBAR</div>
                <div className="container">

                    <BlueMenu>
                        <BlueMenuLink link="">Account Setting</BlueMenuLink>
                        <BlueMenuLink link="">Reset Password</BlueMenuLink>
                    </BlueMenu>
                    <AccountSettingPanel/>
                </div>
            </div>
<<<<<<< HEAD
        </div>
=======
            </div>
        </>
>>>>>>> caec76d974697c3fd6f4dfbb9fc97ae936d2db12
    )
}