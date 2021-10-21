import Button from '@mui/material/Button'
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import {useState} from "react";

export function Card_channel(props) {
    const [showMe, setShowMe] = useState(false);
    function toggle(){
        setShowMe(!showMe);
    }
    return (
        <div className="card_channel_layout">
            <div className="card_channel">
                <div className="connect_group">
                    <img
                        src={props.src} width="40px" height="40px" alt=""/>
                    <label className="tickSVG" onClick={toggle} style={{
                        display: showMe ? "block" : "none"
                    }}>
                        <Button id="connectedBtn" variant="outlined" style={{
                            borderRadius: "10px",
                            paddingLeft: "1rem",
                            color: "white",
                            background: "#2198FA"
                        }}>
                            <CheckCircleIcon sx={{fontSize: 15.4, marginRight: 1}}/>Connected
                        </Button>
                    </label>
                    <label className="tickSVG" onClick={toggle} style={{
                        display: showMe ? "none" : "block"
                    }}>
                        <Button id="connectedBtn" variant="outlined" style={{
                            borderRadius: "10px",
                            paddingLeft: "1rem",
                            color: "#444444",
                            background: "#F5F6F8",
                            border: "none"
                        }}>
                            Connect
                        </Button>
                    </label>
                </div>
                <div className="information_group">
                    <span>{props.name}</span>
                    <p>{showMe ? "Connected " : "Connect "} to {props.name} </p>
                </div>
            </div>
        </div>
    )
}