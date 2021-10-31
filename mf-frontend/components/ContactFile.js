import {TextRadio2} from "./Input";
import {ContactFileImage} from "./Image"
import {MoreImageBadge} from "./Badge";
import {FileLink} from "./FileLink";
import {IconButton, NormalButton} from "./Button";
import {ContactBasicInfo} from "./ContactBasicInfo";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import Button from "@mui/material/Button";
const theme = createTheme({
    palette: {
        neutral: {
            main: '#DEF0FF',
            contrastText: '#2198FA',
        },
        primary: {
            main: '#2198FA',
            contrastText: 'white',
        },
        secondary: {
            main: '#F1B44C',
            contrastText: 'white',
        },
        cancel: {
            main: '#F5F6F8',
            contrastText: '#444444',
        }
    },
});
export function ContactFile() {
    return (
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
                                <span className="fileAmount">2 File, 0 Link</span>
                            </div>
                            <div className="fileLinkGroup">
                                <div className="fileLink">
                                    <div className="iconButtonContainer">
                                        <ThemeProvider theme={theme}>
                                            <Button variant="contained" color="neutral">
                                                <img
                                                    src="fileLogo.svg"
                                                    width="17px" height="17px" alt=""/>
                                            </Button>
                                        </ThemeProvider>
                                    </div>
                                    <div className="fileLinkInfo">
                                        <span className="fileLinkName">Attachment.pdf</span>
                                        <div className=""><span className="fileDate">3 June, 2021</span><span className="fileSize">224KB</span></div>
                                    </div>
                                </div>

                                <div className="fileLink">
                                    <div className="iconButtonContainer">
                                        <ThemeProvider theme={theme}>
                                            <Button variant="contained" color="neutral">
                                                <img
                                                    src="/fileLogo.svg"
                                                    width="17px" height="17px" alt=""/>
                                            </Button>
                                        </ThemeProvider>
                                    </div>
                                    <div className="fileLinkInfo">
                                        <span className="fileLinkName">Attachment2.pdf</span>
                                        <div className=""><span className="fileDate">3 June, 2021</span><span className="fileSize">200KB</span></div>
                                    </div>
                                </div>
                            </div>
                            <NormalButton>
                                View All Files & Links
                            </NormalButton>
                        </div>
                    </div>
                </div>
            </div>
    )
}