import {LineChart} from "../../components/LineChart";

export default function Dashboard() {
    return(
        <>
            <div id={"dashboard"}>
                <LineChart />
            </div>
            <div className="dashboardContainer">
                <img src="./001.png" alt=""/>
                <img src="./002.png" alt=""/>
                <img src="./003.png" alt=""/>
                <img src="./dashBackground.png" alt=""/>
            </div>
        </>

    )
}