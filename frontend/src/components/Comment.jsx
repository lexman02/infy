import React from "react";
import {HandThumbUpIcon, HandThumbDownIcon} from "@heroicons/react/20/solid";
import defaultAvatar from "../img/default-avatar.png";
import {EllipsisHorizontalIcon, FlagIcon, PencilIcon, TrashIcon} from "@heroicons/react/20/solid";
import Popup from "reactjs-popup";
import { UserContext } from "../contexts/UserProvider";
import { useContext, useState } from "react";
import axios from "axios";
import EditComment from "./EditComment";

export default function Comment({ comment }) {
    const liked = 0;
    const disliked = 0;
    const fullName = `${comment.user.profile.first_name} ${comment.user.profile.last_name}`;
    const author = comment.user.username;
    const { userData } = useContext(UserContext);
    const isAdmin = userData.user.isAdmin;
    const [edit, setEdit] = useState(false);

    function handleLike() {
        console.log('Like');
    }

    function handleDislike() {
        console.log('Dislike');
    }

    function handleEdit(){
        setEdit(true);
    }
    
    function handleDelete(){
        axios.delete(`http://localhost:8000/comments/${comment.id}`, {withCredentials: true})
            .then(() => {
                console.log("Comment deleted");
                window.location.reload();
            })
            .catch((error) => {
                console.error(error);
            });
    }

    function handleReport(){
        axios.post(`http://localhost:8000/comments/${comment.id}/report`, {}, {withCredentials: true})
            .then(() => {
                console.log("Comment reported");
                alert("Comment has been reported.");
                window.location.reload();
            })
            .catch((error) => {
                console.error(error);
            });
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
                {edit ? 
                    <EditComment comment={comment} /> 
                    : 
                    <p className="font-medium text-lg">
                        {comment.content}
                    </p>
                }
                <div className="flex space-x-3">
                    <div className="flex space-x-1 items-center">
                        <HandThumbUpIcon className={`h-6 w-6 cursor-pointer ${liked ? 'text-blue-500 hover:text-blue-500/80' : 'text-neutral-400 hover:text-neutral-400/80'}`} onClick={() => handleLike()} />
                        <span className="text-neutral-500 font-light">{comment.likes}</span>
                    </div>
                    <div className="flex space-x-1 items-center">
                        <HandThumbDownIcon className={`h-6 w-6 cursor-pointer ${disliked ? 'text-red-500 hover:text-red-500/80' : 'text-neutral-400 hover:text-neutral-400/80'}`} onClick={() => handleDislike()} />
                        <span className="text-neutral-500 font-light">{comment.dislikes}</span>
                    </div>
                    <div className="flex space-x-3 items-center">
                        <Popup trigger={<EllipsisHorizontalIcon className="h-6 w-6 text-neutral-200 hover:cursor-pointer" />} position="right center">
                            <div className="flex flex-col space-y-2 px-5 py-1 bg-black">
                                {author === userData.user.username ? <button className="text-neutral-200 hover:bg-neutral-800 flex p-2 rounded-lg" onClick={handleEdit}><PencilIcon className="h-5 w-5"/><h1 className="pl-2">Edit</h1> </button>: null}
                                {isAdmin || author == userData.user.username ? <button className="text-neutral-200 hover:bg-neutral-800 flex p-2 rounded-lg" onClick={handleDelete}><TrashIcon className="h-5 w-5"/><h1 className="pl-2">Delete</h1></button> : null}
                                <button className="text-neutral-200 hover:bg-neutral-800 flex p-2 rounded-lg" onClick={handleReport}><FlagIcon className="h-5 w-5"/><h1 className="pl-2">Report</h1></button>
                            </div>
                        </Popup>
                    </div>
                </div>
            </div>
        </div>
    );
}
