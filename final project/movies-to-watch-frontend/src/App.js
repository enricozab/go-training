import { useState, useEffect } from "react";
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import axios from 'axios';
import { css } from "@emotion/react";
import { MoonLoader } from "react-spinners";
import Welcome from './authentication/Welcome';
import SignIn from "./authentication/SignIn";
import SignUp from "./authentication/SignUp";
import NavBar from './main/NavBar';
import Sidebar from './main/Sidebar';
import Movies from './main/Movies'
import Logs from "./main/Logs";
import './App.css';

const override = css`
  display: block;
  margin: auto;
  background-color: transparent;
`;

function App(props) {
    const [isLoading, setIsLoading] = useState(true);
    const [isLogin, setIsLogin] = useState(false);
    const [users, setUsers] = useState([]);
    const [user, setUser] = useState(null);

    const [showSidebar, setShowSidebar] = useState(false);
    const [activeMenu, setActiveMenu] = useState("movies");

    useEffect(() => {
        setIsLoading(true);

        axios.get(`http://localhost:8080/users`)
        .then(response => {
            // console.log(response.data);
            setUsers(response.data);

            let isExisting = false;

            if (![null, undefined].includes(window.localStorage.getItem("id"))) {
                response.data.forEach(user => {
                    if (user.Id === parseInt(window.localStorage.getItem("id"))) {
                        isExisting = true;

                        setUser(user)
                        setIsLogin(true)
                    }
                });
            }

            if (!isExisting) {
                window.localStorage.removeItem("id");
            }
        })
        .catch(error => {
            console.log(error);

            alert("Error: An unexpected error occurred. Please try again later.");
        })
        .finally(() => {
            setIsLoading(false);
        });
    }, []);

    useEffect(() => {
        if (window.location.pathname === "/logs") {
            setActiveMenu("logs");
        }
    }, [])

    return (
        isLoading ? <div className="loader-container"><MoonLoader color={"#CE3300"} loading={isLoading} css={override} size={50} /></div> :
        <BrowserRouter basename={props.baseUrl}>
            <div className="App">
                {!isLogin
                ? <div className="welcome-main">
                    <div className="welcome-layout">
                        <img src="logo512.png" alt="MyMovies Logo" className="welcome-logo" onClick={() => window.location = "/"} />
                        <Routes>
                            <Route path="/" element={<Welcome />} />
                            <Route path="/sign-in" element={<SignIn users={users} />} />
                            <Route path="/sign-up" element={<SignUp users={users} />} />
                            <Route path="*" element={<Navigate to="/" />} />
                        </Routes>
                    </div>
                </div>
                : <>
                    <NavBar user={user} showSidebar={showSidebar} setShowSidebar={setShowSidebar} />

                    <div className='main-container' onClick={() => setShowSidebar(false)}>
                        <Routes>
                            <Route path="/my-movie-list" element={<Movies user={user} />} />
                            <Route path="/logs" element={<Logs user={user} />} />
                            <Route path="*" element={<Navigate to="/my-movie-list" />} />
                        </Routes>
                    </div>

                    <Sidebar showSidebar={showSidebar} activeMenu={activeMenu} setActiveMenu={setActiveMenu} />
                </>}
            </div>
        </BrowserRouter>
    );
}

export default App;