
<template>
  <div class="card">

      <Toolbar class="mb-4">
          <template #start>
              <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openUpsertCompanyDialog" />
              <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"  @click="confirmDeleteSelectedCompanies" :disabled="!selectedCompanies || !selectedCompanies.length"/>
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
          :value="companies"
          paginator
          :rows="10"
          dataKey="ID"
          filterDisplay="row"
          :loading="loading"
          :globalFilterFields="['name', 'country.name']"
          ref="dataTableCompanies"
          v-model:selection="selectedCompanies"
          :class="`p-datatable-sm`"
      >
          <!-- Loaders -->
          <template #empty> No Companies found. </template>
          <template #loading> Loading Companies data. Please wait. </template>

          <!-- Columns Setup -->
          <Column selectionMode="multiple" style="width: 1rem" :exportable="false"></Column>
          <Column :exportable="false" style="min-width:8rem">
              <template #body="slotProps">
                  <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editCompany(slotProps.data)" />
                  <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteCompany(slotProps.data)" />
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
          <Column header="Country" field="Country" style="min-width: 12rem">
              <template #body="{ data }"> {{ data.Country.Name }} </template>
              <template #filter="{ filterModel, filterCallback }">
                  <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
              </template>
          </Column>

      </DataTable>

  </div>

  <!-- Upsert Company Dialog -->
  <Dialog
      v-model:visible="upsertCompanyDialog"
      :style="{width: '450px'}"
      header="Company Details"
      :modal="true"
      class="p-fluid"
  >   
        <!-- NIF -->
        <div class="field">
            <label for="nif">NIF</label>
            <InputNumber id="nif" v-model.trim="editedCompany.NIF"  :useGrouping="false" required="true" autofocus :class="{'p-invalid': submitted && !editedCompany.NIF}" />
            <small class="p-error" v-if="submitted && !editedCompany.NIF">
                NIF is required.
            </small>
        </div>
      
      <!-- Name -->
      <div class="field">
          <label for="name">Name</label>
          <InputText id="name" v-model.trim="editedCompany.Name" required="true" autofocus :class="{'p-invalid': submitted && !editedCompany.Name}" />
          <small class="p-error" v-if="submitted && !editedCompany.Name">
              Name is required.
          </small>
      </div>
      
      <!-- Country -->
      <div class="field">
        <label for="country">Country</label>
        <Dropdown
            v-model="selectedCompanyCountry"
            :options="countries"
            :item-value="'ID'"
            optionLabel="Name"
            placeholder="Select an Country to assign"
        />
        <small v-if="submitted && !editedCompany.CountryId" class="p-error">
            Country is required.
        </small>
      </div>

      <!-- Address -->
      <div class="field">
      <label for="address">Address</label>
      <InputText
          id="address"
          v-model.trim="editedCompany.Address"
          :class="{ 'p-invalid': submitted && !editedCompany.Address }"
      />
      <small class="p-error" v-if="submitted && !editedCompany.Address">
          Please enter a valid Address
      </small>
      </div>

      <!-- Postal Code -->
      <div class="field">
      <label for="postalCode">Postal Code</label>
      <InputText
          id="postalCode"
          v-model="editedCompany.PostalCode"
          :class="{ 'p-invalid': submitted && !editedCompany.PostalCode }"
      />
      <small class="p-error" v-if="submitted && !editedCompany.PostalCode">
          Please enter a valid phone number.
      </small>
      </div>

      <!-- Currency -->
      <div class="field">
      <label for="currency">Currency</label>
      <Dropdown
          v-model="selectedCompanyCurrency"
          :options="currencies"
          :item-value="'ID'"
          optionLabel="Code"
          placeholder="Select the Currency used by the Company"
      />
      <small v-if="submitted && !editedCompany.CurrencyId" class="p-error">
          Currency is required.
      </small>
      </div>

      <!-- Electronic Billing -->
      <div class="field">
      <label for="eletronicBilling">Electronic Billing</label>
      <ToggleButton v-model="editedCompany.HasElectronicBilling" />
      </div>

      <!-- Notes -->
      <div class="field">
      <label for="notes">Notes</label>
      <Textarea
          v-model="editedCompany.Notes"
          rows="5"
          cols="30"
      />
      </div>

      <template #footer>
          <Button label="Cancel" icon="pi pi-times" text @click="closeupsertCompanyDialog"/>
          <Button label="Save" icon="pi pi-check" text @click="upsertCompany" />
      </template>
  </Dialog>

  <!-- Delete Company Dialog -->
  <Dialog v-model:visible="deleteCompanyDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
      <div class="confirmation-content">
          <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
          <span v-if="editedCompany">Are you sure you want to delete <b>{{editedCompany.Name}}</b>?</span>
      </div>
      <template #footer>
          <Button label="No" icon="pi pi-times" text @click="deleteCompanyDialog = false"/>
          <Button label="Yes" icon="pi pi-check" text @click="deleteCompany" />
      </template>
  </Dialog>

  <!-- Delete Various Companies Dialog -->
  <Dialog v-model:visible="deleteSelectedCompaniesDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
      <div class="confirmation-content">
          <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
          <span v-if="editedCompany">Are you sure you want to delete the selected clients?</span>
      </div>
      <template #footer>
          <Button label="No" icon="pi pi-times" text @click="deleteSelectedCompaniesDialog = false"/>
          <Button label="Yes" icon="pi pi-check" text @click="deleteSelectedCompanies" />
      </template>
  </Dialog>

</template>

<script lang="ts">
import { FilterMatchMode } from 'primevue/api';
import crmAPI from '@/service/crmAPI';

