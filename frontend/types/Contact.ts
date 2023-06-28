import Agent from "./User";
import Company from "./Company";
import EnumType from "./EnumType";
import Activity from "./Activity";

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
    public AssociatedActivities : Array<Activity>;

    public isLoading : boolean = false;

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
        this.AssociatedActivities = new Array<Activity>();
    }
}