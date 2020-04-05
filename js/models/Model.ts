export default class Model {
  public assign(data: any): void {
    Object.keys(data).forEach((property) => {
      if (this[property] instanceof Date) {
        this[property] = new Date(data[property]);
      } else {
        this[property] = data[property];
      }
    });
  }

  protected static async buildModels<T extends Model>(
    request: Promise<Response>,
    ModelClass: new () => T,
  ): Promise<T[]> {
    const response = await request;
    const data: any = await response.json();
    if (!response.ok) {
      throw data.Error;
    }

    const models: T[] = new Array<T>(data.length);
    Object.keys(data).forEach((i) => {
      models[i] = new ModelClass();
      models[i].assign(data[i]);
    });

    return models;
  }

  protected static async buildModel<T extends Model>(
    request: Promise<Response>,
    ModelClass: new () => T,
  ): Promise<T> {
    const response = await request;
    const data: any = await response.json();
    if (!response.ok) {
      throw data.Error;
    }

    const model: T = new ModelClass();
    model.assign(data);

    return model;
  }
}
