
<template>
    <div class="card">

        <Toolbar class="mb-4">
            <template #start>
                <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openUpsertCustomerDialog" />
                <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"  @click="confirmDeleteSelected" :disabled="!selectedCustomers || !selectedCustomers.length"/>
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
            :value="customers"
            paginator
            :rows="10"
            dataKey="id"
            filterDisplay="row"
            :loading="loading"
            :globalFilterFields="['name', 'country.name', 'representative.name', 'status']"
            ref="dataTableCustomers"
            v-model:selection="selectedCustomers"
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
            <Column header="Name" field="name" style="min-width: 12rem">
                <template #body="{ data }"> {{ data.name }} </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
                </template>
            </Column>
            <!---->
            <Column header="Country" filterField="country.name" style="min-width: 12rem">
                <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img alt="flag" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" :class="`flag flag-${data.country.code}`" style="width: 24px" />
                        <span>{{ data.country.name }}</span>
                    </div>
                </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by country" />
                </template>
            </Column>
            <!---->
            <Column header="Agent" filterField="representative" :showFilterMenu="false" :filterMenuStyle="{ width: '14rem' }" style="min-width: 14rem">
                <template #body="{ data }">
                    <div class="flex align-items-center gap-2">
                        <img :alt="data.representative.name" :src="`https://primefaces.org/cdn/primevue/images/avatar/${data.representative.image}`" style="width: 32px" />
                        <span>{{ data.representative.name }}</span>
                    </div>
                </template>
                <template #filter="{ filterModel, filterCallback }">
                    <MultiSelect v-model="filterModel.value" @change="filterCallback()" :options="representatives" optionLabel="name" placeholder="Any" class="p-column-filter" style="min-width: 14rem" :maxSelectedLabels="1">
                        <template #option="slotProps">
                            <div class="flex align-items-center gap-2">
                                <img :alt="slotProps.option.name" :src="`https://primefaces.org/cdn/primevue/images/avatar/${slotProps.option.image}`" style="width: 32px" />
                                <span>{{ slotProps.option.name }}</span>
                            </div>
                        </template>
                    </MultiSelect>
                </template>
            </Column>
            <!---->
            <Column field="status" header="Status" :showFilterMenu="false" :filterMenuStyle="{ width: '14rem' }" style="min-width: 12rem">
                <template #body="{ data }">
                    <Tag :value="data.status" :severity="getSeverity(data.status)" />
                </template>
                <template #filter="{ filterModel, filterCallback }">
                    <Dropdown v-model="filterModel.value" @change="filterCallback()" :options="statuses" placeholder="Select One" class="p-column-filter" style="min-width: 12rem" :showClear="true">
                        <template #option="slotProps">
                            <Tag :value="slotProps.option" :severity="getSeverity(slotProps.option)" />
                        </template>
                    </Dropdown>
                </template>
            </Column>
            <!---->
            <Column field="verified" header="Verified" dataType="boolean" style="min-width: 6rem">
                <template #body="{ data }">
                    <i class="pi" :class="{ 'pi-check-circle text-green-500': data.verified, 'pi-times-circle text-red-400': !data.verified }"></i>
                </template>
                <template #filter="{ filterModel, filterCallback }">
                    <TriStateCheckbox v-model="filterModel.value" @change="filterCallback()" />
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
            <InputText id="name" v-model.trim="editedCustomer.name" required="true" autofocus :class="{'p-invalid': submitted && !editedCustomer.name}" />
            <small class="p-error" v-if="submitted && !editedCustomer.name">Name is required.</small>
        </div>
        <!-- Company -->
        <div class="field">
            <label for="company">Company</label>
            <Textarea id="company" v-model="editedCustomer.company" required="true" rows="3" cols="20" />
        </div>
        <!-- Status -->
        <div class="field">
            <label for="status" class="mb-3">Client Status</label>
            <Dropdown id="status" v-model="editedCustomer.status" :options="statuses" optionLabel="label" placeholder="Select a Status">
                <template #value="slotProps">
                    <div v-if="slotProps.value && slotProps.value.value">
                        <Tag :value="slotProps.value.value" :severity="getSeverity(slotProps.value.label)" />
                    </div>
                    <div v-else-if="slotProps.value && !slotProps.value.value">
                        <Tag :value="slotProps.value" :severity="getSeverity(slotProps.value)" />
                    </div>
                    <span v-else>
                        {{slotProps.placeholder}}
                    </span>
                </template>
            </Dropdown>
        </div>
        <!-- Balance -->
        <div class="formgrid grid">
            <div class="field col">
                <label for="balance">Balance</label>
                <InputNumber id="price" v-model="editedCustomer.balance" mode="currency" currency="EUR" locale="pt-PT" />
            </div>
        </div>
        <!-- Activity -->
        <div class="formgrid grid">
            <div class="field col">
                <label for="activity">Activity</label>
                <InputNumber id="activity" v-model="editedCustomer.activity" />
            </div>
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
            <span v-if="editedCustomer">Are you sure you want to delete <b>{{editedCustomer.name}}</b>?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteCustomerDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteCustomer" />
        </template>
    </Dialog>

    <!-- Delete Various Customers Dialog -->
    <Dialog v-model:visible="deleteSelectedCustomersDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedCustomer">Are you sure you want to delete the selected products?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteSelectedCustomersDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteSelectedCustomers" />
        </template>
    </Dialog>

