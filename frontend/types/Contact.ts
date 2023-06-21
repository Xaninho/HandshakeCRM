import Agent from "./User";
import Company from "./Company";
import EnumType from "./EnumType";

export default class Contact {

    public ID : number;
    public CompanyId: number;

    public Name: string;
    public Email: string;
    public PhoneNumber: number;
    public Aniversary: string;
    public DispositionId: number;

    public Company : Company;
    public Disposition : EnumType;

    constructor() {
        this.ID = -1;
        this.Name = '';
        this.Email = '';
        this.PhoneNumber = 0;
        this.Aniversary = '';
        this.CompanyId = 0
        this.DispositionId = 0;
        this.Company = new Company();
        this.Disposition = new EnumType();	
    }
}