import Client from '~~/types/Contact';
import Company from '~~/types/Company';
import Currency from '~~/types/Currency';
import EnumType from '~~/types/EnumType';
import Country from '~~/types/Country';

export default {
  name: 'DashboardCompanies',
  data() {
      return {
          editedCompany: new Company(),
          companies: new Array<Company>(),
          countries: new Array<Country>(),
          selectedCompanies: new Array<Company>(),
          currencies: new Array<Currency>(),
          loading: true as boolean,
          filters: {
              global: { value: null, matchMode: FilterMatchMode.CONTAINS },
              Name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
              Country: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
          },

          submitted: false as boolean,
          upsertCompanyDialog: false as boolean,
          deleteCompanyDialog: false as boolean,
          deleteSelectedCompaniesDialog: false as boolean,
      };
  },
  methods: {
      exportCSV() : void {
          (this.$refs.dataTableCompanies as any).exportCSV();
      },
      openUpsertCompanyDialog() : void {
          this.editedCompany = new Company();
          this.submitted = false;
          this.upsertCompanyDialog = true;
      },
      closeupsertCompanyDialog() : void {
          this.upsertCompanyDialog = false;
          this.submitted = false;
      },
      upsertCompany() : void {
          this.submitted = true;

          if (this.editedCompany.Name.trim()) {
              if (this.editedCompany.ID > 0) {

                  crmAPI.updateCompany(this.editedCompany)
                      .then((response : any) => {
                          console.log(response);
                          this.loadCompanies();
                          this.$toast.add({
                              severity:'success',
                              summary: 'Successful',
                              detail: 'Company Updated',
                              life: 3000
                          });
                      })
                      .catch((error : any) => {
                          console.log(error);
                      });
              
              }
              else {

                  crmAPI.createCompany(this.editedCompany)
                      .then((response : any) => {
                          console.log(response);
                          this.loadCompanies();
                          this.$toast.add({
                              severity:'success',
                              summary: 'Successful',
                              detail: 'Company Created',
                              life: 3000
                          });
                      })
                      .catch((error : any) => {
                          console.log(error);
                      });
                
              }
              
              this.upsertCompanyDialog = false;
              this.editedCompany = new Company();
          }
      },
      editCompany(_company : Company) : void {
          console.log('_company', _company);
          Object.assign(this.editedCompany, _company);
          console.log('this.editedCompany', this.editedCompany);
          this.upsertCompanyDialog = true;
      },
      confirmDeleteCompany(_company : Company) : void {
          Object.assign(this.editedCompany, _company);
          this.deleteCompanyDialog = true;
      },
      deleteCompany() : void {

          crmAPI.deleteCompany(this.editedCompany.ID)
              .then((response : any) => {
                  console.log(response);
                  
                  this.companies = this.companies.filter(val => val.ID !== this.editedCompany.ID);
                  this.deleteCompanyDialog = false;
                  this.editedCompany = new Company();

                  this.$toast.add({
                      severity:'success',
                      summary: 'Successful',
                      detail: 'Company Deleted',
                      life: 3000
                  });
              })
              .catch((error : any) => {
                  console.log(error);
              });

      },
      confirmDeleteSelectedCompanies() : void {
          this.deleteSelectedCompaniesDialog = true;
      },
      deleteSelectedCompanies() : void {

          Promise.all(this.selectedCompanies.map((company: Company) => {
              return crmAPI.deleteCompany(company.ID)
                  .then((response: any) => {
                      console.log(response);
                  })
                  .catch((error: any) => {
                      console.log(error);
                  });
          }))
          .then(() => {
              this.loadCompanies();
              this.deleteSelectedCompaniesDialog = false;
              this.selectedCompanies = new Array<Company>();
              this.$toast.add({
                  severity:'success',
                  summary: 'Successful',
                  detail: 'Deleted the selected Companies',
                  life: 3000
              });
          });
      },

      // #region API Calls

      loadCompanies() : void {
          this.loading = true;
          crmAPI.getCompanies()
              .then((response : any) => {
                  this.companies = response.data.companies;
                  console.log(this.companies);
                  this.loading = false;
              })
              .catch((error : any) => {
                  console.log(error);
                  if (error.response && error.response.status === 401) {
                    this.$router.push('/login');
                    }

                  this.loading = false;
              });
      },
      loadCurrencies() : void {
          crmAPI.getCurrencies()
              .then((response : any) => {
                  this.currencies = response.data.currencies;
              })
              .catch((error : any) => {
                  console.log(error);
              });
      },
      loadCountries() : void {
          crmAPI.getCountries()
              .then((response : any) => {
                  this.countries = response.data.countries;
              })
              .catch((error : any) => {
                  console.log(error);
              });
      },

      // #endregion
  },
  computed: {
      selectedCompanyCountry: {
        get() {
            if (this.editedCompany) {
                return this.countries.find(country => country.ID === this.editedCompany.CountryId);
            } else {
                return null;
            }
        },
        set(value : any) {
            this.editedCompany.CountryId = value.ID;
            this.editedCompany.Country = value;
        }
      },
      selectedCompanyCurrency: {
        get() {
            if (this.editedCompany) {
                return this.currencies.find(currency => currency.ID === this.editedCompany.CurrencyId);
            } else {
                return null;
            }
        },
        set(value : any) {
            this.editedCompany.CurrencyId = value.ID;
            this.editedCompany.Currency = value;
        }
      },
  },

  /* Lifecycle Methods */
  mounted() {
      this.loadCompanies();
      this.loadCurrencies();
      this.loadCountries();
  },

};
</script>