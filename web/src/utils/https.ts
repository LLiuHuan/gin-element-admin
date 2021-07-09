import HttpClient, {HttpClientConfig} from 'axios-mapper';
import {getToken} from "./cookies";

const https = (hasToken: boolean = true) => {
    const config: HttpClientConfig = {
        baseURL: 'http://127.0.0.1:8888',
        headers: {
            Authorization: hasToken ? 'Bearer ' + getToken() : '',
        },
    };

    return new HttpClient(config);
}


// // Add a request interceptor
// https.httpClient.interceptors.request.use(function (config) {
//     // Do something before request is sent
//     return config;
// }, function (error) {
//     // Do something with request error
//     return Promise.reject(error);
// });
//
// // Add a response interceptor
// https.httpClient.interceptors.response.use(function (response) {
//     // Do something with response data
//     return response;
// }, function (error) {
//     // Do something with response error
//     return Promise.reject(error);
// });

export default https;
