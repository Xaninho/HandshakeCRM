export default class Agent {
    public ID : number;
    public Name: string;
    public PhotoUrl: string;
    public Email: string;
    public PositionId: number;
    public PhoneNumber: number;

    constructor() {
        this.ID = -1;
        this.Name = '';
        this.PhotoUrl = '';
        this.Email = '';
        this.PositionId = 0;
        this.PhoneNumber = 0;
    }
}