import {TextRadio2} from "./Input";
import {ContactFileImage} from "./Image"
import {MoreImageBadge} from "./Badge";
import {FileLink} from "./FileLink";
import {NormalButton} from "./Button";
import {ContactBasicInfo} from "./ContactBasicInfo";

export function ContactFile() {
    return (
        <div className="container">
            <div className="contactInfoSet">
                <ContactBasicInfo icon="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg" name="Debra Patel" phone="+852 97650348" contactType="https://www.pngrepo.com/png/158412/512/whatsapp.png" pillColor="teamA" pillContent="Team A" />
                <TextRadio2/>
                <div className="contactFileContainer">
                    <div className="mediaFileLinkContainer">
                        <div className="mediaContainer">
                            <div className="fileSubHeader">
                                <div>Media</div>
                                <span className="fileAmount">4 Photo, 1 Video</span>
                            </div>
                            <div className="mediaGroup">
                                <ContactFileImage
                                    src="https://storage.googleapis.com/yk-cdn/photos/plp/mina-mimbu/rainbow.jpg"/>
                                <ContactFileImage
                                    src="https://storage.googleapis.com/yk-cdn/photos/plp/mina-mimbu/rainbow.jpg"/>
                                <MoreImageBadge/>
                            </div>
                        </div>
                        <div className="fileLinkContainer">
                            <div className="fileSubHeader">
                                <div>Files & Links</div>
                                <span className="fileAmount">2 File, 1 Link</span>
                            </div>
                            <div className="fileLinkGroup">
                                <FileLink />
                                <FileLink />
                                <FileLink />
                            </div>
                            <NormalButton>
                                View All Files & Links
                            </NormalButton>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}