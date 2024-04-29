import React, { useEffect, useState } from "react"
import axios from "axios"

export default function UsersTab() {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        axios.get(`http://localhost:8000/admin/users`, { withCredentials: true })
            .then(response => {
                if (response.data.users.length > 0) {
                    setUsers(response.data.users);
                } else {
                    setUsers(0);
                }
            })
            .catch(error => {
                console.error(error);
            });
    }, []);

    async function toggleAdmin(userID) {
        await axios.put(`http://localhost:8000/admin/users/${userID}`, {}, { withCredentials: true })
            .then(() => {
                const updatedUsers = users.map(user => {
                    if (user.id === userID) {
                        user.isAdmin = !user.isAdmin;
                    }
                    return user;
                });
                setUsers(updatedUsers);
            })
            .catch(error => {
                console.error(error);
            });
    }

    return (
        <div className="divide-y divide-neutral-500">
            {users.length > 0 ? (
                users.map(user =>
                    <div key={user.id} className="flex justify-between items-center space-x-8 bg-black/40 p-4 text-neutral-100 last:rounded-b-lg">
                        <div className="flex flex-col">
                            <h1>{user.username}</h1>
                            <h1>Email: {user.email}</h1>
                            <h1>Admin Status: {user.isAdmin ? "Admin" : "Not Admin"}</h1>
                        </div>
                        <div>
                            {user.isAdmin ? (
                                <button className="bg-red-700 p-2 rounded-lg" onClick={() => toggleAdmin(user.id)}>Remove Admin</button>
                            ) : (
                                <button className="bg-green-700 p-2 rounded-lg" onClick={() => toggleAdmin(user.id)}>Make Admin</button>
                            )}
                        </div>
                    </div>
                )
            ) : (
                <div className="p-4 text-neutral-300 font-medium text-lg bg-black/40 rounded-b-lg">
                    No users found!
                </div>
            )}
        </div>
    )
}