import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import Paper from '@material-ui/core/Paper';
import Checkbox from '@material-ui/core/Checkbox';
import DeleteIcon from '@material-ui/icons/Delete';


function desc(a, b, orderBy) {
    if (b[orderBy] < a[orderBy]) {
        return -1;
    }
    if (b[orderBy] > a[orderBy]) {
        return 1;
    }
    return 0;
}

function stableSort(array, cmp) {
    const stabilizedThis = array.map((el, index) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = cmp(a[0], b[0]);
        if (order !== 0) return order;
        return a[1] - b[1];
    });
    return stabilizedThis.map(el => el[0]);
}

function getSorting(order, orderBy) {
    return order === 'desc' ? (a, b) => desc(a, b, orderBy) : (a, b) => -desc(a, b, orderBy);
}
function displayActionButton(isDeleteButtonEnabled){
    if (isDeleteButtonEnabled){
        return (
            <TableCell component="th" id="delete-head" scope="row" padding="none" className="row-id">
                <span>Delete</span>
            </TableCell>
        )
    }
}

function EnhancedTableHead(props) {
    const { headCells, classes, onSelectAllClick, order, orderBy, numSelected, rowCount, onRequestSort, isDeleteButtonEnabled } = props;
    const createSortHandler = property => event => {
        onRequestSort(event, property);
    };

    return (
        <TableHead>
            <TableRow>
                <TableCell>
                    <Checkbox
                        indeterminate={numSelected > 0 && numSelected < rowCount}
                        checked={numSelected === rowCount}
                        onChange={onSelectAllClick}
                        inputProps={{ 'aria-label': 'select all' }}
                    />
                </TableCell>
                {headCells.map(headCell => (
                    <TableCell
                        key={headCell.id}
                        align={headCell.numeric ? 'right' : 'left'}
                        padding={headCell.disablePadding ? 'none' : 'default'}
                        sortDirection={orderBy === headCell.id ? order : false}
                        className={headCell.hidden ? classes.visuallyHidden : ""}
                    >
                        <TableSortLabel
                            active={orderBy === headCell.id}
                            direction={order}
                            onClick={createSortHandler(headCell.id)}
                        >
                            {headCell.label}
                            {orderBy === headCell.id ? (
                                <span className={classes.visuallyHidden}>
                                    {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                                </span>
                            ) : null}
                        </TableSortLabel>
                    </TableCell>
                ))}
                {displayActionButton(isDeleteButtonEnabled)}
            </TableRow>
        </TableHead>
    );
}

EnhancedTableHead.propTypes = {
    hidden: PropTypes.bool,
    classes: PropTypes.object.isRequired,
    numSelected: PropTypes.number.isRequired,
    onRequestSort: PropTypes.func.isRequired,
    onSelectAllClick: PropTypes.func.isRequired,
    order: PropTypes.oneOf(['asc', 'desc']).isRequired,
    orderBy: PropTypes.string.isRequired,
    rowCount: PropTypes.number.isRequired,
};

const useStyles = makeStyles(theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(3),
    },
    paper: {
        width: '100%',
        marginBottom: theme.spacing(2),
    },
    table: {
        minWidth: 750,
    },
    tableWrapper: {
        overflowX: 'auto',
    },
    visuallyHidden: {
        border: 0,
        clip: 'rect(0 0 0 0)',
        height: 1,
        margin: -1,
        overflow: 'hidden',
        padding: 0,
        position: 'absolute',
        top: 20,
        width: 1,
    },
}));

export default function EnhancedTable(props) {
    const { headCells, data, isDeleteButtonEnabled } = props;
    const classes = useStyles();
    const [order, setOrder] = React.useState('asc');
    const [orderBy, setOrderBy] = React.useState('created_at');
    const [selected, setSelected] = React.useState([]);

    const handleRequestSort = (event, property) => {
        const isDesc = orderBy === property && order === 'desc';
        setOrder(isDesc ? 'asc' : 'desc');
        setOrderBy(property);
    };

    const handleSelectAllClick = event => {
        if (event.target.checked) {
            const newSelecteds = data.map(n => n.name);
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

    const isSelected = name => {
        return selected.indexOf(name) !== -1;
    }

    const deleteLine = (event) => {
        const el = event.target.closest("tr");
        console.log(el)
    }

    const generateRow = (row, index) => {
        console.log(row);
        const isItemSelected = isSelected("row-id-" + index);
        const labelId = `enhanced-table-checkbox-${index}`;
        let cells = []
        for (let headCell in headCells) {
            if (headCells.hasOwnProperty(headCell)) {
                cells.push(
                    <TableCell component="td" id={labelId} scope="row" padding="none"  className={headCells[headCell].hidden ? classes.visuallyHidden : ""}>
                        <span className="ellipsis">{row[headCells[headCell].field]}</span>
                    </TableCell>
                )
            }
        }

        if (isDeleteButtonEnabled){
            cells.push(
                <TableCell component="td" scope="row" padding="none" className="row-id">
                    <DeleteIcon onClick={(event) => {deleteLine(event)} }/>
                </TableCell>
            )
        }

        return (
            <TableRow
                hover
                onClick={event => handleClick(event, "row-id-" + index)}
                role="checkbox"
                aria-checked={isItemSelected}
                key={"row-id-" + index}
                selected={isItemSelected}
            >
                <TableCell>
                    <Checkbox
                        checked={isItemSelected}
                        inputProps={{ 'aria-labelledby': labelId }}
                    />
                </TableCell>
                {
                    cells
                }
            </TableRow>
        );
    }
    return (
        <div className={classes.root}>
            <Paper className={classes.paper}>
                <div className={classes.tableWrapper}>
                    <Table
                        className={classes.table}
                        aria-labelledby="tableTitle"
                        size='medium'
                        aria-label="table"
                    >
                        <EnhancedTableHead
                            headCells={headCells}
                            classes={classes}
                            numSelected={selected.length}
                            order={order}
                            orderBy={orderBy}
                            onSelectAllClick={handleSelectAllClick}
                            onRequestSort={handleRequestSort}
                            rowCount={data.length}
                            isDeleteButtonEnabled={isDeleteButtonEnabled}
                        />
                        <TableBody>
                            {
                                stableSort(data, getSorting(order, orderBy))
                                    .map((row, index) => {
                                        return generateRow(row, index)
                                    })
                            }
                        </TableBody>
                    </Table>
                </div>
            </Paper>
        </div>
    );
}
