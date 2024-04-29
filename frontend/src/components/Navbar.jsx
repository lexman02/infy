import React, { useContext } from "react";
import { useLocation } from "react-router-dom";
import { HomeIcon, MagnifyingGlassIcon, BookmarkIcon, UserCircleIcon } from "@heroicons/react/20/solid";
import NavbarItem from "./NavbarItem.jsx";
import { UserContext } from "../contexts/UserProvider";

export default function Navbar() {
    const { userData } = useContext(UserContext);
    const location = useLocation();

    if (location.pathname === "/login" || location.pathname === "/signup") {
        return null;
    }

    return (
        <div className="text-neutral-50 sticky bottom-0 left-0 z-50 w-full h-12 bg-black/20 md:mt-4">
            <nav className="flex items-center justify-center space-x-14 md:space-x-20 h-full">
                <NavbarItem to="/" icon={HomeIcon} />
                <NavbarItem to="/search" icon={MagnifyingGlassIcon} />
                {userData ? <NavbarItem to={`/profile/${userData.user.username}`} icon={UserCircleIcon} avatar={userData.user.profile.avatar} /> : <NavbarItem to="/login" icon={UserCircleIcon} />}
            </nav>
        </div>
    );
}