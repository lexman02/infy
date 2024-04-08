import Navbar from './components/Navbar.jsx';
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Favorites from './pages/favorites.jsx';
import Profile from './pages/profile.jsx';
import Search from './pages/search.jsx'
import Home from './pages/Home.jsx';
import Login from "./pages/login.jsx";
import Signup from "./pages/signup.jsx";
import DetailedPost from "./pages/DetailedPost.jsx";
import logo from './img/logo.png';


export default function App() {
  return (
    <div className="flex flex-col bg-black/10 min-h-screen text-neutral-100 font-sans">
      <div className="flex justify-center">
        <a href="/">
          <img src={logo} alt="Infy Logo" className="w-48 h-48"/>
        </a>
      </div>
        <Router>
            <Routes>
                <Route exact path="/login" element={<Login/>}/>
                <Route exact path="/signup" element={<Signup/>}/>
                <Route exact path="/" element={<Home/>}/>
                <Route path="/post/:postID" element={<DetailedPost/>}/>
                <Route exact path="/profile" element={<Profile/>}/>
                <Route exact path="/favorites" element={<Favorites/>}/>
                <Route exact path="/search" element={<Search/>}/>
            </Routes>
            <Navbar />
        </Router>
    </div>
  )
}
