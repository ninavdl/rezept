import App from './components/App.vue';
import Vue from 'vue';
import Vuex from 'vuex';
import API from './models/API';

export default class Controller {
    constructor(config, router) {
        this.config = config;
        this.user = null;
        this.isLoggedIn = false;
        this.token = this.getCookie("token");
        this.router = router;

        (async () => {
            if (this.token != "") await this.getLoggedInUser();
        })().then(() => {
            this.initPage();
        })
    }

    requestObject(options) {
        if (this.token != "") {
            if ("headers" in options) {
                options.headers.Authorization = "Bearer " + this.token;
            } else {
                options.headers = {
                    Authorization: "Bearer " + this.token
                }
            }
        }
        console.log(options);
        return options
    }

    buildUri(uri) {
        return this.config.APIPrefix + "/" + uri
    }

    // add prefixes to a given URI and GET it
    async fetch(uri) {
        return fetch(this.buildUri(uri), this.requestObject({}));
    }

    // add prefixes to a given URI and DELETE it
    async delete(uri) {
        return fetch(this.buildUri(uri), this.requestObject({
            method: "DELETE"
        }));
    }

    // add prefixes to a given URI and POST the given data to it
    async post(uri, data) {
        return fetch(this.buildUri(uri), this.requestObject({
            method: "POST",
            body: JSON.stringify(data),
        }));
    };

    // add prefixes to a given URI and PUT the given data to it
    async put(uri, data) {
        return fetch(this.buildUri(uri), this.requestObject({
            method: "PUT",
            body: JSON.stringify(data)
        }));
    };

    async uploadImage(file, progress) {
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest();
            xhr.responseType = "json";
            xhr.open("PUT", this.buildUri("image"), true);
            xhr.onprogress = e => {
                progress(e.uploaded, e.fileSize);
            };
            xhr.onload = () => {
                if (xhr.status == 201) {
                    resolve(xhr.response);
                } else {
                    reject(xhr.response.Error);
                }
            }
            xhr.onerror = () => {
                reject("Error");
            }
            xhr.setRequestHeader("Content-Type", file.type);
            xhr.setRequestHeader("Authorization", "Bearer " + this.token);
            xhr.send(file);
        })
    }

    setCookie(cname, cvalue, exdays) {
        var d = new Date();
        d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
        var expires = "expires=" + d.toUTCString();
        document.cookie = cname + "=" + cvalue + ";" + expires + ";path=" + this.config.PathPrefix + "/";
    }

    getCookie(cname) {
        var name = cname + "=";
        var decodedCookie = decodeURIComponent(document.cookie);
        var ca = decodedCookie.split(';');
        for (var i = 0; i < ca.length; i++) {
            var c = ca[i];
            while (c.charAt(0) == ' ') {
                c = c.substring(1);
            }
            if (c.indexOf(name) == 0) {
                return c.substring(name.length, c.length);
            }
        }
        return "";
    }

    async getLoggedInUser() {
        API.getInstance().setToken(this.token);
        let req = await this.fetch("login")
        let data = await req.json();
        if (!req.ok) {
            alert(data.Error);
            return;
        }
        this.isLoggedIn = true;
        this.user = data;
    }

    initPage() {
        let controller = this;

        Vue.filter('round', (value, decimals) => Math.round(value * Math.pow(10, decimals)) / Math.pow(10, decimals))
        Vue.filter('formatDate', date => date.toLocaleDateString())

        let store = new Vuex.Store({
            state: {
                user: controller.user == null ? {} : controller.user,
                isLoggedIn: controller.user != null
            },
            mutations: {
                setUser(state, newUser) {
                    state.user = Object.assign({}, state.user, newUser);
                    state.isLoggedIn = "ID" in newUser
                }
            }
        })

        Vue.prototype.$controller = this;

        this.page = new Vue({
            router: this.router,
            store,
            render: createElement => createElement(App)
        });

        this.page.$mount("#app");
    }

    async login(username, password) {
        let req = await this.put("login", {
            Username: username,
            Password: password
        });
        let data = await req.json();
        if (!req.ok) {
            throw data.Error;
        }
        this.token = data.SessionID;
        API.getInstance().setToken(this.token);
        this.setCookie("token", data.SessionID, 365);
        await this.getLoggedInUser();
    }

    async logout() {
        let req = await this.delete("login");
        if (!req.ok) {
            throw data.Error;
        }

        this.setCookie("token", "", -1);
        this.token = "";
        this.user = null;
        this.isLoggedIn = false;
    }

    showLoadingScreen() {
    }

    hideLoadingScreen() {
    }
}