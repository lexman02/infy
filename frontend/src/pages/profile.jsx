import {useContext} from "react";
import axios from "axios";
import { UserContext } from "../contexts/UserProvider";
import ProfileInformation from "../components/ProfileInformation";

export default function Profile(){
    const {userData, setUserData} = useContext(UserContext);

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