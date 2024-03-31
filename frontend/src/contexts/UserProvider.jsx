import React, { useState } from "react";

export const UserContext = React.createContext();

export default function UserProvider({ children }) {
    const [userData, setUserData] = useState(null);

    return (
        <UserContext.Provider value={{userData, setUserData}}>
            {children}
        </UserContext.Provider>
    );
}