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
                if (response.data > 0) {
                    setPosts(response.data);
                }
                setPosts(0);
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
                {posts.length > 0 ? (
                    posts.map(post => <Post key={post.post.id} post={post.post} />)
                ) : (
                    <div className="p-4 text-neutral-300 font-medium text-lg bg-black/40 rounded-b-lg">
                        No posts yet... Be the first to post by selecting a movie above!
                    </div>
                )}
            </div>
        </div>
    );
}