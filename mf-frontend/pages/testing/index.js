import {BlueMenu, BlueMenuGroup} from "../../components/BlueMenu";
import {BlueMenuDropdown, BlueMenuLink} from "../../components/BlueMenuLink";
<<<<<<< HEAD
import dynamic from "next/dynamic";
import {ChangingPercentageCard} from "../../components/Cards";

// const LineChart1 = dynamic(
//     () => import('../../components/LineChart1').then(mod => mod.LineChart1),
=======
// import dynamic from "next/dynamic";
// import {ChangingPercentageCard} from "../../components/Cards";
// import {useEffect} from "react";
//
// const LineChart1 = dynamic(
//     async () => {
//         return await  import('../../components/LineChart1').then(mod => mod.LineChart1)
//     },
>>>>>>> 4fc5d1a0ba164bb3d9a2bcd60b1ab3307458e16a
//     {ssr: false}
// );
//
// const MultipleLineChart = dynamic(
<<<<<<< HEAD
//     () => import('../../components/LineChart1').then(mod => mod.MultipleLineChart),
=======
//     async () => {
//         return await  import('../../components/LineChart1').then(mod => mod.MultipleLineChart)
//     },
>>>>>>> 4fc5d1a0ba164bb3d9a2bcd60b1ab3307458e16a
//     {ssr: false}
// );
//
// const LineChartCard = dynamic(
<<<<<<< HEAD
//     () => import('../../components/Cards').then(mod => mod.LineChartCard),
//     {ssr: false}
// );
import {LineChartCard} from "../../components/Cards";
import {LineChart1, MultipleLineChart} from "../../components/LineChart1";

=======
//     async () => {
//         return await import('../../components/Cards').then(mod => mod.LineChartCard)
//     },
//     {ssr: false}
// );
//
//
>>>>>>> 4fc5d1a0ba164bb3d9a2bcd60b1ab3307458e16a
export default function Testing() {
    // useEffect(()=>{
    //
    // },[])

    return (
        <>
            {/*<LineChart1/>*/}
            {/*<MultipleLineChart/>*/}
            {/*<LineChartCard/>*/}
            {/*<ChangingPercentageCard />*/}
        </>
    )
}