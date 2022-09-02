export default interface Post {
    Src?: File;
    File?: File;
    Body: any;
    Title: any;
    Excerpt: any;
    Published: any;
    Id: string;
    Meta: {
      Title: any,
      Keywords: any,
      Description: any
    }
}
