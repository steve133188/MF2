import axios from "axios";

const API_URL = "http://localhost:8080/api/auth/";

const login = (email, password) => {
    return axios
        .post(API_URL + "signin", {
            email,
            password,
        })
        .then((response) => {
            if (response.body.token) {
                localStorage.setItem("user", JSON.stringify(response.body));
            }

            return response.body;
        });
};

const logout = () => {
    localStorage.removeItem("user");
};

const getCurrentUser = () => {
    return JSON.parse(localStorage.getItem("user"));
};

export default {
    login,
    logout,
    getCurrentUser,
};

// res:{
//     header,
//         status,
//         body,
//         data{
//         ...data,
//             token
//     },
//
// }