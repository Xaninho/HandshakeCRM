
export default class EnumType {
    public ID: number;
    public Type: string;
    public Description: string;
    public Active: boolean;

    constructor(){
        this.ID = -1;
        this.Type = "";
        this.Description = "";
        this.Active = false;
    }
}