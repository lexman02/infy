import React, { useState, useContext, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";
import { WrenchIcon } from "@heroicons/react/20/solid";
import { UserContext } from "../contexts/UserProvider";
import defaultAvatar from "../img/default-avatar.png"
import FollowButton from "../components/profile/FollowButton";
import ProfileNav from "../components/profile/ProfileNav";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Profile() {
    const [profileData, setProfileData] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const { userData } = useContext(UserContext);
    const { username } = useParams();
    const navigate = useNavigate();

    function getProfileData() {
        axios.get(`http://localhost:8000/profile/${username}`, { withCredentials: true })
            .then(response => {
                if (response.data) {
                    setProfileData(response.data.user);
                }
            })
            .catch(error => {
                console.error(error);
                setProfileData(null);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('The user profile could not be loaded at this time. Please try again later.');
                }
            });
    }

    if (!profileData) {
        navigate('/404');
    }

    function adminAccess() {
        if (userData.user.isAdmin) {
            window.location.href = "/admin";
        }
    }

    useEffect(() => {
        getProfileData();
    }, []);

    const handleCloseError = () => {
        setErrorMessage('');
    };

    const avatar = profileData && (profileData.profile.avatar ? `http://localhost:8000/avatars/${profileData.profile.avatar}` : defaultAvatar);

    return (
        <div className="md:my-6 md:mx-60 flex-grow rounded-lg">
            {profileData && (
                <div>
                    <div className="bg-black/40 p-5 rounded-t-lg">
                        <div className="flex flex-col space-y-4 md:flex-row md:justify-between md:items-center">
                            <div className="flex space-x-4 items-center">
                                <img src={avatar} className=" w-28 h-28 rounded-full" />
                                <div className="flex items-center space-x-4">
                                    <div>
                                        <h1 className="text-xl">
                                            {profileData.profile.first_name} {profileData.profile.last_name}
                                        </h1>
                                        <h1 className="text-lg">@{profileData.username}</h1>
                                    </div>
                                    {profileData.is_admin ? <WrenchIcon className="w-6 h-6 inline-block ml-2 hover:cursor-pointer" onClick={adminAccess} /> : null}
                                </div>
                            </div>
                            {userData && (
                                <FollowButton isFollowing={false} userData={userData} />
                            )}
                        </div>
                    </div>
                    <div>
                        <ProfileNav user={profileData} />
                    </div>
                </div>
            )}
            <Snackbar Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError} >
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>
        </div>
    );
}