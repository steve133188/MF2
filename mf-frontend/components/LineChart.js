import {Line} from 'react-chartjs-2'

export function LineChart({...props}){
    // const data = (canvas) => {
        // const ctx = canvas.getContext('2d')
        // const gradient = ctx.createLinearGradient(0,0,100,0);
        //
        // return {
        //     backgroundColor: gradient
        //     // ...the rest
        // }
    // }


    return <Line data={{
        labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
    }}
    height={400} width={600}
    />;
}