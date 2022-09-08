export default interface IResponseFiles {
    data: {
      success: boolean,
      message: string,
      src: string,
      files: File[],
      file: string,
      error: string
    }
}
