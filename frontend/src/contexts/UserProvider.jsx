import React, { useEffect, useState } from "react";
import axios from "axios";

export const UserContext = React.createContext();

export default function UserProvider({ children }) {
    const [userData, setUserData] = useState(null);


    function getUserData() {
        if (userData === null) {
            axios.get('http://localhost:8000/auth/user', {withCredentials: true})
            .then(response => {
                if (response.data) {
                    setUserData(response.data);
                }
            })
            .catch(error => {
                if (error.response.status === 401) {
                    setUserData(null);
                }
            });
        }
    }

    useEffect(() => {
        getUserData();
    });

    return (
        <UserContext.Provider value={{userData, setUserData}}>
            {children}
        </UserContext.Provider>
    );
}