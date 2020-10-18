import React from 'react';

import { hot } from 'react-hot-loader';
import AuthPage from "./pages/AuthPage";
import ParkingToolPage from "./pages/ParkingToolPage";

import ParkingListPage from "./pages/CameraTablePage";
import CssBaseline from "@material-ui/core/CssBaseline";

import { applyMiddleware, createStore } from 'redux';
import { Provider } from 'react-redux';
import { createBrowserHistory } from 'history';
import { composeWithDevTools } from 'redux-devtools-extension';
import createSagaMiddleware from 'redux-saga';
import { ConnectedRouter, routerMiddleware } from 'connected-react-router';

import { createRootReducer } from './reducers';
import { sagas } from './sagas';
import { Layout } from './layout/Layout';
import ErrorPage from "./pages/ErrorPage";
import { Route, Switch, Redirect } from 'react-router-dom';

const sagaMiddleware = createSagaMiddleware();

export const history = createBrowserHistory();

const store = createStore(
    createRootReducer(history),
    composeWithDevTools(
        applyMiddleware(
            sagaMiddleware,
            routerMiddleware(history),
        ),
    ),
);

sagaMiddleware.run(sagas);

const App = () => {
 return(
     <React.Fragment>
         <CssBaseline/>
         <Provider store={store}>
             <ConnectedRouter history={history}>
                 <Layout>
                     <Switch>
                         <Route exact path="/admin" component={AuthPage} />
                         <Route exact path="/admin/parking" component={ParkingListPage} />
                         <Route exact path="/admin/parking/tool/:id" component={ParkingToolPage} />
                         <Route component={ErrorPage} />
                         <Redirect to='/404'/>
                     </Switch>
                 </Layout>
             </ConnectedRouter>
         </Provider>
     </React.Fragment>
 )
};

export default hot(module)(App);

