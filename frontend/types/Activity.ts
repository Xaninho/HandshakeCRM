import Contact from "./Contact";
import EnumType from "./EnumType";
import User from "./User";

export default class Activity {
    public ID : number;
    public ContactID: number;
    public UserID: number;
    public Date: string;
    public Title: string;
    public Location: string;
    public Duration: string;
    public Notes: string;
    public OutcomeId: number;

    public Contact : Contact;
    public User : User;
    public Outcome : EnumType;

    constructor() {
        this.ID = -1;
        this.ContactID = 0;
        this.UserID = 0;
        this.Date = '';
        this.Title = '';
        this.Location = '';
        this.Duration = '';
        this.Notes = '';
        this.OutcomeId = 0;

        this.Contact = new Contact();
        this.Outcome = new EnumType();
        this.User = new User();
    }
}