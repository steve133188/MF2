// import {BlueMenu, BlueMenuGroup} from "../../components/BlueMenu";
// import {BlueMenuDropdown, BlueMenuLink} from "../../components/BlueMenuLink";

// import dynamic from "next/dynamic";
import {ChangingPercentageCard} from "../../components/Cards";
import {LineChartCard} from "../../components/Cards";
import {LineChart1, MultipleLineChart} from "../../components/LineChart1";


export default function Testing() {

    return (
        <>
            <LineChart1/>
            <MultipleLineChart/>
            <LineChartCard/>
            <ChangingPercentageCard />
        </>
    )
}