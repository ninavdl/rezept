import 'babel-polyfill';
import main from './main';
import 'buefy/dist/buefy.css';
import '@mdi/font/css/materialdesignicons.css';

main({
  APIPrefix: process.env.API_URL,
  PathPrefix: process.env.BASE_URL,
  PageTitle: process.env.PAGE_TITLE,
});
