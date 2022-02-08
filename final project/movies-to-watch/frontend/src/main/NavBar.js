import axios from "axios";
import MenuIcon from '../assets/menu-icon.png';
import SignOutIcon from '../assets/sign-out-icon.png';

function NavBar(props) {
    return (
        <div className="navbar-main">
            <div className='navbar-left'>
                <img src={MenuIcon} alt="Menu Icon" className="menu-icon" onClick={() => props.setShowSidebar(!props.showSidebar)} title="Menu" />
                <div className="navbar-logo-container" onClick={() => window.location = "/"}>
                    <img src="logo512.png" alt="MyMovies Logo" className="navbar-logo" />
                    <div className="navbar-logo-text">MyMovies</div>
                </div>
            </div>
            <div className='navbar-right'>
                <p style={{ color: "white", marginRight: 20 }}>Welcome, <b>{props.user.Username}</b>!</p>
                <img 
                    src={SignOutIcon} 
                    alt="Sign Out Icon"
                    className="sign-out-icon"
                    onClick={() => {
                        const log = {
                            Description: "User logged out"
                        }

                        axios.post(`http://localhost:8080/logs/${props.user.Id}`, log)
                        .then(response => {
                            // console.log(response.data);
                            
                            window.localStorage.removeItem("id");
                            window.location = "/";
                        })
                        .catch(error => {
                            alert("Error: An unexpected error occurred. Please try again later.");
                            console.log(error);
                        });
                    }}
                    title="Sign Out"
                />
            </div>
        </div>
    )
}

export default NavBar;