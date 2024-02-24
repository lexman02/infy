import {useEffect, useState} from "react";
import axios from "axios";

export default function Profile(){
    const [userData, setUserData] = useState(null);

    async function getUserData() {
        // Check if user data is already in session storage
        let data = sessionStorage.getItem('userData');
        if (data) {
            // Parse the user data from JSON
            data = JSON.parse(data);
        } else {
            try {
                const response = await axios.get('http://localhost:8000/auth/user', {withCredentials: true});
                data = response.data;
                // Save the user data in session storage
                sessionStorage.setItem('userData', JSON.stringify(data));
            } catch (error) {
                console.error(error);
            }
        }

        setUserData(data);
    }

    useEffect(() => {
        getUserData().then(r => console.log(r));
    }, []);

    return (
        <div>
            {userData && (
                <div>
                    <h1 className="text-neutral-50">Welcome back, {userData.user.username}!</h1>
                    <h2 className="text-neutral-50">Email: {userData.user.email}</h2>
                </div>
            )}
        </div>
    );
} 