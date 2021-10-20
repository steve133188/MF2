import {Pill} from "./Pill"

export function Note(props) {
    return (
        <div className="noteContainer">
            <Pill color={props.pillColor} size={props.pillSize} >{props.pillContent}</Pill>
            <div className="noteDetails">
                <div className="nameAndTime">
                    <span className="name">{props.name}</span>
                    <span className="time">{props.time}</span>
                </div>
                <div className="noteContent">{props.children}</div>
            </div>
        </div>
    )
}