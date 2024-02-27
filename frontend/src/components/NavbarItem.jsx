import {Link} from "react-router-dom";
import React from "react";

export default function NavbarItem({to, icon: Icon}) {
    return (
        <Link to={to} className="flex flex-col justify-center items-center content-around space-y-1">
            <Icon className="h-6 w-6"/>
        </Link>
    );
}