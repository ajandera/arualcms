import Permission from "./Permission"

export default interface User {
  Username: string,
  Password: string,
  Id: string,
  ParentId?: string
  Permission?: Permission
  Name?: string
}
