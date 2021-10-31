import {IconButton} from "./Button";

export function FileLink({children,...props}) {
    const {name, icon, date, size} = props;
    return(
        <div className="fileLink">
            <IconButton/>
            <div className="fileLinkInfo">
                <span className="fileLinkName">Attachment.pdf</span>
                <div className=""><span className="fileDate">3 June, 2021</span><span className="fileSize">224KB</span></div>
            </div>
        </div>
    )
}