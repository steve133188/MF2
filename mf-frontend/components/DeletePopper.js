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
                    <span onClick={handleClick}></span>
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
