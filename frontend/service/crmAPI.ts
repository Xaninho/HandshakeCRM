import axios from 'axios';
import Agent from '~~/types/Agent';
import Client from '~~/types/Client';
import Company from '~~/types/Company';

// Define the base URL for the CRM API
const API_BASE_URL = 'http://localhost:3000/';

const apiInstance = axios.create({
    baseURL: API_BASE_URL,
});

export default {

    // #region Get Requests
    
    getClients() {
        return apiInstance.get('/client');
    },
    getClient(id: number) {
        return apiInstance.get(`/client/${id}`);
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

    createClient(_client : Client) {
        return apiInstance.post('/client', _client);
    },
    createCompany(_company : Company) {
        return apiInstance.post('/company', _company);
    },
    createAgent(_agent : Agent) {
        return apiInstance.post('/agent', _agent);
    },

    // #endregion
    // ----------
    
    // #region Put Requests

    updateClient(_client : Client) {
        return apiInstance.put('/client/' + _client.ID, _client);
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

    deleteClient(id: number) {
        return apiInstance.delete(`/client/${id}`);
    },
    deleteCompany(id: number) {
        return apiInstance.delete(`/company/${id}`);
    },
    deleteAgent(id: number) {
        return apiInstance.delete(`/agent/${id}`);
    },

    // #endregion
    
}