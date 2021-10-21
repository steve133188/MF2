import {Pill} from "./Pill"
import {Badge} from "./Badge"
import {NormalButton} from "./Button";

export function CustomerProfileChannel() {
    return (
            <div className="customerProfileChannel">
                <div className="customerChannelContainer">
                    <div className="categoryHeader">
                        <span className="categoryName">Channels</span>
                        <NormalButton>Merge</NormalButton>
                    </div>
                    <div className="channelGrp">
                        <div className="channelSet">
                            <div className="channelName"><img
                                src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                alt=""/>WhatsApp
                            </div>
                        </div>
                        <span className="contactAddress">+852 97650348</span>
                    </div>
                    <div className="channelGrp">
                        <div className="channelSet">
                            <div className="channelName">
                                <img
                                src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                alt=""/>WhatsApp
                            </div>
                        </div>
                        <span className="contactAddress">+852 97650348</span>
                    </div>
                    <div className="pillBadgeContainer">

                    </div>

                    <div className="customerTagsContainer">
                        <div className="categoryHeader">
                            <span className="categoryName">Tags</span>
                            <NormalButton>+</NormalButton>
                        </div>
                        <div className="pillBadgeContainer">
                            <Pill color="vip">VIP</Pill>
                            <Pill color="newCustomer">New Customer</Pill>

                        </div>
                    </div>
                </div>
            </div>
    )
 }