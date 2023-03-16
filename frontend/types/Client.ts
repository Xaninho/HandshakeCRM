import { Country } from "./Country";

export class Client {

    public id : number;
    public firstName: string;
    public lastName: string;
    public displayName: string;
    public note?: string;
    public photoUrl?: string;
    public country: Country;
    public company: string;
    public rating?: number;

    constructor() {
        this.id = -1;
        this.firstName = "";
        this.lastName = "";
        this.displayName = this.firstName + " " + this.lastName;
        this.company = "";
        this.rating = 0;
        this.country = new Country();
    }
}