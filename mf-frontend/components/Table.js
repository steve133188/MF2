import {Pill, StatusPill} from "./Pill";
import {ContactType} from "./ContactType";
import {Badge} from "./Badge";
import {SingleBox} from "./Checkbox";

// Mui table
import * as React from 'react';
import PropTypes from 'prop-types';
import Box from '@mui/material/Box';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableSortLabel from '@mui/material/TableSortLabel';
import Paper from '@mui/material/Paper';
import { visuallyHidden } from '@mui/utils';

export function ContactTable() {
    return(
        <div className="contactTable">
            <table className="table">
                <thead>
                <tr className="headTr">
                    <th className="trCustomerID">Customer ID</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Channel</th>
                    <th>Tags</th>
                    <th>Assignee</th>
                </tr>
                </thead>
                <tbody>
                <tr className="bodyTr">
                    <td>0000001</td>
                    <td>Jacosdas dasdasdb</td>
                    <td>Thorntsadaasdason@com</td>
                    <td className="channel"><ContactType />lori.foster@mail.com</td>
                    <td><span className="tagsGroup"><Pill color="vip">VIP</Pill></span></td>
                    <td>Mary Foster</td>
                </tr>
                <tr>
                    <td>0000002</td>
                    <td>Jacosdas dasdasdb</td>
                    <td>Thorntsadaasdason@com</td>
                    <td className="channel"><ContactType />lori.foster@mail.com</td>
                    <td><span className="tagsGroup"><Pill color="vip">VIP</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill><Pill color="newCustomer">New Customer</Pill></span></td>
                    <td className="assignee">Mary Foster</td>
                </tr>
                </tbody>
            </table>
        </div>
    )
}

export function BroadcastTable() {
    return(
        <div className="broadcastTable">
            <table className="table">
                <thead>
                    <tr className="headTr">
                        <th className="trID">Name</th>
                        <th>Period</th>
                        <th>Group</th>
                        <th>Status</th>
                        <th>Created By</th>
                        <th>Created Date</th>
                        <th>Description</th>
                    </tr>
                </thead>
                <tbody>
                    <tr className="bodyTr">
                        <td>Broadcast 1</td>
                        <td>Sep 30, 2021 7:00 AM - Oct 30, 2021 7:00 AM</td>
                        <td><Badge color="gp1">Group1</Badge></td>
                        <td><StatusPill color="statusActive">Active</StatusPill></td>
                        <td>Lorem Ipsum</td>
                        <td>Mary Foster</td>
                        <td>s</td>
                    </tr>
                    <tr>
                        <td>Broadcast 2</td>
                        <td>Sep 30, 2021 7:00 AM - Oct 30, 2021 7:00 AM</td>
                        <td><Badge color="gp2">Group2</Badge></td>
                        <td><StatusPill color="statusPending">Pending</StatusPill></td>
                        <td>s</td>
                        <td>s</td>
                        <td>Mary Foster</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}

export function NormalTable(props) {
    return (
        <table className="normalTable">
            <tr>
                <th className={props.classname}><SingleBox></SingleBox></th>
                <th>Customer ID</th>
                <th>Name</th>
                <th>Team</th>
                <th>Channel</th>
                <th>Tags</th>
                <th>Assignee</th>
            </tr>
            {props.children}
        </table>
    )
}

function descendingComparator(a, b, orderBy) {
    if (b[orderBy] < a[orderBy]) {
        return -1;
    }
    if (b[orderBy] > a[orderBy]) {
        return 1;
    }
    return 0;
}

function getComparator(order, orderBy) {
    return order === 'desc'
        ? (a, b) => descendingComparator(a, b, orderBy)
        : (a, b) => -descendingComparator(a, b, orderBy);
}

// This method is created for cross-browser compatibility, if you don't
// need to support IE11, you can use Array.prototype.sort() directly
function stableSort(array, comparator) {
    const stabilizedThis = array.map((el, index) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0], b[0]);
        if (order !== 0) {
            return order;
        }
        return a[1] - b[1];
    });
    return stabilizedThis.map((el) => el[0]);
}



