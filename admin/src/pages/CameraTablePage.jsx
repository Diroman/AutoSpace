import * as React from 'react';
import {CameraTable} from "../camera-table/CameraTable";
import {makeStyles} from "@material-ui/styles";

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
            <CameraTable />
        </div>
    );
};

export default CameraTablePage;


