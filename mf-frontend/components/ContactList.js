import {LabelSelect} from "./Select";
import {Contacts} from "./Contacts";
import {Search2} from "./Input"
import {ContactListTopBar} from "./ContactListTopBar"

export function ContactList() {
    return (

            <div className="contactList">
                <ContactListTopBar />
                <div className="contactContainerGrp">
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                    <Contacts/>
                </div>
            </div>

    )
}