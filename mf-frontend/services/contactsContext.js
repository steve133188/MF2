// import {useEffect, createContext, Context, useState} from "react";
// import {useRouter} from "next/router";
// import {redirect} from "next/dist/server/api-utils";
// import AuthService from "../services/auth";
// import axios from "axios";
//
// export const ContactsContext = createContext({
//     userInfo: null,
//     authReady: false,
// })
//
//
// export const AuthContextProvider = ({children}) => {
//
//     const [isInvalid, setInvalid] = useState("");
//     const [customer, setCustomer] = useState({
//         userInfo: {
//             username: null,
//             email: null,
//             role: null,
//             organization: {},
//             authority: [],
//             phone: null,
//             status: "online",
//             notification: 0,
//         },
//         authReady: false,
//     })
//
//     const router = useRouter()
//
//     const customerFetching = async (email, pwd) => {
//         const body = await axios.get('https://mf-api-customer-nccrp.ondigitalocean.app/api/customers', { email, password:pwd }).then(
//             (res) => {
//                 if(res.status==200) {
//                     const {token} = res.data;
//                     console.log(token)
//                     localStorage.setItem('token', token);
//                     const testingJWT = parseJwt(token)
//                     console.log(testingJWT)
//                     setUser({userInfo:{username: testingJWT.username, email: testingJWT.email, role: testingJWT.role}, authReady: true})
//                     console.log(user)
//                     router.push("/testing")
//                 }
//             }
//         ).catch(err=> {
//             if(err) {
//                 console.log(err)
//                 return "Something went Wrong"
//             }
//         })
//     }
//
//     function parseJwt(token) {
//         if (!token) { return; }
//         const base64Url = token.split('.')[1];
//         const base64 = base64Url.replace('-', '+').replace('_', '/');
//         return JSON.parse(atob(base64));
//     }
//
//
//     return (
//         <ContactsContext.Provider value={{customer, customerFetching}}>{children}</ContactsContext.Provider>
//     )
// }