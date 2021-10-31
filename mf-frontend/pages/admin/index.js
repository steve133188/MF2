import Head from 'next/head'
import Image from 'next/image'
import {CheckboxGroup2, Search3} from "../../components/Input";
import {
    CancelButton,
    SelectButton,
    NormalButton,
    NormalButton2,
    TextWithIconButton
} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {NormalTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";
import {useState} from "react";
import Avatar from "@mui/material/Avatar";
import {Pill} from "../../components/Pill";
import {Checkbox1, CheckboxPill, SingleBox} from "../../components/Checkbox"
import {EditColumnPopper} from "../../components/EditColumnPopper";
import {MultipleSelectPlaceholder} from "../../components/Select";
import {AddPopper} from "../../components/AddPopper";
import {DeletePopper} from "../../components/DeletePopper";
import * as React from "react";
import {Dropzone} from "../../components/ImportContact";
import Box from '@mui/material/Box';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import {ColumnCheckbox} from "../../components/Checkbox";
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import FormControl from '@mui/material/FormControl';
import OutlinedInput from "@mui/material/OutlinedInput";
import {useTheme} from "@mui/material/styles";
import {DragDropContext, Droppable, Draggable} from 'react-beautiful-dnd';
import {BlueMenu2} from "../../components/BlueMenu";

export default function Admin() {
    const [isSelectRow, setSelectRow] = useState({"all": false});

    function toggleSelectRow() {
        setSelectRow(!isSelectRow);
    }

    function handleSelect(key, value) {
        setSelectRow(isSelectRow.key == !value, ...isSelectRow)
        if (key == "all") {

        }
    }

    const [isFillCheckbox, setIsFillCheckbox] = useState(false);

    function toggleFill() {
        setIsFillCheckbox(!isFillCheckbox);
    }

    const [open, setOpen] = React.useState(false);

    const handleClick = () => {
        setOpen((prev) => !prev);
    };

    const handleClickAway = () => {
        setOpen(false);
    };

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
    const [openDeletePopper, setOpenDeletePopper] = React.useState(false);

    const handleClickDeletePopper = () => {
        setOpenDeletePopper((prev) => !prev);
    };

    const handleClickAwayDeletePopper = () => {
        setOpenDeletePopper(false);
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

    const names = [
        'Team A',
        'Team B',
        'Team C',
        'Team D',
        'Team E',
    ];
    const selects = [
        {
            selectTitle: 'Team A'
        }
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

    const contactColumns = [
        {
            columnName: 'Customer ID'
        },
        {
            columnName: 'Name'
        },
        {
            columnName: 'Team'
        },
        {
            columnName: 'Channel'
        },
        {
            columnName: 'Tags'
        },
        {
            columnName: 'Assignee'
        }
    ];

    const [columns, updateColumns] = useState(contactColumns);

    function handleOnDragEnd(result) {
        if (!result.destination) return;

        const items = Array.from(columns);
        const [reorderedItem] = items.splice(result.source.index, 1);
        items.splice(result.destination.index, 0, reorderedItem);

        updateColumns(items);
    }

    return (
        <div className="admin-layout">
            {/*<Dropzone/>*/}
            <BlueMenu2 />
            <div className="rightContent">
                <div className="contactsContainer">
                    <div className="topBar">
                        <div className="searchBar">
                            <Search3 type="search">Search</Search3>
                        </div>
                        <div className="buttonGrp">
                            {isSelectRow ? (
                                <span onClick={toggleSelectRow}><SelectButton/></span>
                            ) : (
                                <span onClick={toggleSelectRow}><CancelButton/></span>
                            )}
                            <NormalButton2>+ New Agent</NormalButton2>
                        </div>
                    </div>
                    <div className="navbarPurple">
                        <div className="selectButtonGroup">
                            {selects.map(({selectTitle}) => {
                                return (
                                    <div className="multipleSelectPlaceholder" key={selectTitle}>
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
                                                            return <span>{selectTitle}</span>;
                                                        }
                                                        return selected.join('');
                                                    }}
                                                    MenuProps={MenuProps}
                                                    inputProps={{'aria-label': 'Without label'}}
                                            >
                                                <MenuItem disabled value="">
                                                    <span>{selectTitle}</span>
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
                            })}
                            {/*    */}
                        </div>

                        <div className="tagsButtonGroup">
                            <div>
                                <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="#2198fa"
                                     className="bi bi-pencil" viewBox="0 0 16 16">
                                    <path
                                        d="M12.146.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1 0 .708l-10 10a.5.5 0 0 1-.168.11l-5 2a.5.5 0 0 1-.65-.65l2-5a.5.5 0 0 1 .11-.168l10-10zM11.207 2.5 13.5 4.793 14.793 3.5 12.5 1.207 11.207 2.5zm1.586 3L10.5 3.207 4 9.707V10h.5a.5.5 0 0 1 .5.5v.5h.5a.5.5 0 0 1 .5.5v.5h.293l6.5-6.5zm-9.761 5.175-.106.106-1.528 3.821 3.821-1.528.106-.106A.5.5 0 0 1 5 12.5V12h-.5a.5.5 0 0 1-.5-.5V11h-.5a.5.5 0 0 1-.468-.325z"/>
                                </svg>
                            </div>
                            <div className="deletePopperContainer">
                                <ClickAwayListener onClickAway={handleClickAwayDeletePopper}>
                                    <Box sx={{position: 'relative'}}>

                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="#f46a6b"
                                             cursor="pointer"
                                             className="bi bi-trash" viewBox="0 0 16 16"
                                             onClick={handleClickDeletePopper}>
                                            <path
                                                d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                                            <path fillRule="evenodd"
                                                  d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                                        </svg>
                                        {openDeletePopper ? (
                                            <Box sx={deleteStyles}>
                                                Delete 2 contacts? <br/>
                                                All conversation history will also be erased.
                                                <div className="controlButtonGroup">
                                                    <NormalButton2>Delete</NormalButton2>
                                                    <span
                                                        onClick={handleClickDeletePopper}><CancelButton></CancelButton></span>
                                                </div>
                                            </Box>
                                        ) : null}
                                    </Box>
                                </ClickAwayListener>
                            </div>
                            {/*    */}
                        </div>
                    </div>
                    <div className="broadcastTable">
                        <table className="table">
                            <thead>
                            <tr className="headTr">

                                <th className="trID">Name</th>
                                <th>Role</th>
                                <th>Email</th>
                                <th>Phone</th>
                                <th>No. Of Leads</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr className="bodyTr">

                                <td style={{display: "flex"}}>
                                    <div className="selectStatusOnline"></div>Harry Stewart
                                </td>
                                <td>Admin</td>
                                <td>Harry.stewart@gmail.com</td>
                                <td>+852 9765 0348</td>
                                <td>7</td>
                            </tr>
                            <tr>

                                <td style={{display: "flex"}}>
                                    <div className="selectStatusOffline"></div><span>Jasmine Miller</span></td>
                                <td>Agent</td>
                                <td>jasmine.miller@gmail.com</td>
                                <td>+852 9765 0348</td>
                                <td>6</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    <PaginationControlled/>
                </div>
            </div>

        </div>
    )
}