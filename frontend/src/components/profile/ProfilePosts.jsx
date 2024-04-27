import React from "react";
import Post from "../Post";
import axios from "axios";
import { useState, useEffect } from "react";


export default function ProfilePosts(userID) {
    const [posts, setPosts] = useState([]);

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
            });
    }, []);

    return (
        <div>
            <div className="flex flex-col justify-center ">
                <div className="divide-y divide-neutral-500">
                    {posts.length > 0 ? (
                        posts.map(post => <Post key={post.post.id} post={post} />)
                    ) : (
                        <div className="p-4 text-neutral-300 font-medium text-lg bg-black/40 rounded-b-lg text-center">
                            <h1 className="">No posts yet...</h1>
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}