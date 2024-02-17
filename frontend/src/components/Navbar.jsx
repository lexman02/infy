import React from "react";
import {Link} from "react-router-dom";

export default function Navbar(){
    return (
        <div>
            <nav>
                <div>
                    <Link to="/profile">
                        Profile
                    </Link>
                    <Link to="/favorites">
                        Favorites
                    </Link>
                    <Link to="/search">
                        Search
                    </Link>
                </div>
            </nav>
        </div>
    );
}