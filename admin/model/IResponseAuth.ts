export default interface IResponseAuth {
  data: {
      jwt: string,
      success: boolean,
      message: string
    }
}
