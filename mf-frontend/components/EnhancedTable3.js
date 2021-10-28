import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableSortLabel from "@mui/material/TableSortLabel";
import Box from "@mui/material/Box";
import {visuallyHidden} from "@mui/utils";
import PropTypes from "prop-types";
import * as React from "react";
import Paper from "@mui/material/Paper";
import TableContainer from "@mui/material/TableContainer";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import {Pill} from "./Pill";

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



function EnhancedTable3Head(props) {
    const { order, orderBy, onRequestSort } =
        props;
    const createSortHandler = (property) => (event) => {
        onRequestSort(event, property);
    };

    const headCells2 = [
        {
            id: 'tag',
            numeric: false,
            disablePadding: false,
            label: 'Tag',
        },
        {
            id: 'total',
            numeric: true,
            disablePadding: false,
            label: 'Total',
        },

    ];

    return (
        <TableHead>
            <TableRow>
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

EnhancedTable3Head.propTypes = {
    onRequestSort: PropTypes.func.isRequired,
    order: PropTypes.oneOf(['asc', 'desc']).isRequired,
    orderBy: PropTypes.string.isRequired,
    rowCount: PropTypes.number.isRequired,
};

export function EnhancedTable3() {
    const [order, setOrder] = React.useState('asc');
    const [orderBy, setOrderBy] = React.useState('role');

    function createData(tag, total) {
        return {
            tag,
            total
        };
    }

    const rows = [
        createData(<Pill color="vip" size="longer">VIP</Pill>, 70),
        createData(<Pill color="vvip" size="longer">VVIP</Pill>, 60),
        createData(<Pill color="newCustomer" size="longer">New Customer</Pill>, 55),
        createData(<Pill color="promotions" size="longer">Promotions</Pill>, 58),
    ];

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };


    return (
        <Box sx={{ maxWidth: '1925px' }}>
            <Paper sx={{ width: '100%', mb: 2, boxShadow: "none" }}>
                <TableContainer>
                    <Table
                        sx={{ minWidth: 750 }}
                        aria-labelledby="tableTitle"
                    >
                        <EnhancedTable3Head
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

                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.tag}</TableCell>
                                            <TableCell sx={{padding: "26px", fontSize: "16px"}} align="left">{row.total}</TableCell>


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