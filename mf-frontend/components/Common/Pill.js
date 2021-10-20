export function Pill(props) {
    const name = "badge rounded-pill bg-" + props.color;
    return (
        <span className={name}>{props.children}</span>
    )
}