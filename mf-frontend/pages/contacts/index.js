import Head from 'next/head'
import Image from 'next/image'
import {CheckboxGroup2, Search3} from "../../components/Input";
import {
    CancelButton,
    ToggleButton,
    SelectButton,
    NormalButton,
    NormalButton2,
    TextWithIconButton
} from "../../components/Button";
import {NavbarPurple} from "../../components/NavbarPurple";
import {NormalTable} from "../../components/Table";
import {PaginationControlled} from "../../components/Pagination";
import {useState} from "react";
import {TableItem} from "../../components/Table"
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

export default function Contacts() {
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
    const selects = [
        {
            selectTitle: 'Agent',
            selectItems: [
                'Oliver Hansen',
                'Van Henry',
                'April Tucker',

            ]
        },
        {
            selectTitle: 'Team',
            selectItems: [
                'Ralph Hubbard',
                'Omar Alexander',
            ]
        },
        {
            selectTitle: 'Tags',
            selectItems: [
                'Carlos Abbott',
                'Miriam Wagner',
            ]
        },
        {
            selectTitle: 'Channel',
            selectItems: [
                'Bradley Wilkerson',
                'Virginia Andrews',
                'Kelly Snyder',
            ]
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
        <div className="contacts-layout">
            {/*<Dropzone/>*/}
            <div className="rightContent">
                <div className="contactsContainer">
                    <div className="topBar">
                        <div className="searchBar">
                            <Search3 type="search">Search</Search3>
                        </div>
                        <div className="buttonGrp">
                            {isSelectRow ? (
                                <span onClick={toggleSelectRow}><CancelButton/></span>
                            ) : (
                                <span onClick={toggleSelectRow}><SelectButton/></span>
                            )}
                            {/*<EditColumnPopper/>*/}
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

                                                <DragDropContext onDragEnd={handleOnDragEnd}>
                                                    <Droppable droppableId="columns">
                                                        {(provided) => (
                                                            <ul className="columnGroup" {...provided.droppableProps}
                                                                ref={provided.innerRef}>
                                                                {columns.map(({columnName}, index) => {
                                                                    return (
                                                                        <Draggable key={columnName}
                                                                                   draggableId={columnName}
                                                                                   index={index}>
                                                                            {(provided) => (
                                                                                <li ref={provided.innerRef} {...provided.draggableProps} {...provided.dragHandleProps}
                                                                                    className="columnCheckboxContainer">
                                                                                    <img
                                                                                        src="icon-columnControl.svg"
                                                                                        alt=""/>{columnName}
                                                                                    <input type="checkbox"/>
                                                                                    <span className="checkmark"></span>
                                                                                </li>
                                                                            )}
                                                                        </Draggable>
                                                                    );
                                                                })}
                                                                {provided.placeholder}
                                                            </ul>
                                                        )}
                                                    </Droppable>
                                                </DragDropContext>
                                            </Box>
                                        ) : null}

                                    </Box>
                                </ClickAwayListener>
                            </div>
                            {/**/}
                            <NormalButton>Import</NormalButton>
                            <NormalButton2>+ New Contact</NormalButton2>
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
                            {/*<AddPopper/>*/}

                            <div className="addPopperContainer">
                                <ClickAwayListener onClickAway={handleClickAwayAddPopper}>
                                    <Box sx={{position: 'relative'}}>
                                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25"
                                             fill="currentColor"
                                             className="bi bi-tag" viewBox="0 0 16 16" onClick={handleClickAddPopper}>
                                            <path
                                                d="M6 4.5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0zm-1 0a.5.5 0 1 0-1 0 .5.5 0 0 0 1 0z"/>
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
                                                    <CheckboxPill color={"vip"} checked={"checked"}>VIP</CheckboxPill>
                                                    <CheckboxPill color={"vip"} checked={""}>VIP</CheckboxPill>
                                                </CheckboxGroup2>
                                            </Box>
                                        ) : null}
                                    </Box>
                                </ClickAwayListener>
                            </div>

                            {/**/}
                            <div>
                                <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="#2198fa"
                                     cursor="pointer"
                                     className="bi bi-upload" viewBox="0 0 16 16">
                                    <path
                                        d="M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5z"/>
                                    <path
                                        d="M7.646 1.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708L8.5 2.707V11.5a.5.5 0 0 1-1 0V2.707L5.354 4.854a.5.5 0 1 1-.708-.708l3-3z"/>
                                </svg>
                            </div>

                            {/*<DeletePopper/>*/}
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
                    {/*<NormalTable classname={isSelectRow ? null : "checkBox"}>*/}
                    <table className="normalTable">
                        <tr>
                            <th className={isSelectRow ? null : "checkBox"}><SingleBox key={"all"} isSelect={false}
                                                                                       onClick={toggleFill}
                                                                                       onChange={e => handleSelect(e.target.key, e.target.isSelect)}></SingleBox>
                            </th>
                            <th>Customer ID</th>
                            <th>Name</th>
                            <th>Team</th>
                            <th>Channel</th>
                            <th>Tags</th>
                            <th>Assignee</th>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                        <tr>
                            <TableItem classname={isSelectRow ? null : "checkBox"}><SingleBox
                                fill={isFillCheckbox ? "fillCheckbox" : null}></SingleBox></TableItem>
                            <TableItem>0000001</TableItem>
                            <TableItem>
                                <div className="nameGroup"><Avatar alt="Remy Sharp"
                                                                   src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"/>Debra
                                    Patel
                                </div>
                            </TableItem>
                            <TableItem><Pill color="teamA">Team A</Pill></TableItem>
                            <TableItem>
                                <div className="channel"><img
                                    src="https://www.ethnicmusical.com/wp-content/uploads/2021/02/Whatsapp-PNG-Image-79477.png"
                                    alt=""/></div>
                            </TableItem>
                            <TableItem>
                                <div className="tagsGroup"><Pill color="lightBlue">VIP</Pill><Pill color="lightPurple">New
                                    Customer</Pill></div>
                            </TableItem>
                            <TableItem>
                                <div className="assigneeGroup">
                                    <Pill color="lightYellow" size="roundedPill size30">MF</Pill>
                                    <Pill color="lightBlue" size="roundedPill size30">AX</Pill>
                                    <Pill color="lightGreen" size="roundedPill size30">DS</Pill>
                                    <Pill color="lightPurple" size="roundedPill size30">EW</Pill>
                                    <Pill color="lightRed" size="roundedPill size30">KA</Pill>
                                </div>
                            </TableItem>
                        </tr>
                    </table>
                    {/*</NormalTable>*/}
                    <PaginationControlled/>
                </div>
            </div>

        </div>
    )
}