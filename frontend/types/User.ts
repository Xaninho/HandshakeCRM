import EnumType from "./EnumType";

export default class User {
    public ID : number;
    public Name: string;
    public Email: string;
    public PhoneNumber: number;
    public RoleId: number;

    public Role : EnumType;

    constructor() {
        this.ID = -1;
        this.Name = '';
        this.Email = '';
        this.PhoneNumber = 0;
        this.RoleId = 0;
        
        this.Role = new EnumType();
    }
}