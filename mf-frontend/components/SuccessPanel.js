import Link from "next/link"
import {Search2} from "./Input"
import {NormalButton2} from "./Button"
import {Switch} from "./Switch";

export function SuccessPanel() {
    return (
        <div className="container">
            <div className="successPanel">
                <div className="companyLogo">
                    <img src="MS_logo-square.svg" alt=""/>
                </div>
                <div className="mainContent">
                    <div className="welcomeMessage">
                        <h1>success</h1>
                        <p>If you have an account with this email, <br/>You will receive an email with further instructions.</p>
                    </div>
                    <NormalButton2>
                        Confirm
                    </NormalButton2>
                </div>
            </div>
        </div>
    )
}