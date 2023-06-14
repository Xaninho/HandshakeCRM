import axios from 'axios';

import Agent from '~~/types/Agent';
import Client from '~~/types/Contact';
import Company from '~~/types/Company';
import Contact from '~~/types/Contact';

// Define the base URL for the CRM API
const API_BASE_URL = 'http://localhost:3000/';

const apiInstance = axios.create({
    baseURL: API_BASE_URL,
});

const authCookie = useCookie("authHandshakeCRM");
apiInstance.interceptors.request.use((config) => {

    const authHandshakeCRM = authCookie.value;
    if (authHandshakeCRM) {
      config.headers.Authorization = `${authHandshakeCRM}`;
    }
    return config;
});

export default {

    // #region Get Requests
    
    getContacts() {
        return apiInstance.get('/contact');
    },
    getContact(id: number) {
        return apiInstance.get(`/contact/${id}`);
    },
    getCompanies() {
        return apiInstance.get('/company');
    },
    getCompany(id: number) {
        return apiInstance.get(`/company/${id}`);
    },
    getAgents(){
        return apiInstance.get('/agent');
    },
    getAgent(id: number) {
        return apiInstance.get(`/agent/${id}`);
    },

    getCurrencies() {
        return apiInstance.get('/currency');
    },
    getCurrency(id: number) {
        return apiInstance.get(`/currency/${id}`);
    },
    getCountries() {
        return apiInstance.get('/country');
    },
    getCountry(id: number) {
        return apiInstance.get(`/country/${id}`);
    },

    getEnumTypes(_searchString: string) {
        return apiInstance.get(`/enumType?typeOfEnum=${_searchString}`);
    },

    // #endregion
    // ----------

    // #region Post Requests

    createContact(_contact : Contact) {
        return apiInstance.post('/contact', _contact);
    },
    createCompany(_company : Company) {
        return apiInstance.post('/company', _company);
    },
    createAgent(_agent : Agent) {
        return apiInstance.post('/agent', _agent);
    },
    login(_email: string, _password: string) {
        return apiInstance.post('/login', {email: _email, password: _password});
    },

    // #endregion
    // ----------
    
    // #region Put Requests

    updateContact(_contact : Contact) {
        return apiInstance.put('/contact/' + _contact.ID, _contact);
    },
    updateCompany(_company : Company) {
        return apiInstance.put('/company/' + _company.ID, _company);
    },
    updateAgent(_agent : Agent) {
        return apiInstance.put('/agent/' + _agent.ID, _agent);
    },

    // #endregion
    // ----------
    
    // #region Delete Requests

    deleteContact(id: number) {
        return apiInstance.delete(`/contact/${id}`);
    },
    deleteCompany(id: number) {
        return apiInstance.delete(`/company/${id}`);
    },
    deleteAgent(id: number) {
        return apiInstance.delete(`/agent/${id}`);
    },

    // #endregion
    
}