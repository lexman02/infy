import {useContext, useEffect} from "react";
import axios from "axios";
import { UserContext } from "../contexts/UserProvider";
import ProfileInformation from "../components/ProfileInformation";

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
                    <ProfileInformation userData={userData}/>
                    <button onClick={logout} className="flex justify-center items-center bg-violet-900 text-neutral-50 rounded-lg px-4 py-2 hover:bg-violet-950">Logout</button>
                </div>
            )}
        </div>
    );
}