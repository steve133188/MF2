export function Pill(props) {
    const name = "pill " + props.color + " " + props.size;
    return (
        <span className="pillContainer">
            <span className={name}>{props.children}</span>
        </span>
    )
}

export function StatusPill(props) {
    const name = "pill " + props.color;
    return (
        <span className="pillContainer">
            <span className={name}>
                <li>{props.children}</li>
            </span>
        </span>
    )
}

export function PillInIcon(props) {
    return (
        <span className="pillInIcon">{props.children}</span>

    )
}