import React from "react";
import { useState } from "react";
import ProfilePosts from "./ProfilePosts";
import WatchedPosts from "./WatchedPosts";
import WatchlistPosts from "./WatchlistPosts";

export default function ProfileNav() {

    const [activeTab, setActiveTab] = useState('ProfilePosts')

    const renderTab = (component) => {
        switch (component) {
            case 'ProfilePosts':
                return <ProfilePosts />;
            case 'WatchedPosts':
                return <WatchedPosts />;
            case 'WatchlistPosts':
                return <WatchlistPosts />;
        }
    }

    return (
        <div>
            <div className="flex bg-black/40 border-y-2 border-neutral-700">
                <button className="mr-auto text-xl hover:bg-violet-950 p-3 w-1/3 " onClick={() => setActiveTab('ProfilePosts')}>Posts</button>
                <button className="mr-auto text-xl hover:bg-violet-950 p-3 w-1/3 " onClick={() => setActiveTab('WatchedPosts')}>Watched</button>
                <button className="mr-auto text-xl hover:bg-violet-950 p-3 w-1/3 " onClick={() => setActiveTab('WatchlistPosts')}>Watchlist</button>
            </div>
            {activeTab && renderTab(activeTab)}
        </div>
    );
}