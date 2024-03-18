import React, {useEffect} from 'react';
import axios from 'axios';

export default function Login(){
    useEffect(() => {
        if (sessionStorage.getItem('userData')) {
            window.location.href = '/profile';
        }
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        try {
            await axios.post('http://localhost:8000/auth/login', data, {withCredentials: true});
            // Redirect the user to the home page
            window.location.href = '/profile';
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-80 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Welcome back INFYnaut!</h1>
            </div>
            <br/>
            <div className="flex justify-center content-center align-middle">
                <form onSubmit={handleSubmit}>
                    <label htmlFor="email" className="text-neutral-50">Email:</label><br/>
                    <input type="text" id="email" name="email"></input><br/><br/>
                    <label htmlFor="password" className="text-neutral-50">Password:</label><br/>
                    <input type="text" id="password" name="password"></input><br/><br/>
                    <input type="submit" value="Log In" className="bg-violet-900 text-neutral-50 rounded-lg px-4 py-2 hover:bg-violet-950" />
                </form>
            </div>
        </div>
    )
}