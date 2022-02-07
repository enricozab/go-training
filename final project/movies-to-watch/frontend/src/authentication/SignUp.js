import { useState } from "react";
import { Link } from "react-router-dom";
import axios from 'axios';

function SignUp(props) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const signUp = () => {
        if (username !== "" && password !== "") {
            let checkUser = props.users.find(user => user.Username === username);

            if (![null, undefined].includes(checkUser)) {
                alert("Error: Username already exists. Please try with another one.");
            } else {
                const user = {
                    Username: username,
                    Password: password
                };

                axios.post(`http://localhost:8080/users`, user)
                .then(response => {
                    // console.log(response.data);
                    
                    let users = [...response.data].sort((a, b) => a.Id > b.Id ? 1 : -1)
                    let newUser = users[users.length - 1];

                    alert("Success: Account created successfully.");
                    
                    window.localStorage.setItem("id", newUser.Id);
                    window.location = "/";
                })
                .catch(error => {
                    alert("Error: An unexpected error occurred. Please try again later.");
                    console.log(error);
                });
            }
        } else {
            alert("Error: Please fill out all fields.");
        }
    }

    return(
        <>
            <div className="sign-in-container">
                <p className="sign-in-container-header">Create Account</p>
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
                <button className="sign-in-button" onClick={() => signUp()}>Sign Up</button>
                <br />
                <p style={{ fontSize: 12 }}>Already have an account? <Link to="/sign-in" style={{ textDecoration: "none" }}><span style={{ color: "#CE3300", fontWeight: "bold" }}>Sign In</span></Link></p>
            </div>
        </>
    )
}

export default SignUp;