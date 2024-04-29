import React from 'react';
import axios from 'axios';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import Snackbar from '@mui/material/Snackbar';
import Alert from '@mui/material/Alert';

export default function Signup() {
    const [errorMessage, setErrorMessage] = useState('');

    const handleCloseError = () => {
        setErrorMessage('');
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.target));
        try {
            await axios.post('http://localhost:8000/auth/signup', data, { withCredentials: true });
            // Redirect the user to the profile page
            window.location.href = '/profile';
        } catch (error) {
            console.error(error);
            if (error.response && error.response.data && error.response.data.message) {
                // If the error contains a specific message, set that as the errorMessage
                setErrorMessage(error.response.data.message);
            } else {
                // If no specific message is available, set a generic error message
                setErrorMessage('An error occurred while signing up, please check required fields and try again.');
            }

        }
    }

    return (
        <div className="flex flex-col justify-center items-center mb-10">
            <div className="w-full bg-neutral-800 rounded-lg shadow-lg shadow-violet-900/50 md:mt-0 sm:max-w-md xl:p-0">
                <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <h1 className="text-xl font-bold leading-tight tracking-tight text-neutral-50 md:text">
                        Join the InfyVerse!
                    </h1>
                    <form onSubmit={handleSubmit} className="space-y-4 md:space-y-6">
                        {/*Username field */}
                        <div>
                            <label htmlFor="username" className="block mb-2 text-sm font-medium text-gray">Username:</label>
                            <input type="text" id="username" name="username" className="border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5"></input>
                        </div>

                        {/*Email field */}
                        <div>
                            <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray">Email:</label>
                            <input type="text" id="email" name="email" className="border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5"></input>
                        </div>

                        <div className="flex justify-between items-center">
                            {/*First Name field */}
                            <div>
                                <label htmlFor="first_name" className='block mb-2 text-sm font-medium text-gray'>First Name:</label>
                                <input type="text" id="first_name" name="first_name" className='border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5'></input>
                            </div>

                            {/*Last Name field */}
                            <div>
                                <label htmlFor="last_name" className='block mb-2 text-sm font-medium text-gray'>Last Name:</label>
                                <input type="text" id="last_name" name="last_name" className='border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5'></input>
                            </div>
                        </div>

                        {/*Date of Birth field*/}
                        <div>
                            <label htmlFor="date_of_birth" className='block mb-2 text-sm font-medium text-gray'>Date of Birth:</label>
                            <input type="date" id="date_of_birth" name="date_of_birth" className='border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5'></input>
                        </div>

                        {/*Password field */}
                        <div>
                            <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray">Password:</label>
                            <input type="password" id="password" name="password" className="border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5"></input>
                        </div>

                        {/*Confirm Password field */}
                        <div>
                            <label htmlFor="confirm_password" className="block mb-2 text-sm font-medium text-gray">Confirm Password:</label>
                            <input type="password" id="confirm_password" name="confirm_password" className="border border-gray-400 text-gray-900 sm:text-sm rounded-lg focus:ring-violet-500 focus:border-violet-500 block w-full p-2.5"></input>
                        </div>

                        {/*Submit button */}
                        <button type="submit" value="Submit" className="w-full text-white bg-violet-800 hover:bg-violet-700 focus:ring-4 focus:outline-none focus:ring-violet-400 font-medium rounded-lg text-sm px-5 py-2.5 text-center">Sign Up</button>
                        <p className="text-sm font-light text-gray-500">
                            Already have an account? <Link to="/login" className="font-medium text-violet-500 hover:underline">Sign in</Link>
                        </p>
                    </form>
                </div>
            </div >

            <Snackbar open={!!errorMessage} autoHideDuration={6000} onClose={handleCloseError}>
                <Alert elevation={6} variant="filled" severity="error" onClose={handleCloseError}>
                    {errorMessage}
                </Alert>
            </Snackbar>

        </div >
    )
}