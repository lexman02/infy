import React from "react";
import { useState, useEffect } from "react";
import Post from "./Post";
import axios from "axios";
import { UserContext } from "../contexts/UserProvider";

export default function ProfileNav(){
    
    const [activeTab, setActiveTab] = useState('ProfilePosts')
    
    const renderTab = (component) =>{
        switch (component) {
            case 'ProfilePosts':
              return <ProfilePosts />;
            default:
              return null;
          }
    }
    
    return (
        <div>
            <div className="flex w-auto border-t-4 p-4 border-indigo-700">
                <button className="mr-auto text-xl " onClick={() => setActiveTab('ProfilePosts')}>Posts</button>
                <button className="mr-auto text-xl ">Likes</button>
                <button className="mr-auto text-xl ">Watched</button>
                <button className="mr-auto text-xl ">Watchlist</button>
            </div>
            {activeTab && renderTab(activeTab)}
        </div>
    );
}

function ProfilePosts(){
    const [posts, setPosts] = useState([]);
    const {userData} = React.useContext(UserContext);

    useEffect(() => {
        axios.get(`http://localhost:8000/posts/user/${userData.user.id}`, {withCredentials: true})
            .then(response => {
                if (response.data.length > 0) {
                    setPosts(response.data);
                    console.log(response.data)
                } else {
                    setPosts(0);
                }
            })
            .catch(error => {
                console.error(error);
            });
    }, []);

    return(
        <div>
            <div className="flex flex-col justify-center w-auto">
                <div className="divide-y divide-neutral-500">
                    {posts.length > 0 ? (
                        posts.map(post => <Post key={post.post.id} post={post} />)
                    ) : (
                        <div className="p-4 text-neutral-300 bg-indigo-900 font-medium text-lg bg-black/40 rounded-b-lg text-center">
                            <h1 className="">No posts, yet...</h1>
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}