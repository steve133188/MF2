import {Pill} from "./Pill";
import {ThreeDotsMenu} from "./Button";
import {TextRadio2, LogoInputField} from "./Input";
import {Note} from "./Note"
import {ContactBasicInfo} from "./ContactBasicInfo";

export function ContactNote() {
    return (
        <div className="container">
            <div className="contactInfoSet">
                <ContactBasicInfo icon="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" name="Debra Patel" phone="+852 97650348" contactType="https://www.pngrepo.com/png/158412/512/whatsapp.png" pillColor="teamA" pillContent="Team A" />
                <TextRadio2 />
                <div className="contactNoteContainer">
                    <p>Note <b>3</b></p>
                    <LogoInputField>Write a note...</LogoInputField>
                    <div className="noteSet">
                        <Note pillColor="lightYellow" pillSize="roundedPill" pillContent="MF" name="Mary Foster" time="1 Jun, 11:00 AM" >Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy</Note>
                        <Note pillColor="lightBlue" pillSize="roundedPill" pillContent="AX" name="Mary Foster" time="1 Jun, 11:00 AM" >Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy</Note>
                        <Note pillColor="lightGreen" pillSize="roundedPill" pillContent="DS" name="Mary Foster" time="1 Jun, 11:00 AM" >Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy</Note>
                        <Note pillColor="lightPurple" pillSize="roundedPill" pillContent="EW" name="Mary Foster" time="1 Jun, 11:00 AM" >Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy</Note>
                        <Note pillColor="lightRed" pillSize="roundedPill" pillContent="KA" name="Mary Foster" time="1 Jun, 11:00 AM" >Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy</Note>
                    </div>
                </div>
            </div>
        </div>
    )
}