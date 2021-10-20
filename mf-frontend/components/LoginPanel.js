import Link from "next/link"
import {Search2} from "./Input"
import {NormalButton2} from "./Button"
import {Switch} from "./Switch";

export function LoginPanel() {
    return (
        <div className="container">
            <div className="loginPanel">
                <div className="companyLogo">
                    <img src="MS_logo-square.svg" alt=""/>
                </div>
                <div className="mainContent">
                    <div className="welcomeMessage">
                        <h1>Log In</h1>
                        <p>Welcome back! Login with your data that <br/> you entered during registration</p>
                    </div>
                    <div className="inputSet">
                        <Search2 type="text">Email or Username</Search2>
                        <Search2 type="password">Password</Search2>
                    </div>
                    <div className="passwordSupportSet">
                        <span className="rememberMe"><Switch/>Remember me</span>
                        <Link href="/"><a><span className="forgotPassword">Forgot Password?</span></a></Link>
                    </div>
                    <NormalButton2>
                        Login
                    </NormalButton2>
                </div>
            </div>
        </div>
    )
}