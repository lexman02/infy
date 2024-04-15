import React from "react";
import ProfileNav from "./ProfileNav";
import { useEffect, useState } from "react";
import axios from "axios";
import defaultAvatar from "../../img/default-avatar.png"
import FollowButton from "./FollowButton";
import { WrenchIcon } from "@heroicons/react/20/solid";

export default function UserProfile({ userData }) {
    const [profileData, setProfileData] = useState(null);

    async function getProfileData() {
        await axios.get('http://localhost:8000/profile/user/', { withCredentials: true })
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
        <div className="md:my-6 md:mx-60 flex-grow rounded-lg">
            {profileData && (
                <div>
                    <div className="bg-black/40 p-5 rounded-lg">
                        <div className="flex flex-col space-y-4 md:flex-row md:justify-between md:items-center">
                            <div className="flex space-x-4 items-center">
                                <img src={`${defaultAvatar}`} className=" w-28 h-28 rounded-full" />
                                <div>
                                    <h1 className="text-xl">
                                        {profileData.profile.first_name} {profileData.profile.last_name}
                                        {userData.user.isAdmin ? <WrenchIcon className="w-6 h-6 inline-block ml-2" /> : null}
                                    </h1>
                                    <h1 className="text-lg">@{userData.user.username}</h1>
                                </div>
                            </div>
                            <FollowButton isFollowing={false} userData={userData} />
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