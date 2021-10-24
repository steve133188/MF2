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
                    <img
                        src="https://uxwing.com/wp-content/themes/uxwing/download/17-internet-network-technology/robot-line.png" alt=""/>
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