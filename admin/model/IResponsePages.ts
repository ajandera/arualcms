import Page from "~/model/Page";

export default interface IResponsePages {
    data: {
      success: boolean,
      page: Page[],
      pages: Page[],
      message: string,
      error: string
    }
}
