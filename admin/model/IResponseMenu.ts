import Menu from "~/model/Menu";

export default interface IResponseMenu {
    data: {
      success: boolean,
      message: string,
      menu: Menu[],
      error: string
    }
}
