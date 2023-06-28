
<template>
    <div class="card">

        <Toolbar class="mb-4">
            <template #start>
                <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openupsertContactDialog" />
                <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"  @click="confirmDeleteSelectedContacts" :disabled="!selectedContacts || !selectedContacts.length"/>
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
            :value="contacts"
            paginator
            :rows="10"
            dataKey="ID"
            filterDisplay="row"
            :loading="loading"
            :globalFilterFields="['name', 'country.name']"
            ref="dataTableCustomers"
            v-model:selection="selectedContacts"
            :class="`p-datatable-sm`"
            v-model:expandedRows="expandedRows"
            @rowExpand="onRowExpand"
            @rowCollapse="onRowCollapse"
        >
            <!-- Loaders -->
            <template #empty> No contacts found. </template>
            <template #loading> Loading contacts data. Please wait. </template>

            <!-- Columns Setup -->
            <Column expander style="width: 5rem" />
            <Column selectionMode="multiple" style="width: 1rem" :exportable="false"></Column>
            <Column :exportable="false" style="min-width:8rem">
                <template #body="slotProps">
                    <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editContact(slotProps.data)" />
                    <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteContact(slotProps.data)" />
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

            <template #expansion="slotProps">
                <div class="p-3">
                    <h5>Orders for {{ slotProps.data.name }}</h5>
                    <DataTable :value="slotProps.data.AssociatedActivities">
                        <Column field="Title" header="Title" sortable></Column>
                        <Column field="Date" header="Date" sortable></Column>
                        <Column field="Location" header="Location" sortable></Column>
                    </DataTable>
                </div>
            </template>

        </DataTable>

    </div>

    <!-- Upsert Contac Dialog -->
    <Dialog
        v-model:visible="upsertContactDialog"
        :style="{width: '450px'}"
        header="Contact Details"
        :modal="true"
        class="p-fluid"
    >   
        <!-- Name -->
        <div class="field">
            <label for="name">Name</label>
            <InputText id="name" v-model.trim="editedContact.Name" required="true" autofocus :class="{'p-invalid': submitted && !editedContact.Name}" />
            <small class="p-error" v-if="submitted && !editedContact.Name">
                Name is required.
            </small>
        </div>

        <!-- Email -->
        <div class="field">
        <label for="email">Email</label>
        <InputText
            id="email"
            v-model.trim="editedContact.Email"
            :class="{ 'p-invalid': submitted && !editedContact.Email }"
        />
        <small class="p-error" v-if="submitted && !editedContact.Email">
            Please enter a valid email address.
        </small>
        </div>

        <!-- PhoneNumber -->
        <div class="field">
        <label for="phone-number">Phone Number</label>
        <InputNumber
            id="phone-number"
            v-model="editedContact.PhoneNumber"
            :class="{ 'p-invalid': submitted && !editedContact.PhoneNumber }"
        />
        <small class="p-error" v-if="submitted && !editedContact.PhoneNumber">
            Please enter a valid phone number.
        </small>
        </div>

        <!-- Aniversary -->
        <div class="field">
        <label for="aniversary">Aniversary</label>
        <Calendar
            id="aniversary"
            v-model="editedContact.Aniversary"
            :showIcon="true"
            :dateFormat="'yy-mm-dd'"
        />
        </div>

        <!-- Company -->
        <div class="field">
        <label for="company">Company</label>
        <Dropdown
            v-model="selectedContactCompany"
            :options="companies"
            :item-value="'ID'"
            optionLabel="Name"
            placeholder="Select a company"
        />
        <small
           v-if="submitted && !editedContact.CompanyId"
           class="p-error"
        >
            Company is required.
        </small>
        </div>

        <!-- Contacty Disposition -->
        <div class="field">
        <label for="company">Disposition</label>
        <Dropdown
            v-model="selectedContactDisposition"
            :options="contactDispositions"
            :item-value="'ID'"
            optionLabel="Description"
            placeholder="Select the Disposition of the Contatc"
        />
        <small
           v-if="submitted && !editedContact.DispositionId"
           class="p-error"
        >
            Contact Disposition is required.
        </small>
        </div>


        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="closeupsertContactDialog"/>
            <Button label="Save" icon="pi pi-check" text @click="upsertContact" />
        </template>
    </Dialog>

    <!-- Delete Customer Dialog -->
    <Dialog v-model:visible="deleteContactDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedContact">Are you sure you want to delete <b>{{editedContact.Name}}</b>?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteContactDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteContact" />
        </template>
    </Dialog>

    <!-- Delete Various Customers Dialog -->
    <Dialog v-model:visible="deleteSelectedContactsDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedContact">Are you sure you want to delete the selected contacts?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteSelectedContactsDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteselectedContacts" />
        </template>
    </Dialog>

</template>

<script lang="ts">
import { FilterMatchMode } from 'primevue/api';
import crmAPI from '@/service/crmAPI';

import Contact from '~~/types/Contact';
import Company from '~~/types/Company';
import EnumType from '~~/types/EnumType';
import HelperFunctions from '~~/service/HelperFunctions';

