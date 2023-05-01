
<template>
    <div class="card">

        <Toolbar class="mb-4">
            <template #start>
                <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openUpsertCustomerDialog" />
                <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"  @click="confirmDeleteSelected" :disabled="!selectedClients || !selectedClients.length"/>
                <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV()"  />
            </template>

            <template #end>
                <div class="flex flex-row-reverse justify-content-between">
                    <span class="p-input-icon-left">
                        <i class="pi pi-search" />
                        <InputText v-model="filters['global'].value" placeholder="Global Search" />
                    </span>
                </div>
            </template>
        </Toolbar>
        
        <DataTable
            v-model:filters="filters"
            :value="clients"
            paginator
            :rows="10"
            dataKey="ID"
            filterDisplay="row"
            :loading="loading"
            :globalFilterFields="['name', 'country.name', 'representative.name']"
            ref="dataTableCustomers"
            v-model:selection="selectedClients"
            :class="`p-datatable-sm`"
        >
            <!-- Loaders -->
            <template #empty> No customers found. </template>
            <template #loading> Loading customers data. Please wait. </template>

            <!-- Columns Setup -->
            <Column selectionMode="multiple" style="width: 1rem" :exportable="false"></Column>
            <Column :exportable="false" style="min-width:8rem">
                <template #body="slotProps">
                    <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editCustomer(slotProps.data)" />
                    <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteCustomer(slotProps.data)" />
                </template>
            </Column>
            <!---->
            <Column header="Name" field="Name" style="min-width: 12rem">
                <template #body="{ data }"> {{ data.Name }} </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
                </template>
            </Column>
            <!---->
            <Column header="Company" field="Company" style="min-width: 12rem">
                <template #body="{ data }"> {{ data.Company.Name }} </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
                </template>
            </Column>
            <!---->
            <Column header="Currency" field="Currency" style="min-width: 12rem">
                <template #body="{ data }"> {{ data.Company.Currency.Name }} </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
                </template>
            </Column>
            <!---->
            <Column header="Assigned Agent" filterField="Agent" :showFilterMenu="false" :filterMenuStyle="{ width: '14rem' }" style="min-width: 14rem">
                <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img :alt="data.assignedAgentId" :src="data.AssignedAgent.PhotoUrl" style="width: 32px" />
                        <span>{{ data.AssignedAgent.Name }}</span>
                    </div>
                </template>
                <template #filter="{ filterModel, filterCallback }">
                    <MultiSelect v-model="filterModel.value" @change="filterCallback()" :options="agents" optionLabel="name" placeholder="Any" class="p-column-filter" style="min-width: 14rem" :maxSelectedLabels="1">
                        <template #option="slotProps">
                            <div class="flex align-items-center gap-2">
                                <img :alt="slotProps.option.name" :src="`https://www.meme-arsenal.com/memes/328188a0502635c0b61a76d65dbd51d5.jpg`" style="width: 32px" />
                                <span>{{ slotProps.option.name }}</span>
                            </div>
                        </template>
                    </MultiSelect>
                </template>
            </Column>

        </DataTable>

    </div>

    <!-- Upsert Customer Dialog -->
    <Dialog
        v-model:visible="upsertCustomerDialog"
        :style="{width: '450px'}"
        header="Customer Details"
        :modal="true"
        class="p-fluid"
    >   
        <!-- Name -->
        <div class="field">
            <label for="name">Name</label>
            <InputText id="name" v-model.trim="editedClient.Name" required="true" autofocus :class="{'p-invalid': submitted && !editedClient.Name}" />
            <small class="p-error" v-if="submitted && !editedClient.Name">Name is required.</small>
        </div>

        <!-- Email -->
        <div class="field">
        <label for="email">Email</label>
        <InputText
            id="email"
            v-model.trim="editedClient.Email"
           
        />
        <small class="p-error" v-if="submitted">
            Please enter a valid email address.
        </small>
        </div>

        <!-- PhoneNumber -->
        <div class="field">
        <label for="phone-number">Phone Number</label>
        <InputText
            id="phone-number"
            v-model.trim="editedClient.PhoneNumber"
            :class="{ 'p-invalid': submitted }"
        />
        <small
            class="p-error"
            v-if="submitted"
        >
            Please enter a valid phone number.
        </small>
        </div>

        <!-- Aniversary -->
        <div class="field">
        <label for="aniversary">Aniversary</label>
        <Calendar
            id="aniversary"
            v-model="editedClient.Aniversary"
            :showIcon="true"
            :dateFormat="'mm/dd/yy'"
        />
        </div>

        <!-- Company -->
        <div class="field">
        <label for="company">Company</label>
        <Dropdown
            v-model="selectedClientCompany"
            :options="companies"
            :item-value="'ID'"
            optionLabel="Name"
            placeholder="Select a company"
        />
        <small
           v-if="submitted && !editedClient.CompanyId"
           class="p-error"
        >
            Company is required.
        </small>
        </div>

        <!-- Assigned Agent -->
        <div class="field">
        <label for="company">Agent</label>
        <Dropdown
            v-model="selectedAssignedAgent"
            :options="agents"
            :item-value="'ID'"
            optionLabel="Name"
            placeholder="Select an Agent to assign"
        />
        <small
           v-if="submitted && !editedClient.AssignedAgentId"
           class="p-error"
        >
            Assigned Agent is required.
        </small>
        </div>

        <!-- Client Type -->
        <div class="field">
        <label for="company">Type</label>
        <Dropdown
            v-model="selectedClientType"
            :options="clientTypes"
            :item-value="'ID'"
            optionLabel="Description"
            placeholder="Select the Type of the Client"
        />
        <small
           v-if="submitted && !editedClient.TypeId"
           class="p-error"
        >
            Client Type is required.
        </small>
        </div>

        <!-- Client Position -->
        <div class="field">
        <label for="company">Position</label>
        <Dropdown
            v-model="selectedClientPosition"
            :options="clientPositions"
            :item-value="'ID'"
            optionLabel="Description"
            placeholder="Select the Position of the Client"
        />
        <small
           v-if="submitted && !editedClient.FunctionId"
           class="p-error"
        >
            Client Position is required.
        </small>
        </div>


        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="closeUpsertCustomerDialog"/>
            <Button label="Save" icon="pi pi-check" text @click="saveCustomer" />
        </template>
    </Dialog>

    <!-- Delete Customer Dialog -->
    <Dialog v-model:visible="deleteCustomerDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedClient">Are you sure you want to delete <b>{{editedClient.Name}}</b>?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteCustomerDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteCustomer" />
        </template>
    </Dialog>

    <!-- Delete Various Customers Dialog -->
    <Dialog v-model:visible="deleteselectedClientsDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedClient">Are you sure you want to delete the selected products?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteselectedClientsDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteselectedClients" />
        </template>
    </Dialog>

