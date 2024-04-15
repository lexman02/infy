import { useContext } from "react";
import { UserContext } from "../contexts/UserProvider";
import UserProfile from "../components/profile/UserProfile";

export default function Profile() {
    const { userData } = useContext(UserContext);

    return (
        <div>
            {userData && (
                <div>
                    <UserProfile userData={userData} />
                </div>
            )}
        </div>
    );
}