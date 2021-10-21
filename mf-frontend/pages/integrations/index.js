import {Card_channel} from "../../components/Cards";

export default function Integrations() {
    return (
        <div className="integrations-layout">
            <div className="leftMenu">MENU</div>
            <div className="rightContent">
                <div className="navbar">NAVBAR</div>
                <div className="container-fluid cardChannelGroupContainer">
                    <div className="cardChannelGroup">
                        <h1 style={{fontSize: "24px", fontWeight: "800", marginBottom: "33px"}}>My Channels</h1>
                        <div className="row cardContainer">
                            <Card_channel src="Group 4965.svg" name="WhatsApp"/>
                        </div>
                    </div>
                    <div className="cardChannelGroup">
                        <h1 style={{fontSize: "24px", fontWeight: "800", marginBottom: "33px"}}>Channels</h1>
                        <div className="row cardContainer">
                            <Card_channel src="Group 5167.svg" name="WhatsApp Business API"/>
                            <Card_channel src="Group 4965.svg" name="WeChat"/>
                            <Card_channel src="Group 4965.svg" name="WeChat"/>
                            <Card_channel src="Group 4965.svg" name="WeChat"/>
                            <Card_channel src="Group 4965.svg" name="WeChat"/>
                            <Card_channel src="Group 4965.svg" name="WeChat"/>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    )
}