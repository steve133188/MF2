import * as React from 'react';
import {LabelSelect, MultipleSelectPlaceholder} from "./Select";
import Box from '@mui/material/Box';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import {CancelButton, NormalButton2} from "./Button";
import {DeletePopper} from "./DeletePopper";
import {AddPopper} from "./AddPopper";


export function NavbarPurple() {
    const [open, setOpen] = React.useState(false);

    const handleClick = () => {
        setOpen((prev) => !prev);
    };

    const handleClickAway = () => {
        setOpen(false);
    };
    return (
        <div className="navbarPurple">
            <div className="selectButtonGroup">
                <MultipleSelectPlaceholder/>
                <MultipleSelectPlaceholder/>
                <MultipleSelectPlaceholder/>
                <MultipleSelectPlaceholder/>
            </div>

            <div className="tagsButtonGroup">
                <AddPopper/>

                <span>2</span>

                <DeletePopper/>
            </div>
        </div>
    )
}