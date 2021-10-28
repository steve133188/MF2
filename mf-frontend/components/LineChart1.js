import {useState} from "react";
import dynamic from "next/dynamic";
const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

export function LineChart1({children,...props}) {
    const {title, data, yaxis, total, percentage} = props;
    const [state, setState] = useState({
        series: [{
            name: "Desktops",
            data: data,
            yaxis: "contacts"
        }],
        options: {
            chart: {
                id: 'fb',
                group: 'social',
                type: 'line',
                height: "35%"
            },
           //  subtitle: {
           //      text: ['Total', "32 +5%"],
           //      align: 'left',
           //
           // },
            colors: ['#5B73E8'],
            stroke: {
                curve: 'straight'
            },
            // title: {
            //     text: title,
            //     align: 'left'
            // },
            markers: {
                size: 6
            },
            xaxis: {
                categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
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
                    text: yaxis
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
                    <div><p style={{color: "#495057", fontSize: "16px", fontWeight: "600"}}>{title}</p>
                        <div style={{marginLeft: "25px", fontSize: "12px", color: "#74788D"}}>Total <div style={{color: "#6279EC", fontSize: "20px", fontWeight: "bold"}}>{total}<span style={{marginLeft: "6px", fontSize: "8px", fontWeight: "600", color: "#34C38F"}}>{percentage}</span></div></div>
                    </div>
                    <Chart options={state.options} series={state.series} type="line" />
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
            // subtitle: {
            //     text: 'Longest',
            //     align: 'left',
            // },
            colors: ['#5B73E8', '#68C093', '#F1B44C'],
            chart: {
                id: 'fb',
                group: 'social',
                type: 'line',
                height: "35%"
            },
            stroke: {
                curve: 'straight'
            },
            title: {
                text: '.',
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
                offsetY: -10,
                offsetX: -5
            }
        },

    })


    return (
        <div>
            <div id="wrapper">
                <div id="chart-line">
                    <div><p style={{color: "#495057", fontSize: "16px", fontWeight: "600"}}>Title</p>
                        <div style={{marginLeft: "25px", fontSize: "12px", color: "#74788D"}}>Longest <div style={{color: "#6279EC", fontSize: "20px", fontWeight: "bold"}}>12 mins<span style={{color: "#34C38F", fontSize: "20px", fontWeight: "bold", marginLeft: "20px"}}>12 mins</span><span style={{color: "#F1B44C", fontSize: "20px", fontWeight: "bold", marginLeft: "20px"}}>12 mins</span></div></div>
                    </div>
                    <Chart options={state.options} series={state.series} type="line"/>
                </div>
            </div>
        </div>
    )
}
