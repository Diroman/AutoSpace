import * as React from 'react';
import {CameraTable} from "../camera-table/CameraTable";
import {makeStyles} from "@material-ui/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import MenuIcon from '@material-ui/icons/Menu';

const useStyles = makeStyles(theme => ({
    root: {
        backgroundColor: '#eeeeee',
        minHeight: '100vh',
    },
}));

const CameraTablePage = () => {

    const classes = useStyles();

    React.useEffect(() => {
        document.title = 'CameraTablePage';
    });

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Toolbar variant="dense">
                    <IconButton edge="start" color="inherit" aria-label="menu">
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" color="inherit">
                        Адреса автостоянок
                    </Typography>
                </Toolbar>
            </AppBar>
            <CameraTable />
        </div>
    );
};

export default CameraTablePage;


