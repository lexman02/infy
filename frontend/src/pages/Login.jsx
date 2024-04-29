import React, { useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserProvider';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Login() {
    const { userData } = useContext(UserContext);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    // Send user to profile once logged in
    useEffect(() => {
        if (userData !== null) {
            window.location.href = '/profile';
        }
    });

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        await axios.post('http://localhost:8000/auth/login', data, { withCredentials: true })
            .then(() => {
                window.location.href = '/';
            })
            .catch((error) => {
                console.error(error);
                if (error.response && error.response.data && error.response.data.error) {
                    // If the error contains a specific message, set that as the errorMessage
                    setErrorMessage(error.response.data.error);
                } else {
                    // If no specific message is available, set a generic error message
                    setErrorMessage('An error occurred while logging in. Please check credentials and try again.');
                }
            });
    };

    return (
        <div className="flex flex-col justify-center items-center">
            <div className="w-full bg-neutral-800 rounded-lg shadow-lg shadow-violet-900/50 md:mt-0 sm:max-w-md xl:p-0">
                <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 className="text-xl font-bold leading-tight tracking-tight text-neutral-50 md:text">
                        Welcome back InfyNaut!
                    </h1>
                    <form onSubmit={handleSubmit} className="space-y-4 md:space-y-6">
                        {/*Email field */}
                        <div>
                            <label htmlFor="email" className="block mb-2 text-sm font-medium">Email</label>
                            <input type="text" id="email" name="email" className="border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5"></input>
                        </div>

                        {/*Password field */}
                        <div>
                            <label htmlFor="password" className="block mb-2 text-sm font-medium">Password</label>
                            <input type="password" id="password" name="password" className="border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5"></input>
                        </div>

                        {/*Submit field */}
                        <button type="submit" value="Log In" className="w-full text-white bg-violet-800 hover:bg-violet-700 focus:ring-4 focus:outline-none focus:ring-violet-400 font-medium rounded-lg text-sm px-5 py-2.5 text-center">Sign In</button>
                    </form>
                </div>
            </div>

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>
        </div>
    )
}