import defaultsDeep from 'lodash-es/defaultsDeep';

export const fetchData = (url, options) => {
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    };
    return fetch(url, defaultsDeep(options, {
        headers,
    }));
};
