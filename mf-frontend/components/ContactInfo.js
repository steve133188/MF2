import {Pill} from "./Pill";
import {NormalButton3, NormalButton2, ThreeDotsMenu} from "./Button";
import {TextRadio2} from "./Input";
import {ContactBasicInfo} from "./ContactBasicInfo";

export function ContactInfo() {
    return (

            <div className="contactInfoSet">
                <ContactBasicInfo icon="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" name="Debra Patel" phone="+852 97650348" contactType="https://www.pngrepo.com/png/158412/512/whatsapp.png" pillColor="teamA" pillContent="Team A" />
                <TextRadio2 />
                <div className="extraInfo">
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Birthday</div>
                        <div className="infoContent">1 Jun 1990, 31 years old</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                    <div className="infoGrp">
                        <div className="infoTitle">Email</div>
                        <div className="infoContent">debra.patel@gmail.com</div>
                    </div>
                </div>
                <div className="assigneeAndTagsGroup">
                    <div className="tagsArea">
                        <div className="tagsTitle">Assignee</div>
                        <div className="assigneeTagsGroup">
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="add" size="size30">+</Pill>
                        </div>
                    </div>
                    <div className="tagsArea">
                        <div className="tagsTitle">Tags</div>
                        <div className="tagsGrp">
                            <Pill color="vip2">VIP</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="purple">New Customer</Pill>
                            <Pill color="add">+</Pill>
                        </div>
                    </div>
                </div>
            </div>

    )
}