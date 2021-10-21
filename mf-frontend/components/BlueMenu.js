import Link from "next/link"

export function BlueMenu(props) {
    return(

            <div className="blueMenu">
                <ul className="blueMenuGroup">
                    {props.children}
                </ul>
            </div>

    )
}