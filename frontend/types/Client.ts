import Agent from "./Agent";
import Company from "./Company";
import EnumType from "./EnumType";

export default class Client {

    public ID : number;
    public Name: string;
    public Email: string;
    public PhotoURL: string;
    public PhoneNumber: string;
    public Aniversary: string;
    public CompanyId: number;
    public FunctionId: number;
    public TypeId: number;
    public AssignedAgentId?: number;

    public Company : Company;
    public AssignedAgent : Agent;
    public Type : EnumType;
    public Position : EnumType;

    constructor() {
        this.ID = -1;
        this.Name = '';
        this.Email = '';
        this.PhotoURL = '';
        this.PhoneNumber = '';
        this.Aniversary = '';
        this.CompanyId = 0
        this.FunctionId = 0;
        this.TypeId = 0;
        this.AssignedAgentId = 0;
        this.Company = new Company();
        this.AssignedAgent = new Agent();
        this.Type = new EnumType();
        this.Position = new EnumType();
    }
}