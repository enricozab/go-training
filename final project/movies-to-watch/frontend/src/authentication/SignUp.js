import { useState } from "react";
import { Link } from "react-router-dom";
import axios from 'axios';

function SignUp(props) {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const signUp = () => {
        if (username !== "" && password !== "") {
            const user = {
                Username: username,
                Password: password
            };

            axios.post(`http://localhost:8080/users`, user)
            .then(response => {
                // console.log(response.data);
                
                let users = [...response.data].sort((a, b) => a.Id > b.Id ? 1 : -1)
                let newUser = users[users.length - 1];
                
                const log = {
                    Description: "User logged in"
                }

                alert("Success: Account created successfully.");

                setTimeout(() => {
                    axios.post(`http://localhost:8080/logs/${newUser.Id}`, log)
                    .then(response => {
                        // console.log(response.data);

                        window.localStorage.setItem("id", newUser.Id);
                        window.location = "/";
                    })
                    .catch(error => {
                        alert("Error: An unexpected error occurred. Please try again later.");
                        console.log(error);
                    });
                }, 5)
            })
            .catch(error => {
                console.log(error);

                let errorMessage = error?.response?.data ?? "Error: An unexpected error occurred. Please try again later."
                
                alert(errorMessage);
            });
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