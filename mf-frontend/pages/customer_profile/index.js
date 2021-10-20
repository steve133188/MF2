import {CustomerProfileInfo} from "../../components/CustomerProfileInfo";
import {CustomerProfileCategory} from "../../components/CustomerProfileCategory"
import {CustomerProfileChannel} from "../../components/CustomerProfileChannel";
import {CustomerProfileActivityLog} from "../../components/CustomerProfileActivityLog";

export default function customer_profile() {
    return(
        <div>
            <CustomerProfileCategory />
            <CustomerProfileChannel />
            <CustomerProfileInfo />
            <CustomerProfileActivityLog />
        </div>
    )
}