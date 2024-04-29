import React, { useEffect, useState } from "react";
import Post from "../components/Post";
import NewPost from "../components/NewPost";
import axios from "axios";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Home() {
    const [posts, setPosts] = useState([]);
    const [newPost, setNewPost] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Load Posts
    useEffect(() => {
        axios.get('http://localhost:8000/posts/', { withCredentials: true })
            .then(response => {
                if (response.data.length > 0) {
                    setPosts(response.data);
                } else {
                    setPosts(0);
                }
            })
            .catch(error => {
                if (error.response && error.response.data && error.response.data.error) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.error);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occurred while getting posts.');
                }
            });
    }, [newPost]);

    // Set New Posts
    const handleNewPost = () => {
        setNewPost(!newPost);
    }

    return (
        <div className="md:my-6 md:mx-60 flex-grow">
            {/* button to add a new post */}
            <div className="border-b rounded-t-lg border-neutral-500 bg-black/40 p-4">
                <NewPost onNewPost={handleNewPost} />
            </div>
            {/* display the posts */}
            <div className="divide-y divide-neutral-500">
                {posts.length > 0 ? (
                    posts.map(post => <Post key={post.post.id} post={post} />)
                ) : (
                    <div className="p-4 text-neutral-300 font-medium text-lg bg-black/40 rounded-b-lg">
                        No posts yet... Be the first to post by selecting a movie above!
                    </div>
                )}
            </div>

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>
        </div>
    );
}