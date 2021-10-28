import {ContactList} from "../../components/ContactList";
import {ChatArea} from "../../components/ChatArea"
import {ContactInfo} from "../../components/ContactInfo";
import {ContactFilterList} from "../../components/ContactFilterList";
import {ContactNote} from "../../components/ContactNote";
import {ContactFile} from "../../components/ContactFile"

export default function Live_chat() {
    return (
        <div className="live_chat-layout">
            <ContactList/>
            <ChatArea/>
            <ContactFile/>
        </div>
    )
}