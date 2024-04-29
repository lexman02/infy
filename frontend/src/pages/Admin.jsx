import React, { useContext, useEffect, useState } from "react"
import { UserContext } from "../contexts/UserProvider"
import { FlagIcon, WrenchIcon } from "@heroicons/react/20/solid"
import UsersTab from "../components/admin/UsersTab"
import ReportedPosts from "../components/admin/ReportedPosts";

export default function Admin() {
    const [isLoading, setIsLoading] = useState(true);
    const [activeTab, setActiveTab] = useState('ReportedPosts');
    const { userData } = useContext(UserContext);

    // UseEffect for loading user data
    useEffect(() => {
        if (userData !== null) {
            setIsLoading(false);
        }
    }, [userData]);

    if (isLoading) {
        return (
            <div>
                <h1>Loading...</h1>
            </div>
        )
    }

    function renderTab(state) {
        switch (state) {
            case 'ReportedPosts':
                return <ReportedPosts />;
            case 'AdminStatus':
                return <UsersTab />;
        }
    }

    if (!isLoading && userData === null || !userData.user.isAdmin) {
        console.log("Redirecting to login page");
        window.location.href = "/login";
    }

    return (
        <div className="md:my-6 md:mx-60 flex-grow">
            <div className="flex rounded-t-lg bg-black/40 border-b-2 border-neutral-700">
                <button className={`mr-auto rounded-tl-lg text-xl hover:bg-violet-950 p-3 w-1/2 flex items-center justify-center ${activeTab === "ReportedPosts" ? "bg-violet-900/40" : ""}`} onClick={() => setActiveTab('ReportedPosts')}>Reported Posts <FlagIcon className=" pl-4 h-10 w-10" /></button>
                <button className={`mr-auto rounded-tr-lg text-xl hover:bg-violet-950 p-3 w-1/2 flex items-center justify-center ${activeTab === "AdminStatus" ? "bg-violet-900/40" : ""}`} onClick={() => setActiveTab('AdminStatus')}>Admin Status <WrenchIcon className="pl-4 h-10 w-10" /></button>
            </div>
            <div>
                {renderTab(activeTab)}
            </div>
        </div>
    )
}