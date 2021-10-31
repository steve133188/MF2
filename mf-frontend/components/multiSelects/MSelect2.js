import Select from "@mui/material/Select";
import OutlinedInput from "@mui/material/OutlinedInput";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import * as React from "react";
import {useTheme} from "@mui/material/styles";

export function MSelect2() {
    const theme = useTheme();
    const [personName, setPersonName] = React.useState([]);

    const handleChange = (event) => {
        const {
            target: {value},
        } = event;
        setPersonName(
            // On autofill we get a the stringified value.
            typeof value === 'string' ? value.split(',') : value,
        );
    };
    const ITEM_HEIGHT = 48;
    const ITEM_PADDING_TOP = 8;
    const MenuProps = {
        PaperProps: {
            style: {
                maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
                width: 250,
            },
        },
    };
    return(
        <FormControl sx={{m: 0, width: 171, mt: 1}}>

            <Select sx={{
                height: 28,
                marginBottom: 0.3,
                marginRight: 3,
                borderRadius: 2,
                background: "white"
            }}
                    multiple
                    displayEmpty
                    value={personName}
                    onChange={handleChange}
                    input={<OutlinedInput/>}
                    renderValue={(selected) => {
                        if (selected.length === 0) {
                            return <span>Team</span>;
                        }
                        return selected.join('');
                    }}
                    MenuProps={MenuProps}
                    inputProps={{'aria-label': 'Without label'}}
            >
                <MenuItem disabled value="">
                    <span>Team</span>
                </MenuItem>

                <MenuItem
                    value={"Team A"}
                >
                    {"Team A"}
                </MenuItem>
                <MenuItem
                    value={"Team B"}
                >
                    {"Team B"}
                </MenuItem>
                <MenuItem
                    value={"Team C"}
                >
                    {"Team C"}
                </MenuItem>
                <MenuItem
                    value={"Team D"}
                >
                    {"Team D"}
                </MenuItem>
            </Select>

        </FormControl>
    )
}