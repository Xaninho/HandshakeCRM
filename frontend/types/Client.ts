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
    }
}