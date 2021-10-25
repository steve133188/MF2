// import {BlueMenu, BlueMenuGroup} from "../../components/BlueMenu";
// import {BlueMenuDropdown, BlueMenuLink} from "../../components/BlueMenuLink";

// import dynamic from "next/dynamic";
import {ChangingPercentageCard} from "../../components/Cards";

// const LineChart1 = dynamic(
//     () => import('../../components/LineChart1').then(mod => mod.LineChart1),
// import dynamic from "next/dynamic";
// import {ChangingPercentageCard} from "../../components/Cards";
// import {useEffect} from "react";
//
// const LineChart1 = dynamic(
//     async () => {
//         return await  import('../../components/LineChart1').then(mod => mod.LineChart1)
//     },
//     {ssr: false}
// );
//
// const MultipleLineChart = dynamic(\
//     () => import('../../components/LineChart1').then(mod => mod.MultipleLineChart),
//     async () => {
//         return await  import('../../components/LineChart1').then(mod => mod.MultipleLineChart)
//     },
//     {ssr: false}
// );
//
// const LineChartCard = dynamic(
//     () => import('../../components/Cards').then(mod => mod.LineChartCard),
//     {ssr: false}
// );
import {LineChartCard} from "../../components/Cards";
import {LineChart1, MultipleLineChart} from "../../components/LineChart1";

//     async () => {
//         return await import('../../components/Cards').then(mod => mod.LineChartCard)
//     },
//     {ssr: false}
// );
//
//

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