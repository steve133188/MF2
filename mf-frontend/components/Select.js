import * as React from 'react';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import FormControl from '@mui/material/FormControl';
import {Pill} from "./Pill";

export function LabelSelect() {
    const [age, setAge] = React.useState('');

    const handleChange = (event) => {
        setAge(event.target.value);
    };
    return (
        <div className="selectContainer">
            <Select
                value={age}
                onChange={handleChange}
                displayEmpty

            >
                <MenuItem value="">
                    <span>None</span>
                </MenuItem>
                <MenuItem value={10}>Ten</MenuItem>
                <MenuItem value={20}>Twenty</MenuItem>
                <MenuItem value={30}>Thirty</MenuItem>
            </Select>
        </div>
    )
}


export function TeamFilterSelect({children, ...props}) {
    const {link, title, onClick, isSelect, setSelect} = props
    const [age, setAge] = React.useState('');

    const handleChange = (event) => {
        setAge(event.target.value);
        console.log(event.target.value);
    };

    return (
        <FormControl sx={{m: 1, minWidth: 120}}>
            <Select
                value={age}
                onChange={handleChange}
                displayEmpty
                inputProps={{'aria-label': 'Without label'}}
                sx={{
                    width: 160,
                    height: 31,
                    borderRadius: 19,
                    background: "#D0E9FF",
                    border: "none",
                    textAlign: "center",
                }}
            >
                <MenuItem value="">

                    <span>Mary Foster</span>
                    <div className={"smallPill"}>20</div>
                </MenuItem>
                <MenuItem value={10}>

                    <span>Ten</span>
                    <div className={"smallPill"}>20</div>
                </MenuItem>
                <MenuItem value={20}>

                    <span>Twenty</span>
                    <div className={"smallPill"}>20</div>
                </MenuItem>
                <MenuItem value={30}>

                    <span>Thirty</span>
                    <div className={"smallPill"}>20</div>
                </MenuItem>
            </Select>
        </FormControl>
    )
}


export function LabelSelect2() {
    const [age, setAge] = React.useState('');

    const handleChange = (event) => {
        setAge(event.target.value);
    };
    return (
        <div className="select2Container">
            <Select
                value={age}
                onChange={handleChange}
                displayEmpty

            >
                <MenuItem value="">
                    <span>None</span>
                </MenuItem>
                <MenuItem value={10}>Ten</MenuItem>
                <MenuItem value={20}>Twenty</MenuItem>
                <MenuItem value={30}>Thirty</MenuItem>
            </Select>
        </div>
    )
}

import {useTheme} from '@mui/material/styles';
import OutlinedInput from '@mui/material/OutlinedInput';


export function MultipleSelectPlaceholder({ children,...props }) {
    const {placeholder, selectItems} = props;
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

    const names = [
        'Oliver Hansen',
        'Van Henry',
        'April Tucker',
        'Ralph Hubbard',
        'Omar Alexander',
        'Carlos Abbott',
        'Miriam Wagner',
        'Bradley Wilkerson',
        'Virginia Andrews',
        'Kelly Snyder',
    ];

    function getStyles(name, personName, theme) {
        return {
            fontWeight:
                personName.indexOf(name) === -1
                    ? theme.typography.fontWeightRegular
                    : theme.typography.fontWeightMedium,
        };
    }

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

    return (
        <div className="multipleSelectPlaceholder">
            <FormControl sx={{m: 0, width: 171, mt: 1}}>
                <Select sx={{height: 28, marginBottom: 0.3, marginRight: 3, borderRadius: 2, background: "white"}}
                        multiple
                        displayEmpty
                        value={personName}
                        onChange={handleChange}
                        input={<OutlinedInput/>}
                        renderValue={(selected) => {
                            if (selected.length === 0) {
                                return <span>Agent</span>;
                            }
                            return selected.join('');
                        }}
                        MenuProps={MenuProps}
                        inputProps={{'aria-label': 'Without label'}}
                >
                    <MenuItem disabled value="">
                        <span>Agnet</span>
                    </MenuItem>
                    {names.map((name) => (
                        <MenuItem
                            key={name}
                            value={name}
                            style={getStyles(name, personName, theme)}
                        >
                            {name}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>
        </div>
    );
}


export function SingleSelect({children, ...props}) {
    const {link, title, onClick, isSelect, setSelect} = props
    const [age, setAge] = React.useState('');

    const handleChange = (event) => {
        setAge(event.target.value);
        console.log(event.target.value);
    };

    return (
        <FormControl sx={{m: 1, minWidth: 120}}>
            <Select
                value={age}
                onChange={handleChange}
                displayEmpty
                inputProps={{'aria-label': 'Without label'}}
                sx={{
                    width: 160,
                    height: 31,
                    borderRadius: 19,
                    background: "#F5F6F8",
                    border: "none",
                    textAlign: "center"
                }}
            >
                <MenuItem value="">
                    <div className={'selectStatusOnline'}></div>
                    <span>Mary Foster</span></MenuItem>
                <MenuItem value={10}>
                    <div className={'selectStatusOnline'}></div>
                    <span>Ten</span></MenuItem>
                <MenuItem value={20}>
                    <div className={'selectStatusOffline'}></div>
                    <span>Twenty</span></MenuItem>
                <MenuItem value={30}>
                    <div className={'selectStatusOffline'}></div>
                    <span>Thirty</span></MenuItem>
            </Select>
        </FormControl>
    )
}

export function SingleSelect2({children, ...props}) {
    const {link, title, onClick, isSelect, setSelect} = props
    const [age, setAge] = React.useState('');

    const handleChange = (event) => {
        setAge(event.target.value);
    };
    return (
        <Select
            labelId="demo-simple-select-label"
            id="demo-simple-select"
            value={age}
            label="Age"
            onChange={handleChange}

        >

            <MenuItem value={10}>username</MenuItem>
            <MenuItem value={20}>offline</MenuItem>
            <MenuItem value={30}>offline</MenuItem>
        </Select>

    )
}
