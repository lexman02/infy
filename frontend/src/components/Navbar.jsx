import React from "react";
import {Link} from "react-router-dom";

export default function Navbar(){
    return (
        <div className="text-neutral-50 fixed bottom-0 left-0 right-0">
            <nav className="bg-violet-900 flex flex-row margin justify-center space-x-20 h-12">
                <div className="home-button">
                    <Link to="/">
                        Home
                    </Link>
                </div>
                <div className="profile-button">
                    <Link to="/profile">
                        Profile
                    </Link>
                </div>
                <div className= "favorited-button">
                    <Link to="/favorites">
                        Favorites
                    </Link>
                </div>
                <div className="search-button">
                    <Link to="/search">
                        Search
                    </Link>
                </div>
            </nav>
        </div>
    );
}