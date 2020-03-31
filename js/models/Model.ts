export default class Model {
    static getTypeofProperty<T, K extends keyof T>(o: T, name: K) {
        return typeof o[name];
    }

    public assign(data: any) {
        for (let property in data) {
            if (property in this) {
                if (this[property] instanceof Date) {
                    this[property] = new Date(data[property]);
                } else {
                    this[property] = data[property];
                }
            }
        }
    }

    protected static async buildModels<T extends Model>(request: Promise<Response>, ModelClass: new () => T): Promise<T[]> {
        const response = await request;
        const data: any = await response.json();
        if (!response.ok) {
            throw data.Error;
        }

        let models: T[] = new Array<T>(data.length);
        for (let i in data) {
            models[i] = new ModelClass();
            console.log(models[i]);
            models[i].assign(data[i]);
        }
        return models;
    }

    protected static async buildModel<T extends Model>(request: Promise<Response>, ModelClass: new () => T): Promise<T> {
        const response = await request;
        const data: any = await response.json();
        if (!response.ok) {
            throw data.Error;
        }

        let model: T = new ModelClass();
        model.assign(data);

        return model;
    }
}
