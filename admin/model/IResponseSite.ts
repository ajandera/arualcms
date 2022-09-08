import Site from "~/model/Site";

export default interface IResponseSite {
  data: {
      sites: Site[],
      success: boolean,
      message: string,
      error: string
    }
}
