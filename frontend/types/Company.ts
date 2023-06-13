import Country from "./Country";
import Currency from "./Currency";

export default class Company {
    public ID : number;
    public NIF: number;
    public Name: string;
    public CountryId: number;
    public HasElectronicBilling: boolean;
    public CurrencyId: number;
    public PhotoURL: string;
    public PostalCode: string;
    public Address: string;
    public Notes: string;

    public Country : Country;
    public Currency : Currency;

    constructor() {
        this.ID = -1;
        this.NIF = 0;
        this.Name = '';
        this.CountryId = 0;
        this.CurrencyId = 0;
        this.HasElectronicBilling = false;
        this.PhotoURL = '';
        this.PostalCode = '';
        this.Address = '';
        this.Notes = '';

        this.Country = new Country();
        this.Currency = new Currency();
    }
    
}