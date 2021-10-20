import {Pill, StatusPill} from "./Pill";
import {ContactType} from "./ContactType";
import {Badge} from "./Badge";
import {Checkbox1, SingleBox} from "./Checkbox";
import Avatar from '@mui/material/Avatar';

export function ContactTable() {
    return(
        <div className="contactTable">
            <table className="table">
                <thead>
                <tr className="headTr">
                    <th className="trCustomerID">Customer ID</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Channel</th>
                    <th>Tags</th>
                    <th>Assignee</th>
                </tr>
                </thead>
                <tbody>
                <tr className="bodyTr">
                    <td>0000001</td>
                    <td>Jacosdas dasdasdb</td>
                    <td>Thorntsadaasdason@com</td>
                    <td className="channel"><ContactType />lori.foster@mail.com</td>
                    <td><span className="tagsGroup"><Pill color="vip">VIP</Pill></span></td>
                    <td>Mary Foster</td>
                </tr>
                <tr>
                    <td>0000002</td>
                    <td>Jacosdas dasdasdb</td>
                    <td>Thorntsadaasdason@com</td>
                    <td className="channel"><ContactType />lori.foster@mail.com</td>
                    <td><span className="tagsGroup"><Pill color="vip">VIP</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill></span></td>
                    <td className="assignee">Mary Foster</td>
                </tr>
                </tbody>
            </table>
        </div>
    )
}

export function BroadcastTable() {
    return(
        <div className="broadcastTable">
            <table className="table">
                <thead>
                    <tr className="headTr">
                        <th className="trID">Name</th>
                        <th>Period</th>
                        <th>Group</th>
                        <th>Status</th>
                        <th>Created By</th>
                        <th>Created Date</th>
                        <th>Description</th>
                    </tr>
                </thead>
                <tbody>
                    <tr className="bodyTr">
                        <td>Broadcast 1</td>
                        <td>Sep 30, 2021 7:00 AM - Oct 30, 2021 7:00 AM</td>
                        <td><Badge color="gp1">Group1</Badge></td>
                        <td><StatusPill color="statusActive">Active</StatusPill></td>
                        <td>Lorem Ipsum</td>
                        <td>Mary Foster</td>
                        <td>s</td>
                    </tr>
                    <tr>
                        <td>Broadcast 2</td>
                        <td>Sep 30, 2021 7:00 AM - Oct 30, 2021 7:00 AM</td>
                        <td><Badge color="gp2">Group2</Badge></td>
                        <td><StatusPill color="statusPending">Pending</StatusPill></td>
                        <td>s</td>
                        <td>s</td>
                        <td>Mary Foster</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}

export function NormalTable(props) {
    return (
        <table className="normalTable">
            <tr>
                <th className={props.classname}><SingleBox></SingleBox></th>
                <th>Customer ID</th>
                <th>Name</th>
                <th>Team</th>
                <th>Channel</th>
                <th>Tags</th>
                <th>Assignee</th>
            </tr>
            {props.children}
        </table>
    )
}