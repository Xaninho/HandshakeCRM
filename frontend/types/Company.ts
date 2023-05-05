import Country from "./Country";
import Currency from "./Currency";
import EnumType from "./EnumType";

export default class Company {
    public ID : number;
    public Name: string;
    public CountryId: number;
    public StateId: number;
    public CurrencyId: number;
    public HasElectronicBilling: boolean;
    public PhotoURL: string;
    public PostalCode: string;
    public Address: string;
    public Notes: string;

    public Country : Country;
    public Currency : Currency;
    public State : EnumType;

    constructor() {
        this.ID = -1;
        this.Name = '';
        this.CountryId = 0;
        this.StateId = 0;
        this.CurrencyId = 0;
        this.HasElectronicBilling = false;
        this.PhotoURL = '';
        this.PostalCode = '';
        this.Address = '';
        this.Notes = '';

        this.Country = new Country();
        this.Currency = new Currency();
        this.State = new EnumType();
    }
    
}