</template>

<script lang="ts">
import { FilterMatchMode } from 'primevue/api';
import { CustomerService } from '@/service/CustomerService';
import { Customer } from '~~/types/Customer';
import { Agent } from '~~/types/Representative';

export default {
    data() {
        return {
            editedCustomer: new Customer(),
            customers: new Array<Customer>(),
            selectedCustomers: new Array<Customer>(),
            loading: true as boolean,
            filters: {
                global: { value: null, matchMode: FilterMatchMode.CONTAINS },
                name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                'country.name': { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                representative: { value: null, matchMode: FilterMatchMode.IN },
                status: { value: null, matchMode: FilterMatchMode.EQUALS },
                verified: { value: null, matchMode: FilterMatchMode.EQUALS }
            },
            representatives: [
                { name: 'Amy Elsner', image: 'amyelsner.png' },
                { name: 'Anna Fali', image: 'annafali.png' },
                { name: 'Asiya Javayant', image: 'asiyajavayant.png' },
                { name: 'Bernardo Dominic', image: 'bernardodominic.png' },
                { name: 'Elwin Sharvill', image: 'elwinsharvill.png' },
                { name: 'Ioni Bowcher', image: 'ionibowcher.png' },
                { name: 'Ivan Magalhaes', image: 'ivanmagalhaes.png' },
                { name: 'Onyama Limba', image: 'onyamalimba.png' },
                { name: 'Stephen Shaw', image: 'stephenshaw.png' },
                { name: 'XuXue Feng', image: 'xuxuefeng.png' }
            ] as Array<Agent>,
            statuses: ['unqualified', 'qualified', 'new', 'negotiation', 'renewal', 'proposal'] as Array<string>,
            submitted: false as boolean,
            upsertCustomerDialog: false as boolean,
            deleteCustomerDialog: false as boolean,
            deleteSelectedCustomersDialog: false as boolean,
        };
    },
    methods: {
        getCustomers(data : any) {
            return [...(data || [])].map((d) => {
                d.date = new Date(d.date);

                return d;
            });
        },
        formatDate(value : any) {
            return value.toLocaleDateString('en-US', {
                day: '2-digit',
                month: '2-digit',
                year: 'numeric'
            });
        },
        formatCurrency(value : any) {
            return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
        },
        getSeverity(status : string) {
            switch (status) {
                case 'unqualified':
                    return 'danger';

                case 'qualified':
                    return 'success';

                case 'new':
                    return 'info';

                case 'negotiation':
                    return 'warning';
            }
        },
        exportCSV() : void {
            (this.$refs.dataTableCustomers as any).exportCSV();
        },
        openUpsertCustomerDialog() : void {
            this.editedCustomer = new Customer();
            this.submitted = false;
            this.upsertCustomerDialog = true;
        },
        closeUpsertCustomerDialog() : void {
            this.upsertCustomerDialog = false;
            this.submitted = false;
        },
        saveCustomer() : void {
            this.submitted = true;

			if (this.editedCustomer.name.trim()) {
                if (this.editedCustomer.id) {
                    this.customers[this.findIndexById(this.editedCustomer.id)] = this.editedCustomer;
                    this.$toast.add({severity:'success', summary: 'Successful', detail: 'Product Updated', life: 3000});
                }
                else {
                    this.editedCustomer.id = this.createId();
                    this.customers.push(this.editedCustomer);
                    this.$toast.add({severity:'success', summary: 'Successful', detail: 'Product Created', life: 3000});
                }

                this.upsertCustomerDialog = false;
                this.editedCustomer = new Customer();
            }
        },
        editCustomer(_customer : Customer) : void {
            Object.assign(this.editedCustomer, _customer);
            this.upsertCustomerDialog = true;
        },
        confirmDeleteCustomer(_customer : Customer) : void {
            Object.assign(this.editedCustomer, _customer);
            this.deleteCustomerDialog = true;
        },
        deleteCustomer() : void {
            this.customers = this.customers.filter(val => val.id !== this.editedCustomer.id);
            this.deleteCustomerDialog = false;
            this.editedCustomer = new Customer();
            this.$toast.add({severity:'success', summary: 'Successful', detail: 'Product Deleted', life: 3000});
        },
        confirmDeleteSelected() : void {
            this.deleteCustomerDialog = true;
        },
        deleteSelectedCustomers() : void {
            this.customers = this.customers.filter(val => !this.selectedCustomers.includes(val));
            this.deleteCustomerDialog = false;
            this.selectedCustomers = new Array<Customer>();
            this.$toast.add({severity:'success', summary: 'Successful', detail: 'Customers Deleted', life: 3000});
        },


        findIndexById(id : number) : number {
            let index = -1;
            for (let i = 0; i < this.customers.length; i++) {
                if (this.customers[i].id === id) {
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
            for (let i = 0; i < this.customers.length; i++) {
                let customer = this.customers[i];
                if (customer.id === id) {
                    id = this.createId();
                    break;
                }
            }
            return id;
        },
    },

    /* Lifecycle Methods */
    mounted() {
        CustomerService.getCustomersMedium().then((data) => {
            this.customers = this.getCustomers(data);
            this.loading = false;
        });
    }, 
};
</script>