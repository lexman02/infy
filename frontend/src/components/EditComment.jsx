import React from 'react';
import { PaperAirplaneIcon } from '@heroicons/react/20/solid';
import axios from 'axios';
import { useState } from "react";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function EditPost({ onEdit, comment }) {
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Handles the editted comment
    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        await axios.put(`http://localhost:8000/comments/${comment.id}`, data, { withCredentials: true })
            .then(() => {
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

        onEdit;
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <div className="mt-6 relative">
                    <label htmlFor="content" className="block mb-2 text-md font-medium"></label>
                    {/* disable auto complete and grammarly */}
                    <textarea id="content" name="content" defaultValue={comment.content} className="border rounded-t-lg border-indigo-900 bg-indigo-950/70 text-neutral-200 shadow-inner block w-full p-2.5" data-gramm="false" autoComplete="off" />
                    <div className="border border-t-0 border-indigo-900 rounded-b-lg flex items-center justify-between p-2">
                        <button className="bg-indigo-600 hover:bg-indigo-800 text-neutral-50 rounded-full p-2">
                            <PaperAirplaneIcon className="h-4 w-4" />
                        </button>
                    </div>
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