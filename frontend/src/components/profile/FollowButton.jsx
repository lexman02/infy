import React from "react";
import { useState } from "react";
import axios from "axios";


export default function FollowButton({ followList, isFollowing = false, userData }){
    
    // followList.includes(userData.user.id) ? setFollowing(true) : setFollowing(false);

    const [following, setFollowing] = useState(isFollowing);

    
    // Function to add a follower to the user
    function putFollow(){
        axios.post(`http://localhost:8000/follow/${userData.user.id}`, {}, {withCredentials: true})
        .then(response => { 
            setFollowing(true)
        }).catch(error => {
            console.error(error)
        });
    }

    // Function to remove a follower from the user
    function removeFollow(){
        axios.delete(`http://localhost:8000/follow/${userData.user.id}`, {withCredentials: true})
        .then(response => { 
            setFollowing(false)
        }).catch(error => {
            console.error(error)
        });
    }

    // Function to modify the user's follow properties upon clicking the follow button
    const followToggle = () => {
        if (!following){
            {/*putFollow()*/}
            setFollowing(true)
        }
        else{
            {/*removeFollow()*/}
            setFollowing(false)
        }
    }

    return (
        <button 
            className="flex justify-center items-center bg-violet-900 border-violet-950 border-2 text-neutral-50 rounded-2xl px-4 py-2 hover:bg-violet-950" 
            onClick={followToggle}>{following ? "Following": "Follow"}
        </button>
    )
    
}