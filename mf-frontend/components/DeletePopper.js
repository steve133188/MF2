import * as React from 'react';
import {LabelSelect, MultipleSelectPlaceholder} from "./Select";
import Box from '@mui/material/Box';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import {CancelButton, NormalButton2} from "./Button";


export function DeletePopper() {

    const deleteStyles = {
        position: 'absolute',
        top: 55,
        right: -10,
        borderRadius: 2,
        zIndex: 1,
        boxShadow: 3,
        p: 1,
        bgcolor: 'background.paper',
        textAlign: "center",
        padding: "41px 32px 28px 32px",
        width: 376,
        lineHeight: 2
    };
    const [open, setOpen] = React.useState(false);

    const handleClick = () => {
        setOpen((prev) => !prev);
    };

    const handleClickAway = () => {
        setOpen(false);
    };

    return (
        <div className="deletePopperContainer">
            <ClickAwayListener onClickAway={handleClickAway}>
                <Box sx={{position: 'relative'}}>

                    <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="#f46a6b" cursor="pointer"
                         className="bi bi-trash" viewBox="0 0 16 16" onClick={handleClick}>
                        <path
                            d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                        <path fillRule="evenodd"
                              d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                    </svg>
                    {open ? (
                        <Box sx={deleteStyles}>
                            Delete 2 contacts? <br/>
                            All conversation history will also be erased.
                            <div className="controlButtonGroup">
                                <NormalButton2>Delete</NormalButton2>
                                <span onClick={handleClick}><CancelButton></CancelButton></span>
                            </div>
                        </Box>
                    ) : null}
                </Box>
            </ClickAwayListener>
        </div>
    )
}
