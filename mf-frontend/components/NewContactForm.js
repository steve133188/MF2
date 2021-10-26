import {Input2} from "./Input"
import {Pill} from "./Pill"

export function NewContactForm() {
    return (
        <div className={"newContactFormContainer"}>


            <div className="infoTagContainer">
                <div>
                    <div className={"inputSetContainer"}>
                        <div className="contactDescription">
                            <h6>New Contact</h6>
                            <p>*At least one phone number or an email address is required to create the profile.</p>
                        </div>
                        <div className={"inputSet"}>
                            <Input2 title="Phone*">+852 9765 0348</Input2>
                            <Input2 title="Email*">debra.patel@gmail.com</Input2>
                        </div>
                    </div>
                    <div className={"inputSetContainer"}>
                        <div className="contactDescription">
                            <h6>Basic Information (Optional)</h6>
                        </div>
                        <div className={"inputSet"}>
                            <Input2 title="First Name">Debra</Input2>
                            <Input2 title="Last Name">Patel</Input2>
                        </div>
                        <div className={"inputSet"}>
                            <Input2 title="Birthday">Birthday</Input2>
                            <Input2 title="Country">Hong Kong</Input2>
                        </div>
                        <span className="longInput"><Input2 title="Address">233 Wan Chai Rd, Wan Chai, HK</Input2></span>

                    </div>
                </div>
                <div className={"inputSetContainer"} style={{marginTop: "170px"}}>
                    <div className="contactDescription" style={{display: "flex", maxWidth: "430px"}}>
                        <div className={"tagsGroup"}>
                            <h6>Tags & Assignee</h6>
                            <p>Tags</p>
                            <div className={"tagsGroup"}>
                                <Pill color="vip">VIP</Pill>
                                <Pill color="newCustomer">New customer</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="vvip">VVIP</Pill>
                                <Pill color="promotions">Promotions</Pill>
                                <Pill color={"add"}>+</Pill>
                            </div>
                        </div>
                        <div className={"tagsGroup"}>
                            <p>Assignee</p>
                            <div className={"tagsGroup"}>
                                <Pill color="lightYellow" size="size30">MF</Pill>
                                <Pill color="lightBlue" size="size30">VS</Pill>
                                <Pill color="lightPurple" size="size30">VS</Pill>
                                <Pill color="lightRed" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                                <Pill color="lightGreen" size="size30">VS</Pill>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}