import { Link } from "react-router-dom";

function Sidebar(props) {
    return (
        <div className={`sidebar-container ${props.showSidebar ? "show" : "hide"}`}>
            <div className="sidebar-item-container">
                <Link to="/my-movie-list" style={{ textDecoration: "none" }}><span className={`${props.activeMenu === "movies" && "bold"}`} onClick={() => props.setActiveMenu("movies")}>My Movie List</span></Link>
            </div>
            <div className="sidebar-item-container">
                <Link to="/logs" style={{ textDecoration: "none" }}><span className={`${props.activeMenu === "logs" && "bold"}`} onClick={() => props.setActiveMenu("logs")}>Activity Log</span></Link>
            </div>
        </div>
    )
}

export default Sidebar;