export function Alert() {
    return (
        <div className={"alertContainer"}>
            <img src="iconWhatsappmed.svg" className={"msgTypeIcon"} width={"30px"} height={"30px"} alt=""/>
            <div className={"msgGrp"}>
                <div className={"name"} style={{fontSize:"16px", fontWeight: "bold", color: "#444444"}}><img className={"senderIcon"} src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" width={"25px"} height={"25px"} alt="" style={{borderRadius: "50%", marginRight: "7px"}} />Debra Patel</div>
                <div className={"message"} style={{fontSize: "13px", color: "#444444"}}>Lorem ipsum dolor sit amet</div>
            </div>
        </div>
    )
}