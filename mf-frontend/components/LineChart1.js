import {useState} from "react";
import dynamic from "next/dynamic";
const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

export function LineChart1() {

    const [state, setState] = useState({
        series: [{
            name: "Desktops",
            data: [25, 24, 32, 36, 32, 30, 33]
        }],
        options: {
            chart: {
                id: 'fb',
                group: 'social',
                type: 'line',
                height: 160
            },
            colors: ['#5B73E8'],
            stroke: {
                curve: 'straight'
            },
            title: {
                text: 'All Contacts',
                align: 'left'
            },
            markers: {
                size: 6
            },
            xaxis: {
                categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
                title: {
                    text: 'Month'
                },
                label: {
                    style: {
                        color: "#8b8b8b",
                        fontSize: "10"
                    }
                }
            },
            yaxis: {
                title: {
                    text: 'Contacts'
                },
                label: {
                    style: {
                        fontSize: "10",
                        color: "#8B8B8B"
                    }
                }
            },
        },

    })


    return (
        <div>
            <div id="wrapper">
                <div id="chart-line">
                    <Chart options={state.options} series={state.series} type="line" height={560} width={900}/>
                </div>
            </div>
        </div>
    )
}

export function MultipleLineChart() {
    const [state, setState] = useState({
        series: [{
            name: "Total Contacts",
            data: [10, 41, 35, 51, 49, 62, 69, 91, 148]
        },
            {
                name: "Mary Foster",
                data: [10, 21, 45, 61, 59, 42, 39, 81, 128]
            },
            {
                name: "Harry Stewart",
                data: [3, 21, 35, 21, 19, 32, 49, 91, 38]
            }
        ],
        options: {
            subtitle: {
                text: 'on9',
                align: 'left',

            },
            colors: ['#5B73E8', '#68C093', '#F1B44C'],
            chart: {
                id: 'fb',
                group: 'social',
                type: 'line',
                height: 160
            },
            stroke: {
                curve: 'straight'
            },
            title: {
                text: 'All Contacts',
                align: 'left'
            },
            markers: {
                size: 6,
            },
            xaxis: {
                categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
                title: {
                    text: 'Month'
                }
            },
            yaxis: {
                title: {
                    text: 'Contacts'
                },
            },
            legend: {
                position: 'top',
                horizontalAlign: 'right',
                floating: true,
                offsetY: -25,
                offsetX: -5
            }
        },

    })


    return (
        <div>
            <div id="wrapper">
                <div id="chart-line">
                    <Chart options={state.options} series={state.series} type="line" height={560} width={900}/>
                </div>
            </div>
        </div>
    )
}

// import {useState} from "react";
// import dynamic from "next/dynamic";
// // import ReactApexChart from 'react-apexcharts'
// const Chart = dynamic(() => import('react-apexcharts'), { ssr: false });
//
// export function LineChart1() {
//
//     const [state, setState] = useState({
//         series: [{
//             name: "Desktops",
//             data: [25, 24, 32, 36, 32, 30, 33]
//         }],
//         options: {
//             chart: {
//                 id: 'fb',
//                 group: 'social',
//                 type: 'line',
//                 height: 160
//             },
//             colors: ['#5B73E8'],
//             stroke: {
//                 curve: 'straight'
//             },
//             title: {
//                 text: 'All Contacts',
//                 align: 'left'
//             },
//             markers: {
//                 size: 6
//             },
//             xaxis: {
//                 categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
//                 title: {
//                     text: 'Month'
//                 },
//                 label: {
//                     style: {
//                         color: "#8b8b8b",
//                         fontSize: "10"
//                     }
//                 }
//             },
//             yaxis: {
//                 title: {
//                     text: 'Contacts'
//                 },
//                 label: {
//                     style: {
//                         fontSize: "10",
//                         color: "#8B8B8B"
//                     }
//                 }
//             },
//         },
//
//     })
//
//
//     return (
//         <div>
//             <div id="wrapper">
//                 <div id="chart-line">
//                     <Chart options={state.options} series={state.series} type="line" height={560} width={900}/>
//                 </div>
//             </div>
//         </div>
//     )
// }
//
// export function MultipleLineChart() {
//     const [state, setState] = useState({
//         series: [{
//             name: "Total Contacts",
//             data: [10, 41, 35, 51, 49, 62, 69, 91, 148]
//         },
//             {
//                 name: "Mary Foster",
//                 data: [10, 21, 45, 61, 59, 42, 39, 81, 128]
//             },
//             {
//                 name: "Harry Stewart",
//                 data: [3, 21, 35, 21, 19, 32, 49, 91, 38]
//             }
//         ],
//         options: {
//             subtitle: {
//                 text: 'on9',
//                 align: 'left',
//
//             },
//             colors: ['#5B73E8', '#68C093', '#F1B44C'],
//             chart: {
//                 id: 'fb',
//                 group: 'social',
//                 type: 'line',
//                 height: 160
//             },
//             stroke: {
//                 curve: 'straight'
//             },
//             title: {
//                 text: 'All Contacts',
//                 align: 'left'
//             },
//             markers: {
//                 size: 6,
//             },
//             xaxis: {
//                 categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
//                 title: {
//                     text: 'Month'
//                 }
//             },
//             yaxis: {
//                 title: {
//                     text: 'Contacts'
//                 },
//             },
//             legend: {
//                 position: 'top',
//                 horizontalAlign: 'right',
//                 floating: true,
//                 offsetY: -25,
//                 offsetX: -5
//             }
//         },
//
//     })
//
//
//     return (
//         <div>
//             <div id="wrapper">
//                 <div id="chart-line">
//                     <Chart options={state.options} series={state.series} type="line" height={560} width={900}/>
//                 </div>
//             </div>
//         </div>
//     )
// }

