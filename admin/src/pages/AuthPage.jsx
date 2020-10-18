import * as React from 'react';
import {Authorization} from "../auth/Authorization";

const AuthPage = () => {

    React.useEffect(() => {
        document.title = 'AuthPage';
    });

    return (
        <Authorization />
    );
};

export default AuthPage;
