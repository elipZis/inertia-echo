import {InertiaApp} from '@inertiajs/inertia-svelte'
import axios from 'axios'
import route from '../../node_modules/ziggy-js'

// Import Ziggy globally
window.route = route;
// axios.interceptors.request.use(
//     config => {
//         config.headers['Authorization'] = 'Bearer ' + localStorage.getItem('inertiaToken');
//         return config;
//     },
//     error => Promise.reject(error),
// );
// Add a 401 response interceptor
axios.interceptors.response.use(function (response) {
    return response;
}, function (error) {
    if (401 === error.response.status) {
        window.location = '/login';
    } else {
        return Promise.reject(error);
    }
});

const app = document.getElementById('app')

new InertiaApp({
    target: app,
    props: {
        initialPage: JSON.parse(app.dataset.page),
        resolveComponent: name => import(`./Pages/${name}.svelte`).then(module => module.default),
    },
});
