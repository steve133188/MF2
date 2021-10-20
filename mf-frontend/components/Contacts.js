import {Pill} from "../components/Pill"
import {ContactType} from "./ContactType";
import {IconWithPill} from "./Icon";

export function Contacts() {
    return (
        <div className="contactContainer">
            <div className="contact">
                <div className="addToFavouriteBtn">★</div>
                <div className="contactInfoAndTime">
                    <div className="contactInfoGrp">
                        <IconWithPill src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>
                        <div className="contactChatInfo">
                            <div className="contactName">Debra Patel<ContactType/></div>
                            <Pill color="teamB">Team A</Pill>
                        </div>
                    </div>
                    <div className="msgTimeGrp">
                        <div className="lastOnlineTime">03:45 PM</div>
                        <div className="contactExtraTime">00:00:00</div>
                    </div>
                </div>
            </div>
        </div>
    )
}