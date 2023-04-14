import { Country } from "./Country";
import { Agent } from "./Representative";

export class Customer {
    id: number;
    name: string;
    country: Country;
    company: string;
    date: string;
    status: string;
    verified: boolean;
    activity: number;
    representative: Agent;
    balance: number;

    constructor() {
        this.id = 0;
        this.name = '';
        this.country = new Country();
        this.company = '';
        this.date = '';
        this.status = '';
        this.verified = false;
        this.activity = 0;
        this.representative = new Agent();
        this.balance = 0;
    }
}