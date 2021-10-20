import * as React from 'react';
import {LabelSelect, MultipleSelectPlaceholder} from "./Select";
import Box from '@mui/material/Box';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import {CancelButton, NormalButton2} from "./Button";
import {CheckboxGroup2, Search2, Search3} from "./Input";
import {Checkbox2, CheckboxPill} from "./Checkbox";

export function AddPopper() {
    const [open, setOpen] = React.useState(false);

    const handleClick = () => {
        setOpen((prev) => !prev);
    };

    const handleClickAway = () => {
        setOpen(false);
    };
    const addStyles = {
        position: 'absolute',
        top: 25,
        right: -30,
        borderRadius: 2,
        zIndex: 1,
        boxShadow: 3,
        p: 1,
        bgcolor: 'background.paper',
        textAlign: "center",
        padding: "41px 32px 28px 32px",
        width: 457,
        lineHeight: 2
    };

    return (
        <div className="addPopperContainer">
            <ClickAwayListener onClickAway={handleClickAway}>
                <Box sx={{position: 'relative'}}>
                    <span onClick={handleClick}>1</span>
                    {open ? (
                        <Box sx={addStyles}>
                            <div className="addTagHeader">
                                <span>Add Tag</span>
                                <NormalButton2>Confirm</NormalButton2>
                            </div>
                            <Search3 type="search">Search</Search3>
                            <CheckboxGroup2>
                                <CheckboxPill/>
                                <CheckboxPill/>
                            </CheckboxGroup2>
                        </Box>
                    ) : null}
                </Box>
            </ClickAwayListener>
        </div>
    )
}