export default {
    name: 'Dashboardcontacts',
    data() {
        return {
            editedContact: new Contact(),
            contacts: new Array<Contact>(),
            selectedContacts: new Array<Contact>(),
            loading: true as boolean,
            expandedRows: [],
            filters: {
                global: { value: null, matchMode: FilterMatchMode.CONTAINS },
                Name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Company: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Currency: { value: null, matchMode: FilterMatchMode.IN },
                Agent: { value: null, matchMode: FilterMatchMode.IN }
            },
            companies : new Array<Company>(),
            contactDispositions : new Array<EnumType>(),
            statuses: ['unqualified', 'qualified', 'new', 'negotiation', 'renewal', 'proposal'] as Array<string>,
            submitted: false as boolean,
            upsertContactDialog: false as boolean,
            deleteContactDialog: false as boolean,
            deleteSelectedContactsDialog: false as boolean,
        };
    },
    methods: {
        exportCSV() : void {
            (this.$refs.dataTableCustomers as any).exportCSV();
        },
        openupsertContactDialog() : void {
            this.editedContact = new Contact();
            this.submitted = false;
            this.upsertContactDialog = true;
        },
        closeupsertContactDialog() : void {
            this.upsertContactDialog = false;
            this.submitted = false;
        },
        upsertContact() : void {
            this.submitted = true;

			if (this.editedContact.Name.trim()) {
                if (this.editedContact.ID > 0) {

                    crmAPI.updateContact(this.editedContact)
                        .then((response : any) => {
                            console.log(response);
                            this.loadContacts();
                            this.$toast.add({
                                severity:'success',
                                summary: 'Successful',
                                detail: 'Contact Updated',
                                life: 3000
                            });
                        })
                        .catch((error : any) => {
                            console.log(error);
                        });
                
                }
                else {

                    crmAPI.createContact(this.editedContact)
                        .then((response : any) => {
                            console.log(response);
                            this.loadContacts();
                            this.$toast.add({
                                severity:'success',
                                summary: 'Successful',
                                detail: 'Contact Created',
                                life: 3000
                            });
                        })
                        .catch((error : any) => {
                            console.log(error);
                        });
                  
                }
                
                this.upsertContactDialog = false;
                this.editedContact = new Contact();
            }
        },
        editContact(_contact : Contact) : void {
            Object.assign(this.editedContact, _contact);
            this.upsertContactDialog = true;
        },
        confirmDeleteContact(_contact : Contact) : void {
            Object.assign(this.editedContact, _contact);
            this.deleteContactDialog = true;
        },
        deleteContact() : void {

            crmAPI.deleteContact(this.editedContact.ID)
                .then((response : any) => {
                    console.log(response);
                    
                    this.contacts = this.contacts.filter(val => val.ID !== this.editedContact.ID);
                    this.deleteContactDialog = false;
                    this.editedContact = new Contact();

                    this.$toast.add({
                        severity:'success',
                        summary: 'Successful',
                        detail: 'Contact Deleted',
                        life: 3000
                    });
                })
                .catch((error : any) => {
                    console.log(error);
                });

        },
        confirmDeleteSelectedContacts() : void {
            this.deleteSelectedContactsDialog = true;
        },
        deleteSelectedContacts() : void {

            Promise.all(this.selectedContacts.map((contact: Contact) => {
                return crmAPI.deleteContact(contact.ID)
                    .then((response: any) => {
                        console.log(response);
                    })
                    .catch((error: any) => {
                        console.log(error);
                    });
            }))
            .then(() => {
                this.loadContacts();
                this.deleteSelectedContactsDialog = false;
                this.selectedContacts = new Array<Contact>();
                this.$toast.add({
                    severity:'success',
                    summary: 'Successful',
                    detail: 'Deleted the selected contacts',
                    life: 3000
                });
            });
        },
        onRowExpand(event : any) {

            this.contacts.forEach((contact : Contact) => {
                if (contact.ID == event.data.ID) {
                    contact.isLoading = true;
                }
            });
            
            crmAPI.getActivitiesByContact(event.data.ID)
                .then((response : any) => {
                    console.log(response);

                    this.contacts.forEach((contact : Contact) => {
                        if (contact.ID == event.data.ID) {
                            contact.AssociatedActivities = response.data.activities;
                            contact.isLoading = false;
                        }
                    });
                })
                .catch((error : any) => {
                    console.log(error);

                    this.contacts.forEach((contact : Contact) => {
                        if (contact.ID == event.data.ID) {
                            contact.isLoading = false;
                        }
                    });
                });


            this.$toast.add({
                    severity:'success',
                    summary: 'Successful',
                    detail: 'Row Expand',
                    life: 3000
                });
        },
        onRowCollapse(event : any) {
            console.log(event.data);
            this.$toast.add({
                    severity:'success',
                    summary: 'Successful',
                    detail: 'Row Collapse',
                    life: 3000
                });
        },
        makeAnAssociatedEventForContact(_contact : Contact) : void {
            this.$router.push({ name: 'event', params: { contactID: _contact.ID }});
        },

        // #region API Calls

        loadContacts() : void {
            this.loading = true;
            crmAPI.getContacts()
                .then((response : any) => {
                    this.contacts = response.data.contacts;
                    this.contacts.forEach((contact : Contact) => {
                        contact.Aniversary = HelperFunctions.toStringDateFormat(contact.Aniversary, false);
                    });
                    this.loading = false;
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
        loadContactDispositions() : void {
            crmAPI.getEnumTypes("Disposition")
                .then((response : any) => {
                   this.contactDispositions = response.data.enumTypes;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },

        // #endregion
    },
    computed: {
        selectedContactCompany: {
            get() {
                if (this.editedContact) {
                    return this.companies.find(company => company.ID === this.editedContact.CompanyId);
                } else {
                    return null;
                }
            },
            set(value : any) {
                this.editedContact.CompanyId = value.ID;
                this.editedContact.Company = value;
            }
        },
        selectedContactDisposition: {
            get() {
                if (this.editedContact) {
                    return this.contactDispositions.find(contactDisposition => contactDisposition.ID === this.editedContact.DispositionId);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedContact.DispositionId = value.ID;
                this.editedContact.Disposition = value;
            }
        }
    },


    /* Lifecycle Methods */
    mounted() {
        this.loadContacts();
        this.loadCompanies();
        this.loadContactDispositions();
    },

};
</script>