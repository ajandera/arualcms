export default interface User {
  username: string,
  password: string,
  _id: {
    $oid: string
  }
}
