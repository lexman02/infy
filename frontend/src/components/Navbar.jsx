import React from "react";
import {useLocation} from "react-router-dom";
import {HomeIcon, MagnifyingGlassIcon, StarIcon, UserCircleIcon} from "@heroicons/react/20/solid";
import NavbarItem from "./NavbarItem.jsx";

export default function Navbar(){
    const location = useLocation();
    const data = sessionStorage.getItem('userData');

    if (location.pathname === "/login" || location.pathname === "/signup") {
        return null;
    }

    const Avatar = () => {
        if (!data) {
            return false;
        }
    }

    return (
        <div className="text-neutral-50 fixed bottom-0 left-0 right-0 p-2">
            <nav className="flex items-center justify-center space-x-20 h-12">
                <NavbarItem to="/" icon={HomeIcon}/>
                <NavbarItem to="/search" icon={MagnifyingGlassIcon}/>
                <NavbarItem to="/favorites" icon={StarIcon}/>
                {Avatar() && <NavbarItem to="/profile" icon={UserCircleIcon}/>}
                {!Avatar() && <NavbarItem to="/login" icon={UserCircleIcon}/>}
            </nav>
        </div>
    );
}