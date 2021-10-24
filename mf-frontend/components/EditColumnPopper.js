import * as React from 'react';
import Box from '@mui/material/Box';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import {NormalButton, TextWithIconButton} from "./Button";
import {ColumnCheckbox} from "./Checkbox";

export function EditColumnPopper() {

    const editColumnStyles = {
        position: 'absolute',
        top: 55,
        right: -10,
        borderRadius: 2,
        zIndex: 2,
        boxShadow: 3,
        p: 1,
        bgcolor: 'background.paper',
        textAlign: "center",
        padding: "32px 29px 48px 29px",
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
        <div className="editColumnPopperContainer">
            <ClickAwayListener onClickAway={handleClickAway}>
                <Box sx={{position: 'relative'}}>
                    <TextWithIconButton onClick={handleClick}>Edit Column</TextWithIconButton>
                    {open ? (
                        <Box sx={editColumnStyles}>
                            <div className="topSide">
                                <span>Column Setting</span>
                                <NormalButton>Add</NormalButton>
                            </div>

                            <div className="columnGroup">
                                <ColumnCheckbox>Customer ID</ColumnCheckbox>
                                <ColumnCheckbox>Name</ColumnCheckbox>
                                <ColumnCheckbox>Team</ColumnCheckbox>
                                <ColumnCheckbox>Channel</ColumnCheckbox>
                                <ColumnCheckbox>Tag</ColumnCheckbox>
                                <ColumnCheckbox>Assignee</ColumnCheckbox>
                            </div>

                        </Box>
                    ) : null}
                </Box>
            </ClickAwayListener>
        </div>
    )
}
