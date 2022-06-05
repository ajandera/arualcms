import Setting from "~/model/Setting";

export default interface IResponseSetting {
    data: {
      success: boolean,
      message: string,
      settings: Setting[],
      error: string
    }
}
