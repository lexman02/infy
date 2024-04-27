import Navbar from './components/Navbar.jsx';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Watchlist from './pages/Watchlist.jsx';
import Profile from './pages/Profile.jsx';
import Search from './pages/Search.jsx'
import Home from './pages/Home.jsx';
import Login from "./pages/Login.jsx";
import Signup from "./pages/Signup.jsx";
import DetailedPost from "./pages/DetailedPost.jsx";
import logo from './img/logo.png';
import MovieDetails from './pages/MovieDetails.jsx';
import axios from 'axios';
import { useContext } from 'react';
import { UserContext } from './contexts/UserProvider';
import NotFound from './pages/NotFound.jsx';


export default function App() {

  const { userData, setUserData } = useContext(UserContext);

  function logout() {
    axios.post('http://localhost:8000/auth/logout', {}, { withCredentials: true })
      .then(() => {
        setUserData(null);
        window.location.href = '/';
      })
      .catch(error => {
        console.error(error);
      });
  }

  function handleSignup() {
    window.location.href = '/signup';
  }


  return (
    <div className="flex flex-col bg-black/10 min-h-screen text-neutral-100 font-sans">
      <div className='flex justify-end'>
        {userData === null ? <button onClick={handleSignup} className="flex justify-center w-20 bg-violet-900 text-neutral-50 rounded-lg px-4 py-2 m-2 hover:bg-violet-950">Signup</button> : null}
        {userData !== null ? <button onClick={logout} className="flex justify-center w-20 bg-violet-900 text-neutral-50 rounded-lg px-4 py-2 m-2 hover:bg-violet-950">Logout</button> : null}
      </div>
      <div className="flex justify-center">
        <a href="/">
          <img src={logo} alt="Infy Logo" className="w-48 h-48" />
        </a>
      </div>
      <Router>
        <Routes>
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/signup" element={<Signup />} />
          <Route exact path="/" element={<Home />} />
          <Route path="/post/:postID" element={<DetailedPost />} />
          <Route path="/profile/:username" element={<Profile />} />
          <Route exact path="/watchlist" element={<Watchlist />} />
          <Route exact path="/search" element={<Search />} />
          <Route exact path="/movie/:movieID" element={<MovieDetails />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
        <Navbar />
      </Router>
    </div>
  )
}
