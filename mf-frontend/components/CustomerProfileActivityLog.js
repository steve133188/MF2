import {Pill} from "./Pill";
import Link from 'next/link'
import {LogInputField} from "./Input";

export function CustomerProfileActivityLog() {
    return (

            <div className="customerProfileActivityLog">
                <span className="categoryName">Activity Log</span>
                <div className="activityLogSet">
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Profile Information Updated.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                    <div className="vl"></div>
                    <div className="activityLog">
                        <Pill color="lightYellow">MF</Pill>
                        <div className="logInformation">
                            <div className="activityContent">Customer has been created by <Link href="/"><a className="activityUser">Mary Foster</a></Link>.</div>
                            <div className="activityDate">7 May, 2021</div>
                        </div>
                    </div>
                </div>
                <LogInputField />
            </div>
    )
}