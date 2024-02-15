import React from "react";
import { Nav, NavLink, NavMenu } from "./NavbarElements";

const Navbar = () => {
    return (
        <>
            <Nav>
                <NavMenu>
                    <NavLink to="/profile" activeStyle>
                        Profile
                    </NavLink>
                    <NavLink to="/favorites" activeStyle>
                        Favorites
                    </NavLink>
                    <NavLink to="/search" activeStyle>
                        Search
                    </NavLink>
                </NavMenu>
            </Nav>
        </>
    );
}

export default Navbar;