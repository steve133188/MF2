import {Pill} from "./Pill";
import {LabelSelect, MultipleSelectPlaceholder} from "./Select";
import {Select} from "./Input";
import {ThreeDotsMenu} from "./Button";
import {ThreeDotsHoriMenu} from "./Menu";

export function CustomerProfileInfo() {
    return (
            <div className="customerProfileInfo">
                <ThreeDotsHoriMenu />
                <div className="contactBasicInfo">
                    <img className="icon" src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" alt=""/>
                    <div className="customerName">Debra Patel</div>
                </div>
                <div className="extraInfo">
                    <div className="infoGrp">
                        <div className="infoTitle">Customer ID</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Birthday</div>
                        <div className="infoContent">1 Jun 1990, 31 years old</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Gender</div>
                        <div className="infoContent">F</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Address</div>
                        <div className="infoContent">233 Wan Chai Rd, Wan Chai, HK</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Country</div>
                        <div className="infoContent">Hong Kong</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Contact Status</div>
                        <div className="infoContent">Open</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Created Date</div>
                        <div className="infoContent">1 May, 2021</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Last Contact form You</div>
                        <div className="infoContent">September 30, 2021 11:40 PM</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Last Contact form Customer</div>
                        <div className="infoContent">October 1, 2021 10:00 AM</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Contact Owner</div>
                        <div className="infoContent">Mary Foster</div>
                    </div>
                </div>
            </div>
    )
}