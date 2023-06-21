
<template>
    <div class="card">

        <Toolbar class="mb-4">
            <template #start>
                <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openupsertActivityDialog" />
                <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"  @click="confirmDeleteSelectedActivities" :disabled="!selectedActivities || !selectedActivities.length"/>
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
            :value="activities"
            paginator
            :rows="10"
            dataKey="ID"
            filterDisplay="row"
            :loading="loading"
            :globalFilterFields="['title', 'date']"
            ref="dataTableActivities"
            v-model:selection="selectedActivities"
            :class="`p-datatable-sm`"
        >
            <!-- Loaders -->
            <template #empty> No activities found. </template>
            <template #loading> Loading activities data. Please wait. </template>

            <!-- Columns Setup -->
            <Column selectionMode="multiple" style="width: 1rem" :exportable="false"></Column>
            <Column :exportable="false" style="min-width:8rem">
                <template #body="slotProps">
                    <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editActivity(slotProps.data)" />
                    <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteActivity(slotProps.data)" />
                </template>
            </Column>
            <!---->
            <Column header="Title" field="Title" style="min-width: 12rem">
                <template #body="{ data }"> {{ data.Title }} </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by title" />
                </template>
            </Column>
            <!---->
            <Column header="Date" field="Date" style="min-width: 12rem">
                <template #body="{ data }"> {{ data.Date }} </template>
                <template #filter="{ filterModel, filterCallback }">
                    <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by date" />
                </template>
            </Column>

        </DataTable>

    </div>

    <!-- Upsert Activity Dialog -->
    <Dialog
        v-model:visible="upsertActivityDialog"
        :style="{width: '450px'}"
        header="Activity Details"
        :modal="true"
        class="p-fluid"
    >   
        
        <div class="field">
            <label for="name">Title</label>
            <InputText id="name" v-model.trim="editedActivity.Title" required="true" autofocus :class="{'p-invalid': submitted && !editedActivity.Title}" />
            <small class="p-error" v-if="submitted && !editedActivity.Title">
                Title is required.
            </small>
        </div>

        <div class="field">
        <label for="notes">Notes</label>
        <Textarea
            v-model="editedActivity.Notes"
            rows="5"
            cols="30"
        />
        </div>

        <div class="field">
        <label for="email">Location</label>
        <InputText
            id="location"
            v-model.trim="editedActivity.Location"
            :class="{ 'p-invalid': submitted && !editedActivity.Location }"
        />
        <small class="p-error" v-if="submitted && !editedActivity.Location">
            Please enter a valid location address.
        </small>
        </div>

        <div class="field">
        <label for="phone-number">Duration</label>
        <InputText
            id="duration"
            v-model="editedActivity.Duration"
            :class="{ 'p-invalid': submitted && !editedActivity.Duration }"
        />
        <small class="p-error" v-if="submitted && !editedActivity.Duration">
            Please enter a valid Duration.
        </small>
        </div>

        <div class="field">
        <label for="date">Date</label>
        <Calendar
            id="date"
            v-model="editedActivity.Date"
            :showIcon="true"
            :dateFormat="'yy-mm-dd'"
        />
        </div>

        <div class="field">
        <label for="contact">Contact</label>
        <Dropdown
            v-model="selectedActivityContact"
            :options="contacts"
            :item-value="'ID'"
            optionLabel="Name"
            placeholder="Select a contact"
        />
        <small
           v-if="submitted && !editedActivity.ContactID"
           class="p-error"
        >
            Contact is required.
        </small>
        </div>

        
        <div class="field">
        <label for="outcome">Outcome</label>
        <Dropdown
            v-model="selectedActivityOutcome"
            :options="activityOutcomes"
            :item-value="'ID'"
            optionLabel="Description"
            placeholder="Select the Outcome of the Activity"
        />
        <small
           v-if="submitted && !editedActivity.OutcomeId"
           class="p-error"
        >
            Activity Outcome is required.
        </small>
        </div>

        <div class="field">
        <label for="outcome">Assigned User</label>
        <Dropdown
            v-model="selectedActivityUser"
            :options="users"
            :item-value="'ID'"
            optionLabel="Name"
            placeholder="Select the User assigned to the Activity"
        />
        <small
           v-if="submitted && !editedActivity.UserID"
           class="p-error"
        >
            AssignedUser is required.
        </small>
        </div>


        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="closeupsertActivityDialog"/>
            <Button label="Save" icon="pi pi-check" text @click="upsertActivity" />
        </template>
    </Dialog>

    <!-- Delete Activity Dialog -->
    <Dialog v-model:visible="deleteActivityDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedActivity">Are you sure you want to delete <b>{{editedActivity.Title}}</b>?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteActivityDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteActivity" />
        </template>
    </Dialog>

    <!-- Delete Various Activities Dialog -->
    <Dialog v-model:visible="deleteSelectedActivitiesDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
        <div class="confirmation-content">
            <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
            <span v-if="editedActivity">Are you sure you want to delete the selected activities?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="deleteSelectedActivitiesDialog = false"/>
            <Button label="Yes" icon="pi pi-check" text @click="deleteSelectedActivities" />
        </template>
    </Dialog>

</template>

<script lang="ts">
import { FilterMatchMode } from 'primevue/api';
import crmAPI from '@/service/crmAPI';

import Contact from '~~/types/Contact';
import EnumType from '~~/types/EnumType';
import Activity from '~~/types/Activity';
import User from '~~/types/User';

