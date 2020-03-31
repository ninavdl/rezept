export default class API {
    private static instance: API;
    private apiUrl: string;
    private token: string | null = null;

    private constructor(apiUrl: string) {
        this.apiUrl = apiUrl;
    }

    public static init(apiUrl: string) {
        API.instance = new API(apiUrl);
    }

    public static getInstance(): API {
        if (!API.instance) {
            throw "Must initialize API singleton first";
        }

        return API.instance;
    }

    public setToken(token: string | null): void {
        this.token = token;
    }

    protected buildURL(uri: string): string {
        return `${this.apiUrl}/${uri}`
    }

    requestObject(options: any): Object {
        if (this.token != null) {
            if ("headers" in options) {
                options.headers.Authorization = "Bearer " + this.token;
            } else {
                options.headers = {
                    Authorization: "Bearer " + this.token
                }
            }
        }
        return options
    }

    public async GET(uri: string): Promise<Response> {
        return this.request(uri, "GET");
    }

    public async DELETE(uri: string): Promise<Response> {
        return this.request(uri, "DELETE");
    }

    public async POST(uri: string, data?: any): Promise<Response> {
        return this.request(uri, "POST", data);
    }

    public async PUT(uri: string, data?: any): Promise<Response> {
        return this.request(uri, "PUT", data);
    }

    protected request(uri: string, method: string, data?: any): Promise<Response> {
        let requestObject: any = {
            headers: {},
            method: method,
            body: null
        };

        if (typeof data != "undefined") requestObject.body = JSON.stringify(data);
        if (this.token != null) {
            requestObject.headers.Authorization = `Bearer ${this.token}`;
        }

        return fetch(this.buildURL(uri), requestObject);
    }
}