import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { ArrowLeftIcon } from '@heroicons/react/20/solid';
import Post from '../components/Post';
import Comment from '../components/Comment';
import NewComment from '../components/NewComment';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function DetailedPost() {
    const [post, setPost] = useState(null);
    const [newComment, setNewComment] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const { postID } = useParams();
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    const handleNewComment = () => {
        setNewComment(!newComment);
    }

    useEffect(() => {
        axios.get(`http://localhost:8000/posts/${postID}`, {withCredentials: true})
            .then((response) => {
                setPost(response.data);
                setIsLoading(false);
            })
            .catch((error) => {
                console.error(error);
                setIsLoading(false);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occurred while getting posts. Wait a bit then try again.');
                }
            });
    }, [newComment, postID]);

    if (isLoading) {
        return <div>Loading...</div>
    }

    return (
        <div className="md:my-6 md:mx-60 flex-grow">
            <div className="bg-black/40 px-4 py-2 text-neutral-300 border-b border-neutral-500 rounded-t-lg" onClick={() => window.history.back()}>
                <div className="inline-flex align-middle space-x-1 hover:cursor-pointer">
                    <ArrowLeftIcon className="h-6 w-6"/>
                    <span className="font-medium">
                        Back
                    </span>
                </div>
            </div>

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

            <Post post={post} detailed={true} />
            <div className=" border-t border-neutral-500 bg-black/40 px-4 pb-4 pt-2">
                <NewComment onNewComment={handleNewComment} postID={post.post.id} />
            </div>
            <div className="divide-y divide-neutral-500">
                {post.comments !== null ? (
                    post.comments.map(comment => <Comment key={comment.id} comment={comment} />)
                    // post.comments.map(comment => console.log(comment))
                ) : (
                    <span>No comments yet...</span>
                )}
            </div>
        </div>
    )
}