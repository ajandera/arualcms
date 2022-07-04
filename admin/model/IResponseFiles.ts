export default interface IResponseFiles {
    data: {
      success: boolean,
      message: string,
      files: File[],
      file: File,
      error: string
    }
}
