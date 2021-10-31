import Link from "next/link"
import {Search2} from "./Input"
import {NormalButton2} from "./Button"
import {Switch} from "./Switch";
import {AuthContext} from "../context/authContext";
import {useContext , useState} from "react";
import {Alert} from "./Alert";

export function LoginPanel() {
    const { login,user } = useContext(AuthContext);
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [isInvalid, setInvalid] = useState("");

    function validateForm() {
        return email.length > 0 && password.length > 0;
    }

    function handleSubmit(event) {
        event.preventDefault()
        console.log('clicked log in')
        login(email,password)
        validatePassword()
    }
    function validatePassword() {
        if (login(email,password)=="Something went Wrong") {
            console.log("Email or password invalid")
            setInvalid("wrongPwd")
        } else {
            console.log("Email and password are valid")
            return ""
        }
    }
    return (
        <div className="container">
            <Alert />
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
                            console.log(email)
                        }}>Email or Username</Search2>
                        <Search2 type="password"  value={password} svg={"passwordSVG"} invalid={isInvalid} handleChange={(e)=> {
                            setPassword(e.target.value);
                            console.log(password);
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