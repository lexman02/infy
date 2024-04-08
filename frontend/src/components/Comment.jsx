import React from "react";
import {HandThumbUpIcon, HandThumbDownIcon} from "@heroicons/react/20/solid";
import defaultAvatar from "../img/default-avatar.png";

export default function Comment({ comment }) {
    const liked = 0;
    const disliked = 0;
    const fullName = `${comment.user.profile.first_name} ${comment.user.profile.last_name}`;

    function handleLike() {
        console.log('Like');
    }

    function handleDislike() {
        console.log('Dislike');
    }

    return (
        <div className="bg-black/40 text-neutral-100 p-4 last:rounded-b-lg border-t border-neutral-500">
            <div className="flex space-x-2 items-center">
                <img src={`${comment.user.profile.avatar || defaultAvatar}`} alt={fullName} className="w-11 h-11 rounded-full" />
                <div className="flex flex-col">
                    <h2 className="font-bold">
                        {fullName}
                    </h2>
                    <span className="text-neutral-500 text-sm font-light">@{comment.user.username}</span>
                </div>
            </div>
            <div className="flex justify-between space-y-1 py-2">
                <p className="text-md">
                    {comment.content}
                </p>
                <div className="flex space-x-3">
                    <div className="flex space-x-1 items-center">
                        <HandThumbUpIcon className={`h-6 w-6 cursor-pointer ${liked ? 'text-blue-500 hover:text-blue-500/80' : 'text-neutral-400 hover:text-neutral-400/80'}`} onClick={() => handleLike()} />
                        <span className="text-neutral-500 font-light">{comment.likes}</span>
                    </div>
                    <div className="flex space-x-1 items-center">
                        <HandThumbDownIcon className={`h-6 w-6 cursor-pointer ${disliked ? 'text-red-500 hover:text-red-500/80' : 'text-neutral-400 hover:text-neutral-400/80'}`} onClick={() => handleDislike()} />
                        <span className="text-neutral-500 font-light">{comment.dislikes}</span>
                    </div>
                </div>
            </div>
        </div>
    );
}
