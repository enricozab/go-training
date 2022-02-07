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
                        window.localStorage.removeItem("id");
                        window.location = "/";
                    }}
                    title="Sign Out"
                />
            </div>
        </div>
    )
}

export default NavBar;