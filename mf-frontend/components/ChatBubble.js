export function ChatBubble(props){
    const chatMessageType = props.type + "ChatMessages"
    const MessageTime = props.messageTime

    return(
        <div className={chatMessageType}>
            <div className="chatBubble">
                {props.children}
            </div>
            <div className="messageTime">{MessageTime}</div>
        </div>
    )
}