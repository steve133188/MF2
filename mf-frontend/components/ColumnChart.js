import {useState} from "react";
import dynamic from "next/dynamic";
const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

export function ColumnChart() {

    const [state, setState] = useState({
        series: [{
            name: 'Unhandled Contacts',
            data: [44, 55, 41, 67, 22, 43, 14, 16, 28, 60]
        }, {
            name: 'Delivered Contacts',
            data: [13, 23, 20, 8, 13, 27, 10, 46, 50, 40]
        }, {
            name: 'Active Contacts',
            data: [11, 17, 15, 15, 21, 14, 13, 40, 67, 10]
        }],
        options: {
            colors: ['#2198FA', '#6279EC', '#34C38F'],
            chart: {
                type: 'bar',
                height: 350,
                stacked: true,
                toolbar: {
                    show: true
                },
                zoom: {
                    enabled: true
                },

            },

            plotOptions: {
                bar: {
                    horizontal: false,

                },
            },
            xaxis: {
                type: 'category',
                categories: ['Harry Stewart', 'Jasmine Miller', 'Doris Mendoza', 'Mike Mendoza', 'Eugene Jackson', 'Bianca Hayes', 'Elizabeth Coleman', 'Will Fisher', 'Timothy Powell', 'Emma Marshall'
                ],
            },
            legend: {
                position: 'bottom',
                offsetX: -10,
                offsetY: 10
            },
            fill: {
                opacity: 1
            }
        },
    })

    return (

        <div id="chart">
            <Chart options={state.options} series={state.series} type="bar" height={350} />
        </div>
    )
}