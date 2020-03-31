import 'babel-polyfill';
import main from './js/main.ts';
import 'buefy/dist/buefy.css'
import '@mdi/font/css/materialdesignicons.css'

main({
    "APIPrefix": "http://npad2:8080",
    "AssetsPrefix": "/",
    "PathPrefix": "",
    "PageTitle": "rezept"
});