import React, { useState } from 'react';
import { PaperAirplaneIcon } from '@heroicons/react/20/solid';
import axios from 'axios';

export default function NewComment({ onNewComment, postID }) {
    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        data.post_id = postID;
        await axios.post('http://localhost:8000/comments/', data, {withCredentials: true})
            .then(() => {
                onNewComment();
            })
            .catch(error => {
                console.error(error);
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
        </div>
    );
}