export default class Currency {
    public ID: number;
    public Code: string;
    public Name: string;
    public ExchangeRate: number;

    constructor() {
        this.ID = -1;
        this.Code = '';
        this.Name = '';
        this.ExchangeRate = 0;
    }
    
}