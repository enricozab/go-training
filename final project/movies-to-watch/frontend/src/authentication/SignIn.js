import { useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

function SignIn(props) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const signIn = () => {
        if (username !== "" && password !== "") {
            let isExisting = false;

            props.users.forEach(user => {
                if (user.Username === username && user.Password === password) {
                    isExisting = true
                    
                    const log = {
                        Description: "User logged in"
                    }

                    axios.post(`http://localhost:8080/logs/${user.Id}`, log)
                    .then(response => {
                        // console.log(response.data);
                        
                        window.localStorage.setItem("id", user.Id);
                        window.location = "/";
                    })
                    .catch(error => {
                        alert("Error: An unexpected error occurred. Please try again later.");
                        console.log(error);
                    });
                }
            });

            if (!isExisting) {
                setUsername("");
                setPassword("");
                alert("Error: Incorrect username/password. Please try again.");
            }
        } else {
            alert("Error: Please fill out all fields.");
        }
    }

    return(
        <>
            <div className="sign-in-container">
                <p className="sign-in-container-header">Sign In</p>
                <br />
                <br />
                <p className="input-header-style">Username</p>
                <input className="input-style" placeholder="Enter username" value={username} onChange={(e) => setUsername(e.target.value)} />
                <br />
                <br />
                <p className="input-header-style">Password</p>
                <input className="input-style" type="password" placeholder="Enter password" value={password} onChange={(e) => setPassword(e.target.value)} />
                <br />
                <br />
                <br />
                <button className="sign-in-button" onClick={() => signIn()}>Sign In</button>
                <br />
                <Link to="/" style={{ textDecoration: "none" }}><p className="sign-in-back">Back</p></Link>
            </div>
        </>
    )
}

export default SignIn;