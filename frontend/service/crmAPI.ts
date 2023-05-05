import axios from 'axios';
import Client from '~~/types/Client';

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

    // #endregion
    // ----------
    
    // #region Put Requests

    updateClient(_client : Client) {
        return apiInstance.put('/client/' + _client.ID, _client);
    },

    // #endregion
    // ----------
    
    // #region Delete Requests

    deleteClient(id: number) {
        return apiInstance.delete(`/client/${id}`);
    },

    // #endregion
    
}