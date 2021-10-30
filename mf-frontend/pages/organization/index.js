import {BlueMenu} from "../../components/BlueMenu";
import {Search3} from "../../components/Input";
import {NormalButton, NormalButton2} from "../../components/Button";
import {PaginationControlled} from "../../components/Pagination";
import * as React from "react";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableSortLabel from "@mui/material/TableSortLabel";
import Box from "@mui/material/Box";
import {visuallyHidden} from "@mui/utils";
import PropTypes from "prop-types";
import Paper from "@mui/material/Paper";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import {SingleBox} from "../../components/Checkbox";

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



function EnhancedTable2Head(props) {
    const { order, orderBy, onRequestSort } =
        props;
    const createSortHandler = (property) => (event) => {
        onRequestSort(event, property);
    };

    const headCells2 = [
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
            id: 'email',
            numeric: true,
            disablePadding: false,
            label: 'Email',
        },
        {
            id: 'phone',
            numeric: true,
            disablePadding: false,
            label: 'Phone',
        },
        {
            id: 'noOfLeads',
            numeric: true,
            disablePadding: false,
            label: 'No. Of Leads'
        }
    ];

    return (
        <TableHead>
            <TableRow >
                <th style={{ width: "30px", textAlign: "center", borderBottom: "1px solid #e0e0e0"}}><SingleBox /></th>
                {headCells2.map((headCell2) => (
                    <TableCell
                        key={headCell2.id}
                        align="left"
                        padding={headCell2.disablePadding ? 'none' : 'normal'}
                        sortDirection={orderBy === headCell2.id ? order : false}
                        sx={{padding: "26px"}}
                    >
                        <TableSortLabel
                            sx={{ fontWeight: "bold", color: "#495057"}}
                            active={orderBy === headCell2.id}
                            direction={orderBy === headCell2.id ? order : 'asc'}
                            onClick={createSortHandler(headCell2.id)}
                        >
                            {headCell2.label}
                            {orderBy === headCell2.id ? (
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

EnhancedTable2Head.propTypes = {
    onRequestSort: PropTypes.func.isRequired,
    order: PropTypes.oneOf(['asc', 'desc']).isRequired,
    orderBy: PropTypes.string.isRequired,
    rowCount: PropTypes.number.isRequired,
};

export default function Organization() {
    const [order, setOrder] = React.useState('asc');
    const [orderBy, setOrderBy] = React.useState('role');

    function createData(name, role, email, phone, noOfLeads) {
        return {
            name,
            role,
            email,
            phone,
            noOfLeads
        };
    }

    const rows = [
        createData(<div style={{display: "flex"}}><div className="selectStatusOnline"></div><span>Harry Stewart</span></div>, "Admin", "Harry.stewart@gmail.com", "+852 9765 0348", 7),

        createData(<div style={{display: "flex"}}><div className="selectStatusOffline"></div><span>Jasmine Miller</span></div>, "Agent", "jasmine.miller@gmail.com", "+852 9765 0348", 6),

        createData(<div style={{display: "flex"}}><div className="selectStatusOnline"></div><span>Chris Chavez</span></div>, "Agent", "jasmine.miller@gmail.com", "+852 9765 0348", 6),

        createData(<div style={{display: "flex"}}><div className="selectStatusOnline"></div><span>Harry Stewart</span></div>, "Agent", "Harry.stewart@gmail.com", "+852 9765 0348", 7),
    ];

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };

    return (
        <div className="organization-layout">
            <BlueMenu></BlueMenu>
            <div className="rightContent">
                <div className="broadcastContainer">
                    <div className="topBar">
                        <div className="searchBar">
                            <Search3 type="search">Search</Search3>
                        </div>
                        <div className="buttonGrp">
                            <NormalButton>Select</NormalButton>
                            <NormalButton2>+ New Team</NormalButton2>
                            <NormalButton2>+ New Division</NormalButton2>
                        </div>
                    </div>
                    <div className="navbarPurple">
                        <div>Team2 ( 8 )</div>
                        <NormalButton2>+ New Agent</NormalButton2>
                    </div>
                    <Box sx={{ maxWidth: '1925px' }}>
                        <Paper sx={{ width: '100%', mb: 2, boxShadow: "none" }}>
                            <TableContainer>
                                <Table
                                    sx={{ minWidth: 750 }}
                                    aria-labelledby="tableTitle"
                                >
                                    <EnhancedTable2Head
                                        order={order}
                                        orderBy={orderBy}
                                        onRequestSort={handleRequestSort}
                                        rowCount={rows.length}
                                    />
                                    <TableBody>
                                        {stableSort(rows, getComparator(order, orderBy))
                                            .map((row, index) => {
                                                const labelId = `enhanced-table-checkbox-${index}`;
                                                return (
                                                    <TableRow
                                                        hover
                                                        role="checkbox"
                                                        tabIndex={-1}
                                                        key={row.name}
                                                    >
                                                        <td style={{ width: "30px", textAlign: "center", borderBottom: "1px #e0e0e0 solid"}}><SingleBox /></td>
                                                        <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.name}</TableCell>
                                                        <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.role}</TableCell>
                                                        <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.email}</TableCell>
                                                        <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.phone}</TableCell>
                                                        <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.noOfLeads}</TableCell>
                                                    </TableRow>
                                                );
                                            })}
                                    </TableBody>
                                </Table>
                            </TableContainer>
                        </Paper>
                    </Box>
                    <PaginationControlled/>
                </div>
            </div>
        </div>
    )
}