import { Country } from "./Country";
import { Representative } from "./Representative";

export class Customer {
    id: number;
    name: string;
    country: Country;
    company: string;
    date: string;
    status: string;
    verified: boolean;
    activity: number;
    representative: Representative;
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
        this.representative = new Representative();
        this.balance = 0;
    }
}