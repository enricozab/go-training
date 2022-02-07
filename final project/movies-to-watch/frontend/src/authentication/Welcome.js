import { Link } from "react-router-dom";

function Welcome() {
    return (
        <>
            <div className="welcome-text">
                Welcome to MyMovies
                <br />
                <Link to="/sign-in"><button className="welcome-sign-in-button" style={{ textDecoration: "none" }}>Sign In</button></Link>
                <p>Don't have an account? <Link to="/sign-up" style={{ textDecoration: "none" }}><span>Create Account</span></Link></p>
            </div>
        </>
    );
}
  
export default Welcome;