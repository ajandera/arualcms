import Post from "~/model/Post";

export default interface IResponsePosts {
    data: {
      success: boolean,
      post: Post[],
      posts: Post[],
      message: string,
      error: string
    }
}
