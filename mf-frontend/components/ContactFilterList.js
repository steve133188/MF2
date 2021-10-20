import {ContactListTopBar} from "./ContactListTopBar"
import {CheckboxGroup1, CheckboxGroup2, Checkbox3} from "./Input";
import {LabelSelect, LabelSelect2} from "./Select";
import {Pill} from "./Pill";
import {CancelButton, NormalButton, NormalButton2, NormalButton3} from "./Button";
import {Checkbox1, Checkbox2} from "./Checkbox";

export function ContactFilterList() {
    return (
        <div className="container">
            <div className="contactList">
                <ContactListTopBar/>
                <div className="filterArea">
                    <CheckboxGroup1 title="Filter">
                        <Checkbox1 checked="checked">Unread</Checkbox1>
                        <Checkbox1>Unassigned</Checkbox1>
                        <Checkbox1>ChatBot</Checkbox1>
                    </CheckboxGroup1>
                    <CheckboxGroup2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            All Channel
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png" checked="checked">
                            WhatsApp
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            WhatsApp Business API
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            Messager
                        </Checkbox2>
                        <Checkbox2 src="https://clipart.info/images/ccovers/1499955335whatsapp-icon-logo-png.png">
                            WeChat
                        </Checkbox2>
                    </CheckboxGroup2>
                    <div className="agentFilter">
                        <p>Agent</p>
                        <LabelSelect2/>
                        <div className="agentGroup">
                            <Pill color="lightYellow" size="size30">MF</Pill>
                            <Pill color="lightBlue" size="size30">MF</Pill>
                        </div>
                    </div>
                    <div className="tagFilter">
                        <p>Tag</p>
                        <div className="tagGroup">
                            <Pill color="vip">VIP</Pill>
                            <Checkbox3/>
                        </div>
                    </div>
                    <div className="buttonGrp">
                        <NormalButton2>Confirm</ NormalButton2>
                        <CancelButton>Cancel</ CancelButton>
                    </div>
                </div>
            </div>
        </div>
    )
}