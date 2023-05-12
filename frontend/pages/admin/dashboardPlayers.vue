
<template>
  <div class="card">

      <Toolbar class="mb-4">
          <template #start>
              <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openUpsertAgentDialog" />
              <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"  @click="confirmDeleteSelectedAgents" :disabled="!selectedAgents || !selectedAgents.length"/>
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
          :value="agents"
          paginator
          :rows="10"
          dataKey="ID"
          filterDisplay="row"
          :loading="loading"
          :globalFilterFields="['name', 'phoneNumber', 'email', 'position.description']"
          ref="dataTableAgents"
          v-model:selection="selectedAgents"
          :class="`p-datatable-sm`"
      >
          <!-- Loaders -->
          <template #empty> No Agents found. </template>
          <template #loading> Loading Agents data. Please wait. </template>

          <!-- Columns Setup -->
          <Column selectionMode="multiple" style="width: 1rem" :exportable="false"></Column>
          <Column :exportable="false" style="min-width:8rem">
              <template #body="slotProps">
                  <Button icon="pi pi-pencil" outlined rounded class="mr-2" @click="editAgent(slotProps.data)" />
                  <Button icon="pi pi-trash" outlined rounded severity="danger" @click="confirmDeleteAgent(slotProps.data)" />
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
          <Column header="PhoneNumber" field="PhoneNumber" style="min-width: 12rem">
              <template #body="{ data }"> {{ data.PhoneNumber }} </template>
              <template #filter="{ filterModel, filterCallback }">
                  <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
              </template>
          </Column>
          <!---->
          <Column header="Email" field="Email" style="min-width: 12rem">
              <template #body="{ data }"> {{ data.Email}} </template>
              <template #filter="{ filterModel, filterCallback }">
                  <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Search by name" />
              </template>
          </Column>
          <!---->
          <Column header="Position" filterField="Position" :showFilterMenu="false" :filterMenuStyle="{ width: '14rem' }" style="min-width: 14rem">
              <template #body="{ data }">
                  <div class="flex align-items-center gap-2">
                      <span>{{ data.Position.Description }}</span>
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

  <!-- Upsert Agent Dialog -->
  <Dialog
      v-model:visible="upsertAgentDialog"
      :style="{width: '450px'}"
      header="Agent Details"
      :modal="true"
      class="p-fluid"
  >   
      <!-- Name -->
      <div class="field">
          <label for="name">Name</label>
          <InputText id="name" v-model.trim="editedAgent.Name" required="true" autofocus :class="{'p-invalid': submitted && !editedAgent.Name}" />
          <small class="p-error" v-if="submitted && !editedAgent.Name">
              Name is required.
          </small>
      </div>

      <!-- Email -->
      <div class="field">
      <label for="email">Email</label>
      <InputText
          id="email"
          v-model.trim="editedAgent.Email"
          :class="{ 'p-invalid': submitted && !editedAgent.Email }"
      />
      <small class="p-error" v-if="submitted && !editedAgent.Email">
          Please enter a valid email address.
      </small>
      </div>

      <!-- PhoneNumber -->
      <div class="field">
      <label for="phone-number">Phone Number</label>
      <InputNumber
          id="phone-number"
          v-model="editedAgent.PhoneNumber"
          :class="{ 'p-invalid': submitted && !editedAgent.PhoneNumber }"
      />
      <small class="p-error" v-if="submitted && !editedAgent.PhoneNumber">
          Please enter a valid phone number.
      </small>
      </div>

      <!-- Agent Position -->
      <div class="field">
      <label for="company">Agent Position</label>
      <Dropdown
          v-model="selectedAgentPosition"
          :options="agentPositions"
          :item-value="'ID'"
          optionLabel="Description"
          placeholder="Select the Position of the Agent"
      />
      <small
         v-if="submitted && !editedAgent.PositionId"
         class="p-error"
      >
          Agent Position is required.
      </small>
      </div>


      <template #footer>
          <Button label="Cancel" icon="pi pi-times" text @click="closeUpsertAgentDialog"/>
          <Button label="Save" icon="pi pi-check" text @click="upsertAgent" />
      </template>
  </Dialog>

  <!-- Delete Agent Dialog -->
  <Dialog v-model:visible="deleteAgentDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
      <div class="confirmation-content">
          <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
          <span v-if="editedAgent">Are you sure you want to delete <b>{{editedAgent.Name}}</b>?</span>
      </div>
      <template #footer>
          <Button label="No" icon="pi pi-times" text @click="deleteAgentDialog = false"/>
          <Button label="Yes" icon="pi pi-check" text @click="deleteAgent" />
      </template>
  </Dialog>

  <!-- Delete Various Agents Dialog -->
  <Dialog v-model:visible="deleteSelectedAgentDialog" :style="{width: '450px'}" header="Confirm" :modal="true">
      <div class="confirmation-content">
          <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
          <span v-if="editedAgent">Are you sure you want to delete the selected agents?</span>
      </div>
      <template #footer>
          <Button label="No" icon="pi pi-times" text @click="deleteSelectedAgentDialog = false"/>
          <Button label="Yes" icon="pi pi-check" text @click="deleteSelectedAgents" />
      </template>
  </Dialog>

