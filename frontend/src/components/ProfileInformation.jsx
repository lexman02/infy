import ProfileNav from "./ProfileNav";
import {useEffect, useState} from "react";
import axios from "axios";
import defaultAvatar from "../img/default-avatar.png";

export default function ProfileInformation({userData}){
    const [profileData, setProfileData] = useState(null);

    async function getProfileData() {
        await axios.get('http://localhost:8000/profile/user/', {withCredentials: true})
        .then(response => {
            if (response.data) {
                setProfileData(response.data);
            }
        })
        .catch(error => {
            console.error(error);
            setProfileData(null);
        });
    }
    
    useEffect(() => {
        getProfileData();
    }, []);

    return (
        <div className="flex justify-center">
            {profileData && (
            <div className=" bg-indigo-900 w-1/2 rounded-lg">
                <div>
                    <div className="flex items-start bg-indigo-900 p-5 rounded-lg">
                        <div>
                            <img src={`${defaultAvatar}`} className=" w-28 h-28 rounded-full" />
                            <br/>
                            <h1 className="text-xl">{profileData.profile.first_name} {profileData.profile.last_name}</h1>
                            <h1 className="text-lg">@{userData.user.username}</h1>
                            <h1 className="text-base">{profileData.profile.rank}</h1>
                        </div>
                    </div>
                </div>
                <div>
                    <ProfileNav />
                </div>
            </div>
            )}
        </div>
    );
}