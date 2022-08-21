export default interface Post {
    src?: File;
    file?: File;
    body: any;
    title: any;
    excerpt: any;
    published: any;
    id: string;
    meta: {
      title: any,
      keywords: any,
      description: any
    }
}
