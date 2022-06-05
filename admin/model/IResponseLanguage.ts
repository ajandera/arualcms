import Language from "~/model/Language";

export default interface IResponseLanguage {
  data: {
      languages: Language[],
      success: boolean,
      message: string,
      error: string
    }
}