</template>

<script lang="ts">
import { FilterMatchMode } from 'primevue/api';
import crmAPI from '@/service/crmAPI';

import Agent from '~~/types/Agent';
import EnumType from '~~/types/EnumType';

export default {
  name: 'DashboardAgents',
  data() {
      return {
          editedAgent: new Agent(),
          agents: new Array<Agent>(),
          selectedAgents: new Array<Agent>(),
          loading: true as boolean,
          filters: {
              global: { value: null, matchMode: FilterMatchMode.CONTAINS },
              Name: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
              PhoneNumber: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
              Email: { value: null, matchMode: FilterMatchMode.IN },
              Position: { value: null, matchMode: FilterMatchMode.IN }
          },
          agentPositions : new Array<EnumType>(),
          submitted: false as boolean,
          upsertAgentDialog: false as boolean,
          deleteAgentDialog: false as boolean,
          deleteSelectedAgentDialog: false as boolean,
      };
  },
  methods: {
      exportCSV() : void {
          (this.$refs.dataTableAgents as any).exportCSV();
      },
      openUpsertAgentDialog() : void {
          this.editedAgent = new Agent();
          this.submitted = false;
          this.upsertAgentDialog = true;
      },
      closeUpsertAgentDialog() : void {
          this.upsertAgentDialog = false;
          this.submitted = false;
      },
      upsertAgent() : void {
          this.submitted = true;

          if (this.editedAgent.Name.trim()) {
              if (this.editedAgent.ID > 0) {

                  crmAPI.updateAgent(this.editedAgent)
                      .then((response : any) => {
                          console.log(response);
                          this.loadAgents();
                          this.$toast.add({
                              severity:'success',
                              summary: 'Successful',
                              detail: 'Agent Updated',
                              life: 3000
                          });
                      })
                      .catch((error : any) => {
                          console.log(error);
                      });
              
              }
              else {

                  crmAPI.createAgent(this.editedAgent)
                      .then((response : any) => {
                          console.log(response);
                          this.loadAgents();
                          this.$toast.add({
                              severity:'success',
                              summary: 'Successful',
                              detail: 'Agent Created',
                              life: 3000
                          });
                      })
                      .catch((error : any) => {
                          console.log(error);
                      });
                
              }
              
              this.upsertAgentDialog = false;
              this.editedAgent = new Agent();
          }
      },
      editAgent(_agent : Agent) : void {
          Object.assign(this.editedAgent, _agent);
          this.upsertAgentDialog = true;
      },
      confirmDeleteAgent(_agent : Agent) : void {
          Object.assign(this.editedAgent, _agent);
          this.deleteAgentDialog = true;
      },
      deleteAgent() : void {

          crmAPI.deleteAgent(this.editedAgent.ID)
              .then((response : any) => {
                  console.log(response);
                  
                  this.agents = this.agents.filter(val => val.ID !== this.editedAgent.ID);
                  this.deleteAgentDialog = false;
                  this.editedAgent = new Agent();

                  this.$toast.add({
                      severity:'success',
                      summary: 'Successful',
                      detail: 'Agent Deleted',
                      life: 3000
                  });
              })
              .catch((error : any) => {
                  console.log(error);
              });

      },
      confirmDeleteSelectedAgents() : void {
          this.deleteSelectedAgentDialog = true;
      },
      deleteSelectedAgents() : void {

          Promise.all(this.selectedAgents.map((agent: Agent) => {
              return crmAPI.deleteAgent(agent.ID)
                  .then((response: any) => {
                      console.log(response);
                  })
                  .catch((error: any) => {
                      console.log(error);
                  });
          }))
          .then(() => {
              this.loadAgents();
              this.deleteSelectedAgentDialog = false;
              this.selectedAgents = new Array<Agent>();
              this.$toast.add({
                  severity:'success',
                  summary: 'Successful',
                  detail: 'Deleted the selected Agents',
                  life: 3000
              });
          });
      },

      // #region API Calls

      loadAgents() : void {
          crmAPI.getAgents()
              .then((response : any) => {
                  this.agents = response.data.agents;
                  this.loading = false;
              })
              .catch((error : any) => {
                  console.log(error);
                  this.loading = false;
              });
      },
      loadAgentPositions() : void {
          crmAPI.getEnumTypes("Position")
              .then((response : any) => {
                 this.agentPositions = response.data.enumTypes;
              })
              .catch((error : any) => {
                  console.log(error);
              });
      }

      // #endregion
  },
  computed: {
      selectedAgentPosition: {
            get() {
                if (this.editedAgent) {
                    return this.agentPositions.find(agentPosition => agentPosition.ID === this.editedAgent.PositionId);
                } else {
                    return null;
                }
            },
            set(value: any) {
                this.editedAgent.PositionId = value.ID;
                this.editedAgent.Position = value;
            }
        }
  },


  /* Lifecycle Methods */
  mounted() {
      this.loadAgents();
      this.loadAgentPositions();
  },

};
</script>