import * as React from 'react';
import {LabelSelect, MultipleSelectPlaceholder} from "./Select";
import Box from '@mui/material/Box';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import {CancelButton, NormalButton2} from "./Button";
import {CheckboxGroup2, Search2, Search3} from "./Input";
import {Checkbox2, CheckboxPill} from "./Checkbox";


export function AddPopper() {
    const [openAddPopper, setOpenAddPopper] = React.useState(false);

    const handleClickAddPopper = () => {
        setOpenAddPopper((prev) => !prev);
    };

    const handleClickAwayAddPopper = () => {
        setOpenAddPopper(false);
    };
    const addStyles = {
        position: 'absolute',
        top: 50,
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
            <ClickAwayListener onClickAway={handleClickAwayAddPopper}>
                <Box sx={{position: 'relative'}}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor"
                         className="bi bi-tag" viewBox="0 0 16 16" onClick={handleClickAddPopper}>
                        <path d="M6 4.5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm-1 0a.5.5 0 1 0-1 0 .5.5 0 0 0 1 0z"/>
                        <path
                            d="M2 1h4.586a1 1 0 0 1 .707.293l7 7a1 1 0 0 1 0 1.414l-4.586 4.586a1 1 0 0 1-1.414 0l-7-7A1 1 0 0 1 1 6.586V2a1 1 0 0 1 1-1zm0 5.586 7 7L13.586 9l-7-7H2v4.586z"/>
                    </svg>
                    {openAddPopper ? (
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