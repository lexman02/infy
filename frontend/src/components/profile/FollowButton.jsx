import React from "react";
import { useState } from "react";
import axios from "axios";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function FollowButton({ isFollowing = false, userID }) {
    const [following, setFollowing] = useState(isFollowing);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Function to add a follower to the user
    function followUser() {
        axios.post(`http://localhost:8000/profile/follow/${userID}`, {}, { withCredentials: true })
            .then(() => {
                setFollowing(true)
            }).catch(() => {
                setErrorMessage('An error occurred while following the user.');
            });
    }

    // Function to remove a follower from the user
    function unfollowUser() {
        axios.delete(`http://localhost:8000/profile/unfollow/${userID}`, { withCredentials: true })
            .then(() => {
                setFollowing(false)
            }).catch(() => {
                setErrorMessage('An error occurred while unfollowing the user.');
            });
    }

    return (
        <div>
            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

            {following ? (
                <button
                    className="flex justify-center items-center bg-violet-900 border-violet-950 border-2 text-neutral-50 rounded-2xl px-4 py-2 hover:bg-violet-950"
                    onClick={unfollowUser}>
                    Unfollow
                </button>
            ) : (
                <button
                    className="flex justify-center items-center bg-violet-900 border-violet-950 border-2 text-neutral-50 rounded-2xl px-4 py-2 hover:bg-violet-950"
                    onClick={followUser}>
                    Follow
                </button>
            )}
        </div>
    )
}