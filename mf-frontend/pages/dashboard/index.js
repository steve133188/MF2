import {LineChart1, MultipleLineChart} from "../../components/LineChart1";
import {SingleSelect2} from "../../components/Select";
import {EnhancedTable} from "../../components/Table";
import {EnhancedTable2} from "../../components/EnhancedTable2";
import {EnhancedTable3} from "../../components/EnhancedTable3";
import {LineChartCard} from "../../components/Cards";

export default function Dashboard() {


    return (
        <div className="dashboard-layout">
            <div className="navbarPurple">
            </div>
            <div className="chartGroup">
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"All Contacts"} data={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} yaxis={"Contacts"} total={"34"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Active Contacts"} data={[12, 17, 19, 22, 24, 20, 18, 26, 20, 17, 15, 19]} yaxis={"Contacts"} total={"19"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"Total Messages Sent"} data={[23, 38, 30, 17, 26, 18, 34, 13, 19, 39, 22, 14]} yaxis={"Messages"} total={"14"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Total Messages Received"} data={[17, 18, 17, 13, 40, 17, 36, 33, 25, 34, 36, 15]} yaxis={"Messages"} total={"15"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"All Contacts"} data={[40, 24, 37, 39, 21, 14, 19, 36, 27, 31, 28, 14]} yaxis={"Enquiries"} total={"14"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Newly Added Contacts"} data={[21, 18, 17, 35, 38, 16, 40, 18, 12, 24, 30, 20]} yaxis={"Contacts"} total={"20"} percentage={"+5%"} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><LineChart1 title={"Average Response Time"} data={[16, 24, 23, 36, 19, 20, 25, 29, 29, 22, 34, 37]} yaxis={"Mintes"} total={"37"} percentage={"+5%"} /></div>
                    <div className="dashboardColumn"><LineChart1 title={"Most Communication Hours"} data={[28, 30, 17, 18, 36, 13, 23, 36, 34, 23, 15, 26]} yaxis={"Hours"} total={"26"} percentage={"+5%"} /></div>
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
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[25, 24, 32, 36, 32, 30, 33, 33, 20, 17, 19, 34]} data2={[10, 15, 8, 20, 17, 15, 13, 17, 16, 14, 5, 27]} data3={[15,9,24,16,15,15,20,16,4,3,14,7]} min1={"12"} min2={12} min3={12} /></div>
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[12, 17, 19, 22, 24, 20, 18, 26, 20, 17, 15, 19]} data2={[5,12,14,4,12,6,7,12,16,3,5,7]} data3={[7,5,5,18,12,14,11,14,4,14,10,12]} min1={"12"} min2={12} min3={12} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[23, 38, 30, 17, 26, 18, 34, 13, 19, 39, 22, 14]} data2={[17,18,20,7,16,11,19,10,12,30,14,10]} data3={[6,20,10,10,10,7,15,3,7,9,8,4]} min1={"12"} min2={12} min3={12} /></div>
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[17, 18, 17, 13, 40, 17, 36, 33, 25, 34, 36, 15]} data2={[14,13,13,7,29,10,20,14,16,20,17,10]} data3={[3,5,4,6,11,7,16,19,9,14,19,5]} min1={"12"} min2={12} min3={12} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[40, 24, 37, 39, 21, 14, 19, 36, 27, 31, 28, 14]} data2={[30,10,14,19,14,10,6,10,17,20,21,10]} data3={[10,14,23,20,7,4,13,26,10,11,7,4]} min1={"12"} min2={12} min3={12} /></div>
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[21, 18, 17, 35, 38, 16, 40, 18, 12, 24, 30, 20]} data2={[14,13,13,21,22,10,23,10,4,19,18,13]} data3={[7,5,4,14,16,6,17,8,8,5,12,7]} min1={"12"} min2={12} min3={12} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[16, 24, 23, 36, 19, 20, 25, 29, 29, 22, 34, 37]} data2={[10,18,14,30,12,13,20,18,17,13,30,30]} data3={[6,6,9,6,7,7,5,11,12,9,4,7]} min1={"12"} min2={12} min3={12} /></div>
                    <div className="dashboardColumn"><MultipleLineChart title={"All Contacts"} yaxis={"Contacts"} data1={[28, 30, 17, 18, 36, 13, 23, 36, 34, 23, 15, 26]} data2={[21,22,14,13,15,7,16,15,17,18,14,13]} data3={[7,8,3,5,21,6,7,21,17,5,1,13]} min1={"12"} min2={12} min3={12} /></div>
                </div>
                <div className="dashboardRow">
                    <div className="tableSet">
                        <div className="dashboardColumn" style={{width:"55%"}}><EnhancedTable2/></div>
                    </div>
                </div>
            </div>

            <div className="navbarPurple"></div>
            <div className="lineCardGroupSet">
                <div className="lineCardGroup1">
                    <LineChartCard />
                    <LineChartCard />
                    <LineChartCard />
                    <LineChartCard />
                </div>
                <div className="lineCardGroup2">

                </div>
            </div>

        </div>

    )
}