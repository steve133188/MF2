import {useEffect, createContext, Context, useState, useReducer} from "react";
import {useRouter} from "next/router";
import {redirect} from "next/dist/server/api-utils";
import AuthService from "../services/auth";
import axios from "axios";

export const AuthContext = createContext()

const initialState = {
    auth: {},
    user: []
};

const ACTIONS = {
    AUTH: "AUTH",
    ADD_USERS: "ADD_USERS"
};

const reducer = (stateAuth, action) => {
    switch(action.type) {
        case ACTIONS.AUTH:
            return {
                ...stateAuth,
                auth: action.payload
            };
        case ACTIONS.ADD_USERS:
            return {
                ...stateAuth,
                user: action.payload
            };
        default:
            return stateAuth;
    };
};

export const AuthContextProvider = ({children}) => {
    const [stateAuth, dispatchAuth] = useReducer(reducer, initialState);

    const [isInvalid, setInvalid] = useState("");
    const [user, setUser] = useState({
        userInfo: {
            username: null,
            email: null,
            role: null,
            organization: {},
            authority: [],
            phone: null,
            status: "online",
            notification: 0,
        },
        authReady: false,
    })
    const [customer, setCustomer] = useState({

    })
    const router = useRouter()
    useEffect(() => {
        const firstLogin = localStorage.getItem("firstLogin");
        if (firstLogin === true) {
            async (token) => {
                const res = await fetch("http://localhost:3000/api/users/", {
                    method: "GET",
                    headers: {
                        "Authorization": token
                    }
                });
                await res.json().then(res => {
                    res.err ? localStorage.removeItem("firstLogin") : dispatchAuth({
                        type: ACTIONS.AUTH,
                        payload: {
                            token: res.token,
                            user: res.user
                        }
                    });
                });
            };
        }
    }, []);

    const loginUser = async (email, pwd) => {
        const body = await axios.post('https://mf-api-user-sj8ek.ondigitalocean.app/mf-2/api/users/login', { email, password:pwd }).then(
            (res) => {
                if(res.status==200) {
                    const {token} = res.data;
                    console.log("token: ",token)
                    localStorage.setItem('token', token);
                    const JWT = parseJwt(token)
                    console.log(JWT)
                    setUser({userInfo:{username: JWT.username, email: JWT.email, role: JWT.role}, authReady: true})
                    console.log("user: ",user)
                    // checkTokenExpirationMiddleware(user)
                    // console.log("checkTokenExpirationMiddleware: " + user)
                    router.push("/")
                }
            }
        ).catch(err=> {
            if(err) {
                console.log(err)
                return "Something went Wrong"
            }
        })
    }

    const customerFetching = async (email, pwd) => {
        axios.defaults.headers.common['Authorization'] = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN0ZXZlY2hha2N5QGdtYWlsLmNvbSIsImV4cCI6MTYzNjE0NzI2MCwicGFzc3dvcmQiOiIxMjM0NSIsInJvbGUiOiIiLCJ1c2VybmFtZSI6InN0ZXZlIn0.-6vWKoi2OtH4tu4ilGum0ZUTTQBfKk3fl78ItpGzhJw";
        const body = await axios.get('https://mf-api-customer-nccrp.ondigitalocean.app/api/customers').then(
            (res) => {
                if(res.status==200) {
                    const {token} = res.data;
                    console.log(token)
                    localStorage.setItem('token', token);
                    const testingJWT = parseJwt(token)
                    console.log(testingJWT)
                    setUser({userInfo:{username: testingJWT.username, email: testingJWT.email, role: testingJWT.role}, authReady: true})
                    console.log(user)
                    router.push("/testing")
                }
            }
        ).catch(err=> {
            if(err) {
                console.log(err)
                return "Something went Wrong"
            }
        })
    }

    function parseJwt(token) {
        if (!token) { return; }
        const base64Url = token.split('.')[1];
        const base64 = base64Url.replace('-', '+').replace('_', '/');
        return JSON.parse(atob(base64));
    }

        // function validatePassword() {
        // if (login(email, pwd) == "Something went Wrong") {
        //     console.log("Email or password invalid")
        //     setInvalid("wrongPwd")
        // } else {
        //     console.log("Email and password are valid")
        //     return ""
        // }
        // }

        // if(email == "wiva.wei@matrixsense.tech" && pwd == "1234"){
        //     console.log("login success")
        //     setUser({userInfo:{name:"Wiva " , email: email , role:"super admin" , organization: {"Matrixsense":"CEO"}}, authReady: true})
        //     router.push("/dashboard")
        // }else if (email == "steve.chak@matrixsense.tech" && pwd =="1234"){
        //     setUser({userInfo:{name:"Steve.Chak " }, authReady: true})
        //     router.push("/dashboard")
        // }else if(email =="ben.cheng@matrixsense.tech" && pwd == "1234"){
        //     setUser({userInfo:{name:"Ben.cheng " , email: email , role:"admin" , organization: {"Matrixsense":"Developer"}}, authReady: true})
        //     router.push("/dashboard")
        // }else if(email =="lewis.chan@matrixsense.tech"  && pwd == "1234"){
        //     setUser({userInfo:{name:"Lewis.chan " , email: email , role:"admin" , organization: {"Matrixsense":"Developer"}}, authReady: true})
        //     router.push("/dashboard")
        // }else{
        //     console.log("Something went Wrong")
        //     return "Something went Wrong"
        // }

    return (
        <AuthContext.Provider value={{ stateAuth, dispatchAuth}}>{children}</AuthContext.Provider>
    )
}