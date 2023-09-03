import Menu from "~/model/Menu";

export default interface IResponseMenu {
    data: {
      success: boolean,
      message: string,
      menu: Menu[],
      error: string
    }
}

export interface IResponseMenuRoot {
  data: {
    success: boolean,
    message: string,
    menu: Menu,
    error: string
  }
}
