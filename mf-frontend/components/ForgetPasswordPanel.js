import Link from "next/link"
import {Search2} from "./Input"
import {NormalButton2} from "./Button"
import {Switch} from "./Switch";

export function ForgetPasswordPanel() {
    return (
        <div className="container">
            <div className="forgetPasswordPanel">
                <div className="companyLogo">
                    <img src="MS_logo-square.svg" alt=""/>
                </div>
                <div className="mainContent">
                    <div className="welcomeMessage">
                        <h1>Forget Paswword</h1>
                        <p>Enter the email address associated with <br/>your account and we’ll send you a link <br/> to reset your password.</p>
                    </div>
                    <div className="inputSet">
                        <Search2 type="text">Email or Username</Search2>
                    </div>
                    <NormalButton2>
                        Request Reset
                    </NormalButton2>
                </div>
            </div>
        </div>
    )
}