import React, { useState } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import {HandThumbUpIcon, HandThumbDownIcon, ChatBubbleOvalLeftIcon} from "@heroicons/react/20/solid";
import defaultAvatar from "../img/default-avatar.png";

export default function Post({ post, detailed = false }) {
    const [liked, setLiked] = useState(post.liked);
    const [disliked, setDisliked] = useState(post.disliked);
    const [likes, setLikes] = useState(post.likes);
    const [dislikes, setDislikes] = useState(post.dislikes);

    const navigate = useNavigate();
    const navigateToPost = () => {
        if (!detailed) {
            navigate(`/post/${post.post.id}`);
        }
    }

    const fullName = `${post.post.user.profile.first_name} ${post.post.user.profile.last_name}`;
    
    function handleLike() {
        axios.post(`http://localhost:8000/posts/${post.post.id}/like`, {is_liked: liked}, {withCredentials: true})
            .then(() => {
                if (liked) {
                    setLikes(likes - 1);
                    if (disliked) {
                        setDislikes(dislikes + 1);
                    }
                    setLiked(false);
                } else {
                    setLikes(likes + 1);
                    if (disliked) {
                        setDislikes(dislikes - 1);
                    }
                    setLiked(true);
                }
                setDisliked(false);
            })
            .catch((error) => {
                console.error(error);
            });
    }
    
    function handleDislike() {
        axios.post(`http://localhost:8000/posts/${post.post.id}/dislike`, {is_disliked: disliked}, {withCredentials: true})
            .then(() => {
                if (disliked) {
                    setDislikes(dislikes - 1);
                    if (liked) {
                        setLikes(likes + 1);
                    }
                    setDisliked(false);
                } else {
                    setDislikes(dislikes + 1);
                    if (liked) {
                        setLikes(likes - 1);
                    }
                    setDisliked(true);
                }
                setLiked(false);
            })
            .catch((error) => {
                console.error(error);
            });
    }

    return (
        <div className="flex justify-between bg-black/40 p-4 text-neutral-100 last:rounded-b-lg">
            <div className="flex flex-col justify-around">
                {/* Post author details */}
                <div className="flex space-x-2 items-center">
                    <img src={`${post.post.user.profile.avatar || defaultAvatar}`} alt={fullName} className="w-11 h-11 rounded-full" />
                    <div>
                        <div className="flex items-end space-x-1">
                            <h2 className="font-bold">
                                {fullName}
                            </h2>
                            <span className="text-neutral-500 text-sm font-light">@{post.post.user.username}</span>
                        </div>
                        <div className="flex items-end space-x-1 text-sm text-neutral-400">
                            <p className="font-light">
                                watched
                            </p>
                            <Link to={"/movie/" + post.post.movie.id}>
                                <span className="font-medium">{post.post.movie.title}</span> 
                            </Link>
                        </div>
                    </div>
                </div>
                {/* Post content */}
                <p className="font-medium text-lg">
                    {post.post.content}
                </p>
                {/* Post interaction buttons */}
                <div className="flex space-x-3">
                    <div className="flex space-x-2 items-center">
                        <HandThumbUpIcon className={`h-6 w-6 cursor-pointer ${liked ? 'text-blue-500 hover:text-blue-500/80' : 'text-neutral-400 hover:text-neutral-400/80'}`} onClick={() => handleLike()} />
                        <span className="text-neutral-500 font-light">{likes}</span>
                    </div>
                    <div className="flex space-x-2 items-center">
                        <HandThumbDownIcon className={`h-6 w-6 cursor-pointer ${disliked ? 'text-red-500 hover:text-red-500/80' : 'text-neutral-400 hover:text-neutral-400/80'}`} onClick={() => handleDislike()} />
                        <span className="text-neutral-500 font-light">{dislikes}</span>
                    </div>
                    {!detailed && <ChatBubbleOvalLeftIcon className="h-6 w-6 text-neutral-200 hover:cursor-pointer" onClick={navigateToPost} />}
                </div>
            </div>
            <img src={`https://image.tmdb.org/t/p/original/${post.post.movie.poster_path}`} alt={post.post.movie.title} className="w-20 h-32 object-cover rounded-lg" />
        </div>
    )
}