export default {
    name: 'DashboardActivities',
    data() {
        return {
            editedActivity: new Activity(),
            activities: new Array<Activity>(),
            contacts : new Array<Contact>(),
            selectedActivities: new Array<Activity>(),
            loading: true as boolean,
            filters: {
                global: { value: null, matchMode: FilterMatchMode.CONTAINS },
                Title: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Date: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
            },
            users : new Array<User>(),
            activityOutcomes : new Array<EnumType>(),
            statuses: ['unqualified', 'qualified', 'new', 'negotiation', 'renewal', 'proposal'] as Array<string>,
            submitted: false as boolean,
            upsertActivityDialog: false as boolean,
            deleteActivityDialog: false as boolean,
            deleteSelectedActivitiesDialog: false as boolean,
        };
    },
    methods: {
        exportCSV() : void {
            (this.$refs.dataTableActivities as any).exportCSV();
        },
        openupsertActivityDialog() : void {
            this.editedActivity = new Activity();
            this.submitted = false;
            this.upsertActivityDialog = true;
        },
        closeupsertActivityDialog() : void {
            this.upsertActivityDialog = false;
            this.submitted = false;
        },
        upsertActivity() : void {
            this.submitted = true;

			if (this.editedActivity.Title.trim()) {
                if (this.editedActivity.ID > 0) {

                    crmAPI.updateActivity(this.editedActivity)
                        .then((response : any) => {
                            console.log(response);
                            this.loadActivities();
                            this.$toast.add({
                                severity:'success',
                                summary: 'Successful',
                                detail: 'Activity Updated',
                                life: 3000
                            });
                        })
                        .catch((error : any) => {
                            console.log(error);
                        });
                
                }
                else {

                    crmAPI.createActivity(this.editedActivity)
                        .then((response : any) => {
                            console.log(response);
                            this.loadActivities();
                            this.$toast.add({
                                severity:'success',
                                summary: 'Successful',
                                detail: 'Activity Created',
                                life: 3000
                            });
                        })
                        .catch((error : any) => {
                            console.log(error);
                        });
                  
                }
                
                this.upsertActivityDialog = false;
                this.editedActivity = new Activity();
            }
        },
        editActivity(_activity : Activity) : void {
            Object.assign(this.editedActivity, _activity);
            this.upsertActivityDialog = true;
        },
        confirmDeleteActivity(_activity : Activity) : void {
            Object.assign(this.editedActivity, _activity);
            this.deleteActivityDialog = true;
        },
        deleteActivity() : void {

            crmAPI.deleteActivity(this.editedActivity.ID)
                .then((response : any) => {
                    console.log(response);
                    
                    this.activities = this.activities.filter(val => val.ID !== this.editedActivity.ID);
                    this.deleteActivityDialog = false;
                    this.editedActivity = new Activity();

                    this.$toast.add({
                        severity:'success',
                        summary: 'Successful',
                        detail: 'Activity Deleted',
                        life: 3000
                    });
                })
                .catch((error : any) => {
                    console.log(error);
                });

        },
        confirmDeleteSelectedActivities() : void {
            this.deleteSelectedActivitiesDialog = true;
        },
        deleteSelectedActivities() : void {

            Promise.all(this.selectedActivities.map((activity: Activity) => {
                return crmAPI.deleteActivity(activity.ID)
                    .then((response: any) => {
                        console.log(response);
                    })
                    .catch((error: any) => {
                        console.log(error);
                    });
            }))
            .then(() => {
                this.loadActivities();
                this.deleteSelectedActivitiesDialog = false;
                this.selectedActivities = new Array<Activity>();
                this.$toast.add({
                    severity:'success',
                    summary: 'Successful',
                    detail: 'Deleted the selected activities',
                    life: 3000
                });
            });
        },

        // #region API Calls

        loadActivities() : void {
            this.loading = true;
            crmAPI.getActivities()
                .then((response : any) => {
                    this.activities = response.data.activities;
                    console.log(this.activities);
                    this.loading = false;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadUsers() : void {
            crmAPI.getUsers()
                .then((response : any) => {
                    this.users = response.data.users;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadContacts() : void {
            crmAPI.getContacts()
                .then((response : any) => {
                    this.contacts = response.data.contacts;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },
        loadActivityOutcomes() : void {
            crmAPI.getEnumTypes("ActivityOutcome")
                .then((response : any) => {
                   this.activityOutcomes = response.data.enumTypes;
                })
                .catch((error : any) => {
                    console.log(error);
                });
        },

        // #endregion
    },
    computed: {
        selectedActivityOutcome: {
            get() {
                if (this.editedActivity) {
                    return this.activityOutcomes.find(activityOutcome => activityOutcome.ID === this.editedActivity.OutcomeId);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedActivity.OutcomeId = value.ID;
                this.editedActivity.Outcome = value;
            }
        },
        selectedActivityContact: {
            get() {
                if (this.editedActivity) {
                    return this.contacts.find(contact => contact.ID === this.editedActivity.ContactID);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedActivity.ContactID = value.ID;
                this.editedActivity.Contact = value;
            }
        },
        selectedActivityUser: {
            get() {
                if (this.editedActivity) {
                    return this.users.find(user => user.ID === this.editedActivity.UserID);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedActivity.UserID = value.ID;
                this.editedActivity.User = value;
            }
        },
    },


    /* Lifecycle Methods */
    mounted() {
        this.loadActivities();
        this.loadUsers();
        this.loadContacts();
        this.loadActivityOutcomes();
    },

};
</script>