</template>

<script lang="ts">
import { FilterMatchMode } from 'primevue/api';
import Client from '~~/types/Client';
import Agent from '~~/types/Agent';
import crmAPI from '@/service/crmAPI';
import Company from '~~/types/Company';
import EnumType from '~~/types/EnumType';

export default {
    name: 'DashboardClients',
    data() {
        return {
            editedClient: new Client(),
            clients: new Array<Client>(),
            selectedClients: new Array<Client>(),
            loading: true as boolean,
            filters: {
                global: { value: null, matchMode: FilterMatchMode.CONTAINS },
                Name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Company: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Currency: { value: null, matchMode: FilterMatchMode.IN },
                Agent: { value: null, matchMode: FilterMatchMode.IN }
            },
            agents : new Array<Agent>(),
            companies : new Array<Company>(),
            clientPositions : new Array<EnumType>(),
            clientTypes: new Array<EnumType>(),
            statuses: ['unqualified', 'qualified', 'new', 'negotiation', 'renewal', 'proposal'] as Array<string>,
            submitted: false as boolean,
            upsertCustomerDialog: false as boolean,
            deleteCustomerDialog: false as boolean,
            deleteselectedClientsDialog: false as boolean,
        };
    },
    methods: {
        exportCSV() : void {
            (this.$refs.dataTableCustomers as any).exportCSV();
        },
        openUpsertCustomerDialog() : void {
            this.editedClient = new Client();
            this.submitted = false;
            this.upsertCustomerDialog = true;
        },
        closeUpsertCustomerDialog() : void {
            this.upsertCustomerDialog = false;
            this.submitted = false;
        },
        saveCustomer() : void {
            this.submitted = true;

			if (this.editedClient.Name.trim()) {
                if (this.editedClient.ID) {
                    this.clients[this.findIndexById(this.editedClient.ID)] = this.editedClient;
                    this.$toast.add({severity:'success', summary: 'Successful', detail: 'Product Updated', life: 3000});
                }
                else {
                    this.editedClient.ID = this.createId();
                    this.clients.push(this.editedClient);
                    this.$toast.add({severity:'success', summary: 'Successful', detail: 'Product Created', life: 3000});
                }
                this.upsertCustomerDialog = false;
                this.editedClient = new Client();
            }
        },
        editCustomer(_customer : Client) : void {
            Object.assign(this.editedClient, _customer);
            this.upsertCustomerDialog = true;
        },
        confirmDeleteCustomer(_customer : Client) : void {
            Object.assign(this.editedClient, _customer);
            this.deleteCustomerDialog = true;
        },
        deleteCustomer() : void {
            this.clients = this.clients.filter(val => val.ID !== this.editedClient.ID);
            this.deleteCustomerDialog = false;
            this.editedClient = new Client();
            this.$toast.add({severity:'success', summary: 'Successful', detail: 'Product Deleted', life: 3000});
        },
        confirmDeleteSelected() : void {
            this.deleteCustomerDialog = true;
        },
        deleteselectedClients() : void {
            this.clients = this.clients.filter(val => !this.selectedClients.includes(val));
            this.deleteCustomerDialog = false;
            this.selectedClients = new Array<Client>();
            this.$toast.add({severity:'success', summary: 'Successful', detail: 'Customers Deleted', life: 3000});
        },
        findIndexById(id : number) : number {
            let index = -1;
            for (let i = 0; i < this.clients.length; i++) {
                if (this.clients[i].ID === id) {
                    index = i;
                    break;
                }
            }
            return index;
        },
        createId() : number {
            // generate a and return a random int number
            let id = Math.floor(Math.random() * 10000);
            // check if the id already exists
            for (let i = 0; i < this.clients.length; i++) {
                let client = this.clients[i];
                if (client.ID === id) {
                    id = this.createId();
                    break;
                }
            }
            return id;
        },

        // #region API Calls

        loadClients() : void {
            crmAPI.getClients()
                .then((response : any) => {
                    this.clients = response.data.clients;
                    this.loading = false;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadAgents() : void {
            crmAPI.getAgents()
                .then((response : any) => {
                    this.agents = response.data.agents;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadCompanies() : void {
            crmAPI.getCompanies()
                .then((response : any) => {
                    this.companies = response.data.companies;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadClientPositions() : void {
            crmAPI.getEnumTypes("Position")
                .then((response : any) => {
                   this.clientPositions = response.data.enumTypes;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadClientTypes() : void {
            crmAPI.getEnumTypes("ClientType")
                .then((response : any) => {
                    this.clientTypes = response.data.enumTypes;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },

        // #endregion
    },
    computed: {
        selectedClientCompany: {
            get() {
                if (this.editedClient) {
                    return this.companies.find(company => company.ID === this.editedClient.CompanyId);
                } else {
                    return null;
                }
            },
            set(value : any) {
                this.editedClient.CompanyId = value.ID;
                this.editedClient.Company = value;
            }
        },
        selectedAssignedAgent: {
            get() {
                if (this.editedClient) {
                    return this.agents.find(agent => agent.ID === this.editedClient.AssignedAgentId);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedClient.AssignedAgentId = value.ID;
                this.editedClient.AssignedAgent = value;
            }
        },
        selectedClientType: {
            get() {
                if (this.editedClient) {
                    return this.clientTypes.find(clientType => clientType.ID === this.editedClient.TypeId);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedClient.TypeId = value.ID;
                this.editedClient.Type = value;
            }
        },
        selectedClientPosition: {
            get() {
                if (this.editedClient) {
                    return this.clientPositions.find(clientPosition => clientPosition.ID === this.editedClient.FunctionId);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedClient.FunctionId = value.ID;
                this.editedClient.Position = value;
            }
        }
    },


    /* Lifecycle Methods */
    mounted() {
        this.loadClients();
        this.loadCompanies();
        this.loadAgents();
        
        this.loadClientTypes();
        this.loadClientPositions();
    },

};
</script>