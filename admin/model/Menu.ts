export default interface Menu {
    Id: string;
    Name: any;
    Url: string;
    Children: Menu[];
    PostId: string;
    PageId: string;
}
