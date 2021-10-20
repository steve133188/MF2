import Button from '@mui/material/Button'

export function Card_channel() {
    return(
        <div className="card_channel_layout">
            <div className="card_channel">
                <div className="connect_group">
                    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/6/6b/WhatsApp.svg/479px-WhatsApp.svg.png" width="40px" height="40px" alt=""/>
                    <label className="tickSVG">
                        <Button variant="outlined" style={{borderRadius: "10px", paddingLeft: "2.5rem", color: "white", background: "#2198FA"}}>
                            Connected
                        </Button>
                     </label>
                </div>
                <div className="information_group">
                    <span>Whatsapp</span>
                    <p>Connected to WhatsApp Business</p>
                </div>
            </div>
        </div>
    )
}