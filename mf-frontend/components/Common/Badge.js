export function Badge(props) {
    const name = "badge bg-" + props.color;
    return (
        <span className={name}>{props.children}</span>
    )
}

