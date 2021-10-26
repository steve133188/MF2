import Button from "@mui/material/Button";
import {createTheme, ThemeProvider} from '@mui/material/styles';

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

export function IconButton() {
    return (
        <div className="iconButtonContainer">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="neutral">
                    <img
                        src="https://uxwing.com/wp-content/themes/uxwing/download/17-internet-network-technology/robot-line.png"
                        width="17px" height="17px" alt=""/>
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function NormalButton(props) {
    return (
        <div className="normalButton">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="neutral">
                    {props.children}
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function ToggleButton(props) {
     return(
         <div className={props.classname}>
             <ThemeProvider theme={theme}>
                 <Button variant="contained" color="neutral">
                     {props.children}
                 </Button>
             </ThemeProvider>
         </div>
     )
}

export function SelectButton() {
    return (
        <div className="normalButton">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="neutral">
                    Select
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function NormalButton2({ children ,...props}) {
    const {disabled , onClick} = props
    return (
        <div className="newContactButton">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="primary"  disabled={disabled} onClick={onClick}>
                    {children}
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function CancelButton() {
    return (
        <div className="cancelButton">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="cancel">
                    Cancel
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function NormalButton3(props) {
    return (
        <div className="newNoteButton">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="secondary" onClick={onClick}>
                    {props.children}
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function TextWithIconButton({ children, ...props }) {
    const {onClick} = props
    return (
        <div className="textWithIconButton">
            <ThemeProvider theme={theme}>
                <Button variant="contained" color="neutral" onClick={onClick}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                         className="bi bi-pencil" viewBox="0 0 16 16" style={{marginRight: "4px"}}>
                        <path
                            d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5 13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"/>
                    </svg>
                    <span>{children}</span>
                </Button>
            </ThemeProvider>
        </div>
    )
}

export function ThreeDotsMenu() {
    return (
        <div className="threeDotsMenu"></div>
    )
}