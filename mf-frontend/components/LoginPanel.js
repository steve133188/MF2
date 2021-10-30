import Link from "next/link"
import {Search2} from "./Input"
import {NormalButton2} from "./Button"
import {Switch} from "./Switch";
import {AuthContext} from "../context/authContext";
import {useContext , useState} from "react";

export function LoginPanel() {
    const { loginUser,user ,customerFetching } = useContext(AuthContext);
    const [email, setEmail] = useState("");
    const [pwd, setPwd] = useState("");
    const [isInvalid, setInvalid] = useState("");

    function validateForm() {
        return email.length > 0 && pwd.length > 0;
    }

    function handleSubmit(event) {
        event.preventDefault()
        loginUser(email,pwd)
    }
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
                        <Search2 type="text" value={email} svg={"emailSVG"} invalid={isInvalid} handleChange={(e)=> {
                            setEmail(e.target.value)
                        }}>Email or Username</Search2>
                        <Search2 type="password"  value={pwd} svg={"passwordSVG"} invalid={isInvalid} handleChange={(e)=> {
                            setPwd(e.target.value);
                        }}>Password</Search2>
                    </div>
                    <div className="passwordSupportSet">
                        <span className="rememberMe"><Switch/>Remember me</span>
                        <Link href="/recovery"><a><span className="forgotPassword">Forgot Password?</span></a></Link>
                    </div>
                    <NormalButton2 disabled={!validateForm()} onClick={handleSubmit}>
                        Login
                    </NormalButton2>
                </div>
            </div>
        </div>
    )
}