import {Card_channel} from "../../components/Cards";

export default function Integrations() {
    return (
        <div className="integrations-layout">
            <div className="container-fluid cardChannelGroupContainer">
                <div className="cardChannelGroup">
                    <h1 style={{fontSize: "24px", fontWeight: "800", marginBottom: "33px"}}>My Channels</h1>
                    <div className="row cardContainer">
                        <Card_channel/>
                    </div>
                </div>
                <div className="cardChannelGroup">
                    <h1 style={{fontSize: "24px", fontWeight: "800", marginBottom: "33px"}}>Channels</h1>
                    <div className="row cardContainer">
                        <Card_channel/>
                        <Card_channel/>
                        <Card_channel/>
                        <Card_channel/>
                        <Card_channel/>
                    </div>
                </div>
            </div>
        </div>
    )
}