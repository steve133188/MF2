import {LineChart1, MultipleLineChart} from "../../components/LineChart1";
import {SingleSelect2} from "../../components/Select";
import {EnhancedTable} from "../../components/Table";
import {EnhancedTable2} from "../../components/EnhancedTable2";
import {EnhancedTable3} from "../../components/EnhancedTable3";

export default function Dashboard() {


    return (
        <div className="dashboard-layout">
            <div className="navbarPurple">
            </div>
            <div className="chartGroup">
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"All Contacts"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Contacts"} total={"32"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Active Contacts"} data={[12, 17, 19, 22, 24, 20, 18, 26, 20, 17, 15]} yaxis={"Contacts"} total={"32"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"Total Messages Sent"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Messages"} total={"32"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Total Messages Received"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Messages"} total={"32"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"All Contacts"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Enquiries"} total={"32"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Newly Added Contacts"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Contacts"} total={"32"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"Average Response Time"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Mintes"} total={"32"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Most Communication Hours"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Hours"} total={"32"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="tableSet">
                        <div className="dashboardColumn" style={{width: "55%"}}><EnhancedTable3/></div>
                    </div>
                </div>
            </div>



            <div className="navbarPurple">

            </div>
            <div className="chartGroup">
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                    <div className="dashboardColumn"><MultipleLineChart/></div>
                </div>
                <div className="dashboardRow">
                    <div className="tableSet">
                        <div className="dashboardColumn" style={{width:"55%"}}><EnhancedTable2/></div>
                    </div>
                </div>
            </div>
        </div>

    )
}