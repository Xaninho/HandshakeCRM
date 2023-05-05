export default class Country {
    public ID: number;
    public Name: string;
    public Code: string;
    public FlagUrl?: string;

    constructor() {
        this.ID = -1;
        this.Name = "";
        this.Code = "";
        this.FlagUrl = "";
    } 
}