import React from 'react';
import axios from 'axios';

export default function Signup(){
    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        try {
            await axios.post('http://localhost:8000/auth/signup', data);
            // Redirect the user to the login page
            window.location.href = '/login';
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-80 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Fill out your INFYnaut Application!</h1>
            </div>
            <br/>
            <div className="flex justify-center align-middle">
                <form onSubmit={handleSubmit}>
                    <label htmlFor="username" className="text-neutral-50">Username:</label><br/>
                    <input type="text" id="username" name="username" className="hover:bg-neutral-400"></input><br/><br/>
                    <label htmlFor="password" className="text-neutral-50">Password:</label><br/>
                    <input type="password" id="password" name="password" className="hover:bg-neutral-400"></input><br/><br/>
                    <label htmlFor="email" className="text-neutral-50">Email:</label><br/>
                    <input type="text" id="email" name="email" className="hover:bg-neutral-400"></input><br/><br/>
                    <input type="submit" value="Submit" className="bg-violet-900 text-neutral-50 px-4 py-2 rounded-lg hover:bg-violet-950" />
                </form>
            </div>
        </div>
    )
}