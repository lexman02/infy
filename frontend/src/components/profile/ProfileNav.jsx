import React from "react";
import { useState } from "react";
import ProfilePosts from "./ProfilePosts";
import WatchedPosts from "./WatchedPosts";
import WatchlistPosts from "./WatchlistPosts";

export default function ProfileNav(user) {
    const [activeTab, setActiveTab] = useState('ProfilePosts')
    const watched = user.user.profile.preferences.watched;
    const watchlist = user.user.profile.preferences.watchlist

    // Tabs on the profile page
    const renderTab = (component) => {
        switch (component) {
            case 'ProfilePosts':
                return <ProfilePosts userID={user.user.id} />;
            case 'WatchedPosts':
                return <WatchedPosts watched={watched} />;
            case 'WatchlistPosts':
                return <WatchlistPosts watchlist={watchlist} />;
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