import React, { useState } from 'react';
import { PaperAirplaneIcon } from '@heroicons/react/20/solid';
import axios from 'axios';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function NewComment({ onNewComment, postID }) {
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        data.post_id = postID;
        await axios.post('http://localhost:8000/comments/', data, {withCredentials: true})
            .then(() => {
                onNewComment();
                window.location.reload(); // Refresh the page
            })
            .catch(error => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occured.');
                }
            });
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <div className="relative">
                    <label htmlFor="content" className="block mb-2 text-md font-medium">New Comment</label>
                    <textarea id="content" name="content" placeholder="Anything to say?" className="border rounded-lg border-indigo-500/30 bg-indigo-950/70 text-neutral-200 shadow-inner block w-full p-2.5" data-gramm="false" autoComplete="off"/>
                    {/* disable auto complete and grammarly */}
                    <button className="absolute bottom-2 right-2 bg-indigo-900 hover:bg-indigo-800 text-neutral-50 rounded-full p-2">
                        <PaperAirplaneIcon className="h-4 w-4" />
                    </button>
                </div>
            </form>

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

        </div>
    );
}