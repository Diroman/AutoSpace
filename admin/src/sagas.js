import { all } from 'redux-saga/effects';

// import { saga as authSaga } from './auth/duck/saga';

export function* sagas() {
    yield all([
        //authSaga(),
    ]);
}
