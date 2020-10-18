import React from 'react';
import clsx from 'clsx';
import { Card, CardContent, Divider, Typography, makeStyles } from '@material-ui/core';
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import {Link} from "react-router-dom";

const useStyles = makeStyles(theme => ({
    root: {
        width: '350px',
        display: 'flex',
        flexDirection: 'column',
        marginRight: theme.spacing(2),
    },
    statsItem: {
        alignItems: 'center',
        display: 'flex',
    },
    statsIcon: {
        marginRight: theme.spacing(1),
        color: theme.palette.primary.main,
    },
    button:{
        backgroundColor: theme.palette.action.active,
        color: 'white',
        '&:hover': {
            backgroundColor: theme.palette.action.active,
            color: '#fafafa',
        },
    }

}));

export const CameraCard = ({ address, index, className, ...rest }) => {

    const classes = useStyles();

    return (
        <Card
            className={clsx(classes.root, className)}
            {...rest}
        >
            <CardContent>
                <Typography
                    align="left"
                    color="textPrimary"
                    gutterBottom
                    variant="h4"
                >
                    Камера {index}
                </Typography>
                <Typography
                    align="left"
                    color="textPrimary"
                    gutterBottom
                    variant="h6"
                >
                    {address}
                </Typography>
            </CardContent>
            <Divider />
            <CardActions>
                <Button size="small" className={classes.button} variant='contained' component={Link} to={`/admin/parking/tool/${index}`}>
                    Просмотреть
                </Button>
            </CardActions>
        </Card>
    );
};

