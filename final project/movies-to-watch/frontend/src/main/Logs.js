import { useState, useEffect } from "react";
import axios from "axios";

function Logs(props) {
    const [logs, setLogs] = useState([]);

    useEffect(() => {
        axios.get(`http://localhost:8080/logs/${props.user.Id}`)
        .then(response => {
            // console.log(response.data);

            setLogs(response.data);
        })
        .catch(error => {
            alert("Error: An unexpected error occurred. Please try again later.");
            console.log(error);
        });
    }, []);

    return (
        <>
            <div className="logs-container">
                <div className="logs-header-container">
                    <p>Activity Log</p>
                </div>
                <br />
                <div style={{ maxHeight: "calc(100vh - 180px)", overflowY: "auto" }}>
                    <table>
                        <tr>
                            <th style={{ width: "20%" }}>Date</th>
                            <th style={{ width: "80%" }}>Description</th>
                        </tr>
                        {logs.map(log => {
                            return <tr>
                                <td>{log.Date}</td>
                                <td>{log.Description}</td>
                            </tr>
                        })}
                    </table>
                </div>
            </div>
        </>
    )
}

export default Logs;