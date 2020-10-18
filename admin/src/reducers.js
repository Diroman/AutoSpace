import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

//import { IState as IAuth, reducer as authReducer } from '@app/auth/duck/reducer';

export const createRootReducer = (history) => combineReducers({
    router: connectRouter(history),
    //
    // auth: authReducer,
});
