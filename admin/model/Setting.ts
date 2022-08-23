export default interface Setting {
    Key: string;
    Value: SettingObject;
    Id: string;
}

interface SettingObject {
  [key: string]: string | number;
}