function EnhancedTableHead(props) {
    const { order, orderBy, onRequestSort } =
        props;
    const createSortHandler = (property) => (event) => {
        onRequestSort(event, property);
    };

    const headCells = [
        {
            id: 'name',
            numeric: false,
            disablePadding: false,
            label: 'Name',
        },
        {
            id: 'role',
            numeric: true,
            disablePadding: false,
            label: 'Role',
        },
        {
            id: 'status',
            numeric: true,
            disablePadding: false,
            label: 'status',
        },
        {
            id: 'averageDailyOnlineTime',
            numeric: true,
            disablePadding: false,
            label: 'Average Daily Online Time',
        },
        {
            id: 'assignedContacts',
            numeric: true,
            disablePadding: false,
            label: 'Assigned Contacts',
        },
        {
            id: 'activeContacts',
            numeric: true,
            disablePadding: false,
            label: 'Active Contacts',
        },
        {
            id: 'deliveredContacts',
            numeric: true,
            disablePadding: false,
            label: 'Delivered Contacts',
        },
        {
            id: 'unhandledContacts',
            numeric: true,
            disablePadding: false,
            label: 'Unhandled Contacts',
        },
        {
            id: 'totalMessagesSent',
            numeric: true,
            disablePadding: false,
            label: 'Total Messages Sent',
        },
        {
            id: 'averageResponseTime',
            numeric: true,
            disablePadding: false,
            label: 'Average Response Time',
        },
        {
            id: 'averageFirstResponseTime',
            numeric: true,
            disablePadding: false,
            label: 'Average First Response Time',
        },
    ];

    return (
        <TableHead>
            <TableRow>
                {headCells.map((headCell) => (
                    <TableCell
                        key={headCell.id}
                        align="left"
                        padding={headCell.disablePadding ? 'none' : 'normal'}
                        sortDirection={orderBy === headCell.id ? order : false}
                        sx={{padding: "26px"}}
                    >
                        <TableSortLabel
                            sx={{ fontWeight: "bold", color: "#495057"}}
                            active={orderBy === headCell.id}
                            direction={orderBy === headCell.id ? order : 'asc'}
                            onClick={createSortHandler(headCell.id)}
                        >
                            {headCell.label}
                            {orderBy === headCell.id ? (
                                <Box component="span" sx={visuallyHidden}>
                                    {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                                </Box>
                            ) : null}
                        </TableSortLabel>
                    </TableCell>
                ))}
            </TableRow>
        </TableHead>
    );
}

EnhancedTableHead.propTypes = {
    numSelected: PropTypes.number.isRequired,
    onRequestSort: PropTypes.func.isRequired,
    onSelectAllClick: PropTypes.func.isRequired,
    order: PropTypes.oneOf(['asc', 'desc']).isRequired,
    orderBy: PropTypes.string.isRequired,
    rowCount: PropTypes.number.isRequired,
};

export function EnhancedTable() {
    const [order, setOrder] = React.useState('asc');
    const [orderBy, setOrderBy] = React.useState('role');
    const [selected, setSelected] = React.useState([]);

    function createData(name, role, status, averageDailyOnlineTime, assignedContacts, activeContacts, deliveredContacts, unhandledContacts, totalMessagesSent, averageResponseTime, averageFirstResponseTime) {
        return {
            name,
            role,
            status,
            averageDailyOnlineTime,
            assignedContacts,
            activeContacts,
            deliveredContacts,
            unhandledContacts,
            totalMessagesSent,
            averageResponseTime,
            averageFirstResponseTime,
        };
    }

    const rows = [
        createData('Cupcake', "Admin", "Online", "09:00:00", 4, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Donut', "Admin", "Online", "09:00:00", 4, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Eclair', "Admin", "Online", "09:00:00", 6, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Frozen yoghurt', "Admin", "Online", "09:00:00", 4, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Gingerbread', "Admin", "Online", "09:00:00", 3, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Honeycomb', "Admin", "Online", "09:00:00", 6, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Ice cream sandwich', "Admin", "Online", "09:00:00", 4, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Jelly Bean', "Agent", "Online", "09:00:00", 0, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('KitKat', "Agent", "Online", "09:00:00", 7, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Lollipop', "Agent", "Online", "09:00:00", 0, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Marshmallow', "Agent", "Online", "09:00:00", 2, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Nougat', "Agent", "Offline", "09:00:00", 37, 5, 9, 8, 7, "6 mins", "6 mins"),
        createData('Oreo', "Agent", "Offline", "09:00:00", 4, 5, 9, 8, 7, "6 mins", "6 mins"),
    ];

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };

    const handleSelectAllClick = (event) => {
        if (event.target.checked) {
            const newSelecteds = rows.map((n) => n.name);
            setSelected(newSelecteds);
            return;
        }
        setSelected([]);
    };

    const handleClick = (event, name) => {
        const selectedIndex = selected.indexOf(name);
        let newSelected = [];

        if (selectedIndex === -1) {
            newSelected = newSelected.concat(selected, name);
        } else if (selectedIndex === 0) {
            newSelected = newSelected.concat(selected.slice(1));
        } else if (selectedIndex === selected.length - 1) {
            newSelected = newSelected.concat(selected.slice(0, -1));
        } else if (selectedIndex > 0) {
            newSelected = newSelected.concat(
                selected.slice(0, selectedIndex),
                selected.slice(selectedIndex + 1),
            );
        }

        setSelected(newSelected);
    };

    const isSelected = (name) => selected.indexOf(name) !== -1;

    return (
        <Box sx={{ maxWidth: '1925px' }}>
            <Paper sx={{ width: '100%', mb: 2, boxShadow: "none" }}>
                <TableContainer>
                    <Table
                        sx={{ minWidth: 750 }}
                        aria-labelledby="tableTitle"
                    >
                        <EnhancedTableHead
                            numSelected={selected.length}
                            order={order}
                            orderBy={orderBy}
                            onSelectAllClick={handleSelectAllClick}
                            onRequestSort={handleRequestSort}
                            rowCount={rows.length}
                        />
                        <TableBody>
                            {stableSort(rows, getComparator(order, orderBy))
                                .map((row, index) => {
                                    const isItemSelected = isSelected(row.name);
                                    const labelId = `enhanced-table-checkbox-${index}`;

                                    return (
                                        <TableRow
                                            hover
                                            onClick={(event) => handleClick(event, row.name)}
                                            role="checkbox"
                                            aria-checked={isItemSelected}
                                            tabIndex={-1}
                                            key={row.name}
                                            selected={isItemSelected}
                                        >

                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.name}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.role}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.status}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.averageDailyOnlineTime}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.assignedContacts}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.activeContacts}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.deliveredContacts}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.unhandledContacts}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.totalMessagesSent}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.averageResponseTime}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.averageFirstResponseTime}</TableCell>
                                        </TableRow>
                                    );
                                })}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Paper>
        </Box>
    );
}