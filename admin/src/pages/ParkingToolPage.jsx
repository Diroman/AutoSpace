import * as React from 'react';
import ParkingTool from "../parking-tool/ParkingTool";
import {makeStyles} from "@material-ui/styles";


const useStyles = makeStyles(theme => ({
    root: {
        backgroundColor: '#eeeeee',
        minHeight: '100vh',
    },
}));

const ParkingToolPage = () => {

    const classes = useStyles();

    React.useEffect(() => {
        document.title = 'ParkingToolPage';
    });

    return (
        <div className={classes.root}>
            <ParkingTool />
        </div>
    );
};

export default ParkingToolPage;
