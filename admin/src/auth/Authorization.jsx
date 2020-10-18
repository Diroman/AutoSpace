import React from 'react';

import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';

import logo from '../assets/img/logo.png';
import main from '../assets/img/main.png';
import { makeStyles } from '@material-ui/core/styles';
import { InputAdornment } from '@material-ui/core';
import IconButton from '@material-ui/core/IconButton';
import { Visibility, VisibilityOff } from '@material-ui/icons';
import {fetchData} from "../common/helpers/fetchData";

const useStyles = makeStyles(theme => ({
    root: {
        height: '100vh',
        backgroundColor: '#eeeeee',
    },
    image: {
        backgroundImage: `url(${main})`,
        backgroundRepeat: 'no-repeat',
        backgroundColor: theme.palette.grey[50],
        backgroundSize: 'cover',
        backgroundPosition: 'center',
    },
    paper: {
        margin: theme.spacing(12, 4),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    picture: {
        margin: theme.spacing(1),
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
        backgroundColor: theme.palette.action.active,
        color: 'white',
        '&:hover': {
            backgroundColor: theme.palette.action.active,
            color: '#fafafa',
        },
    },
}));

export const Authorization = ({ isAuth, authorize }) => {

    const classes = useStyles();
    const [values, setValues] = React.useState({
        username: '',
        password: '',
        showPassword: false,
    });

    const handleChange = (prop) => (e) => {
        setValues({ ...values, [prop]: e.target.value });
    };

    const onSubmit = (e) => {
        e.preventDefault();
        fetchData(`http://192.168.31.44:8080/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                    login: values.username,
                    password: values.password,
                }
            ),
            mode: 'no-cors',
        })
            .then(data => console.log(data))
            .catch(err => console.error(err));
    };

    const handleClickShowPassword = () => {
        setValues({ ...values, showPassword: !values.showPassword });
    };

    const handleMouseDownPassword = (e) => {
        e.preventDefault();
    };

    return (
        <Grid container component='main' className={classes.root}>

            <Grid item xs={false} sm={4} md={7} className={classes.image}/>
            <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>

                <div className={classes.paper}>
                    <img className={classes.picture} src={logo} alt='Logo' height='150px'/>
                    <form className={classes.form} noValidate>
                        <TextField
                            variant='outlined'
                            margin='normal'
                            required
                            fullWidth
                            id='username'
                            label='Username'
                            name='username'
                            autoComplete='username'
                            onChange={handleChange('username')}
                        />
                        <TextField
                            variant='outlined'
                            margin='normal'
                            required
                            fullWidth
                            name='password'
                            label='Пароль'
                            type={values.showPassword ? 'text' : 'password'}
                            value={values.password}
                            id='password'
                            autoComplete='password'
                            onChange={handleChange('password')}
                            InputProps={{
                                endAdornment:
                                    <InputAdornment position='end'>
                                        <IconButton
                                            aria-label='toggle password visibility'
                                            onClick={handleClickShowPassword}
                                            onMouseDown={handleMouseDownPassword}
                                            edge='end'
                                        >
                                            {values.showPassword ? <Visibility/> : <VisibilityOff/>}
                                        </IconButton>
                                    </InputAdornment>
                                ,
                            }}
                        />
                        <Button
                            type='submit'
                            fullWidth
                            variant='contained'
                            color='inherit'
                            className={classes.submit}
                            onClick={onSubmit}
                        >Войти</Button>

                    </form>
                </div>
            </Grid>
        </Grid>
    );
};
