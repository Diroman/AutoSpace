import * as React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import Box from '@material-ui/core/Box';
import {CameraCard} from "./CameraCard";
import {fetchData} from "../common/helpers/fetchData";
import axios from 'axios';

const useStyles = makeStyles(theme => ({
    root: {
        backgroundColor: '#eeeeee',
        minHeight: '100%',
        paddingRight: theme.spacing(6),
        paddingLeft: theme.spacing(6),
        paddingBottom: theme.spacing(3),
        paddingTop: theme.spacing(3),
    },
    news: {
        width: '100%',
        paddingBottom: theme.spacing(3),
    },
    productCard: {
        height: '100%',
        marginBottom: theme.spacing(2),
    },
}));

export const CameraTable = () => {
    const classes = useStyles();
    const [data, setData] = React.useState([]);

    React.useEffect(() => {
        const options = {
            headers: {
                'Content-Type': 'application/json',
            },
            mode: 'no-cors',
        };

        axios.get(`http://192.168.31.44:8080/all-cameras`)
            .then(res => {
                const data = res.data.cameras;
                setData(data);
            })

    },[]);

    return (
        <Grid container className={classes.root}>
            {data && data.map((item, index) =>
                <Box key={index} className={classes.news}>
                    <CameraCard
                        index={item.id}
                        className={classes.productCard}
                        address={item.address}
                        key={item.id}/>
                </Box>
            )}
        </Grid>
    );
};
