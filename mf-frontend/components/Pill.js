export function Pill({children,...props}) {
    const {color, size} = props;
    const name = "pill " + color + " " + size;
    return (
        <span className="pillContainer">
            <span className={name}>{children}</span>
        </span>
    )
}

export function StatusPill({children,...props}) {
    const {color} = props;
    const name = "pill " + color;
    return (
        <span className="pillContainer">
            <span className={name}>
                <li>{children}</li>
            </span>
        </span>
    )
}

export function PillInIcon({children, ...props}) {
    return (
        <span className="pillInIcon">{children}</span>

    )
}


