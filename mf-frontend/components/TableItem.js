export function TableItem(props) {
    return(
        <td className={props.classname}>{props.children}</td>
    )
}