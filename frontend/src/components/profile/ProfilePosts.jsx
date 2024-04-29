import React from "react";
import Post from "../Post";
import axios from "axios";
import { useState, useEffect } from "react";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';


export default function ProfilePosts(userID) {
    const [posts, setPosts] = useState([]);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Fetch posts
    useEffect(() => {
        axios.get(`http://localhost:8000/posts/user/${userID.userID}`, { withCredentials: true })
            .then(response => {
                if (response.data.length > 0) {
                    setPosts(response.data);
                } else {
                    setPosts(0);
                }
            })
            .catch(error => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occurred while fetching posts, or there are no posts.');
                }
            });
    }, [userID]);

    return (
        <div>
            <div className="flex flex-col justify-center ">
                <div className="divide-y divide-neutral-500">
                    {posts.length > 0 ? (
                        posts.map(post => <Post key={post.post.id} post={post} />)
                    ) : (
                        <div className="p-4 text-neutral-300 font-medium text-lg bg-black/40 rounded-b-lg text-center">
                            <h1>No posts yet...</h1>
                        </div>
                    )}
                </div>
            </div>

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

        </div>
    )
}