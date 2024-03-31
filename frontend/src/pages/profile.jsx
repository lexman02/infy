import {useContext, useEffect} from "react";
import axios from "axios";
import { UserContext } from "../contexts/UserProvider";

export default function Profile(){
    const {userData, setUserData} = useContext(UserContext);

    function getUserData() {
        if (userData === null) {
            axios.get('http://localhost:8000/auth/user', {withCredentials: true})
            .then(response => {
                if (response.data) {
                    setUserData(response.data);
                }
            })
            .catch(error => {
                console.error(error);
                window.location.href = '/login';
            });
        }
    }

    function logout() {
        axios.post('http://localhost:8000/auth/logout', {}, {withCredentials: true})
        .then(() => {
            setUserData(null);
            window.location.href = '/';
        })
        .catch(error => {
            console.error(error);
        });
    }

    useEffect(() => {
        getUserData();
    });

    return (
        <div>
            {userData && (
                <div>
                    <h1 className="text-neutral-50">Welcome back, {userData.user.username}!</h1>
                    <h2 className="text-neutral-50">Email: {userData.user.email}</h2>
                    <button onClick={logout}>Logout</button>
                </div>
            )}
        </div>
    );
} 