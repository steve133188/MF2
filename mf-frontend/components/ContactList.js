import {LabelSelect} from "./Select";
import {Contacts} from "./Contacts";
import {Search2, Search3} from "./Input"
import {ContactListTopBar} from "./ContactListTopBar"
import {useState} from "react";
import {ContactFilterList} from "./ContactFilterList";

export function ContactList() {
    const [isFilter, setIsFilter] = useState(false);

    return (

            <div className="contactList">
                <ContactListTopBar isFilter={isFilter} setIsFilter={setIsFilter} />
                {isFilter?<ContactFilterList/>:<div className="contactContainerGrp">
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
                </div>}
            </div>

    )
}