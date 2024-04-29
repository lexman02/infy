import React, { useEffect, useState } from "react"
import axios from "axios"
import Post from "../Post";

export default function ReportedPosts() {
    const [posts, setPosts] = useState([]);

    useEffect(() => {
        axios.get(`http://localhost:8000/admin/reports/posts`, { withCredentials: true })
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
        <div className="divide-y divide-neutral-500">
            {posts.length > 0 ? (
                posts.map(post => <Post key={post.post.id} post={post} />)
            ) : (
                <div className="p-4 text-neutral-300 font-medium text-lg bg-black/40 rounded-b-lg">
                    No reported posts! Good job!
                </div>
            )}
        </div>
    )
}