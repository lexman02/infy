import { Link } from "react-router-dom";
import React from "react";

export default function NavbarItem({ to, icon: Icon, avatar }) {
    const avatarIcon = () => {
        if (avatar) {
            return (<img src={`http://localhost:8000/avatars/${avatar}`} className="w-7 h-7 rounded-full" />);
        } else {
            return (<Icon className="h-7 w-7 md:h-6 md:w-6" />);
        }
    }

    return (
        <Link to={to} className="flex flex-col justify-center items-center content-around space-y-1">
            {avatarIcon()}
        </Link>
    );
}