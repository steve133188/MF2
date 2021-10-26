export function Badge(props) {
    const name = "badge " + props.color;
    return (
        <span className="badgeContainer">
            <span className={name}>{props.children}</span>
        </span>
    )
}

export function MoreImageBadge(props) {
    return (
        <span className="moreImageBadgeContainer">
            <span className="moreImageBadge">
                +99
            </span>
        </span>
    )
}