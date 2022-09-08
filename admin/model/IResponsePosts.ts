import Post from "~/model/Post";

export default interface IResponsePosts {
    data: {
      success: boolean,
      post: Post[],
      posts: Post[],
      files: any[],
      message: string,
      error: string
    }
}
