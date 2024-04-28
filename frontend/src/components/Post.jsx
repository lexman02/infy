import React, { useState, useContext } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";
import { HandThumbUpIcon, HandThumbDownIcon, ChatBubbleOvalLeftIcon, EllipsisHorizontalIcon, FlagIcon, PencilIcon, TrashIcon } from "@heroicons/react/20/solid";
import defaultAvatar from "../img/default-avatar.png";
import Popup from "reactjs-popup"
import { UserContext } from "../contexts/UserProvider";
import EditPost from "./EditPost";
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Post({ post, detailed = false }) {
    const [liked, setLiked] = useState(post.liked);
    const [disliked, setDisliked] = useState(post.disliked);
    const [likes, setLikes] = useState(post.likes);
    const [dislikes, setDislikes] = useState(post.dislikes);
    const [edit, setEdit] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');
    const { userData } = useContext(UserContext);

    const navigate = useNavigate();
    const navigateToPost = () => {
        if (!detailed) {
            navigate(`/post/${post.post.id}`);
        }
    }

    const handleCloseError = () => {
        setErrorMessage('');
    }

    const isAdmin = userData && (userData.user ? userData.user.isAdmin : false);
    const author = post.post.user.username;
    const avatar = post.post && post.post.user.profile.avatar ? `http://localhost:8000/avatars/${post.post.user.profile.avatar}` : defaultAvatar;
    const fullName = `${post.post.user.profile.first_name} ${post.post.user.profile.last_name}`;

    function handleLike() {
        axios.post(`http://localhost:8000/posts/${post.post.id}/like`, { is_liked: liked }, { withCredentials: true })
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
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occured while liking.');
                }
            });
    }

    function handleDislike() {
        axios.post(`http://localhost:8000/posts/${post.post.id}/dislike`, { is_disliked: disliked }, { withCredentials: true })
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
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occured while disliking.');
                }
            });
    }

    function handleEdit() {
        setEdit(true);
    }

    function handleDelete() {
        axios.delete(`http://localhost:8000/posts/${post.post.id}`, { withCredentials: true })
            .then(() => {
                console.log("Post deleted");
                window.location.reload();
            })
            .catch((error) => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occured during deletion.');
                }
            });
    }

    function handleReport() {
        axios.get(`http://localhost:8000/posts/${post.post.id}/report`, { withCredentials: true })
            .then(() => {
                console.log("Post reported");
                alert("Post has been reported.");
                window.location.reload();
            })
            .catch((error) => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.message) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.message);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occured during reporting.');
                }
            });
    }

    return (
        <div className="flex justify-between space-x-8 bg-black/40 p-4 text-neutral-100 last:rounded-b-lg">
            <div className="flex flex-col justify-around w-full">
                {/* Post author details */}
                <div className="flex space-x-2 items-center">
                    <Link to={"/profile/" + post.post.user.username}>
                        <img src={avatar} alt={post.post.user.username} className="w-10 h-10 rounded-full" />
                    </Link>
                    <div>
                        <Link to={"/profile/" + post.post.user.username} className="flex items-end space-x-1">
                            <h2 className="font-bold">
                                {fullName}
                            </h2>
                            <span className="text-neutral-500 text-sm font-light">@{post.post.user.username}</span>
                        </Link>
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

                {edit ?
                    <EditPost post={post.post} />
                    :
                    <p className="font-medium text-lg">
                        {post.post.content}
                    </p>
                }
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
                    <div className="flex space-x-3 items-center">
                        <Popup trigger={<EllipsisHorizontalIcon className="h-6 w-6 text-neutral-200 hover:cursor-pointer" />} position="right center">
                            <div className="flex flex-col space-y-2 px-5 py-1 bg-neutral-950 rounded-lg">
                                {userData && author === userData.user.username ? <button className="text-neutral-200 hover:bg-neutral-800 flex p-2 rounded-lg" onClick={handleEdit}><PencilIcon className="h-5 w-5" /><h1 className="pl-2">Edit</h1> </button> : null}
                                {userData && (isAdmin || author === userData.user.username ? <button className="text-neutral-200 hover:bg-neutral-800 flex p-2 rounded-lg" onClick={handleDelete}><TrashIcon className="h-5 w-5" /><h1 className="pl-2">Delete</h1></button> : null)}
                                <button className="text-neutral-200 hover:bg-neutral-800 flex p-2 rounded-lg" onClick={handleReport}><FlagIcon className="h-5 w-5" /><h1 className="pl-2">Report</h1></button>
                            </div>
                        </Popup>
                    </div>
                </div>
            </div>
            <img src={`https://image.tmdb.org/t/p/original/${post.post.movie.poster_path}`} alt={post.post.movie.title} className="w-20 h-32 object-cover rounded-lg" />

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>
        </div>
    )
}