import React, { useEffect, useState } from "react";
import Post from "../components/Post";
import NewPost from "../components/NewPost";
import axios from "axios";

export default function Home(){
    const [posts, setPosts] = useState([]);
    const [newPost, setNewPost] = useState(false);

    useEffect(() => {
        const fetchPosts = axios.get('http://localhost:8000/posts/', {withCredentials: true})
            .then(response => {
                setPosts(response.data);
            })
            .catch(error => {
                console.error(error);
            });
        
        fetchPosts;
    }, [newPost]);

    const handleNewPost = () => {
        setNewPost(!newPost);
    }

    return(
        <div className="md:my-6 md:mx-60">
            {/* button to add a new post */}
            <div className="border-b rounded-t-lg border-neutral-500 bg-black/40 p-4">
                <NewPost onNewPost={handleNewPost} />
            </div>
            {/* display the posts */}
            <div className="divide-y divide-neutral-500">
                {posts.map(post => <Post key={post.post.id} post={post.post} />)}
            </div>
        </div>
    );
}