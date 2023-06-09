<template>
  <div>
    <h1>Login</h1>
    <div class="flex flex-column gap-2">
      <label for="email">Email</label>
      <InputText id="email" v-model="email"/>
      
      <label for="password">Password</label>
      <Password v-model="password" :feedback="false" />

      <Button type="submit" @click="login">Login</Button>
    </div>
  </div>
</template>

<script setup lang="ts">
  definePageMeta({
      layout: "landing-layout",
  });
</script>

<script lang="ts">
import LandingLayout from '~/layouts/landingLayout.vue';
import crmAPI from '~~/service/crmAPI';
import Cookies from 'universal-cookie';

export default defineComponent({
  name: 'LoginPage',
  layout: LandingLayout,
  data() {
    return {
      email: '' as string,
      password: '' as string,
    };
  },
  methods: {
    login() {

     
      crmAPI.login(this.email, this.password)
        .then((response : any) => {

          const authCookie = response.data.token;
        
          const cookies = new Cookies();
          const current = new Date();
          const nextYear = new Date();

          nextYear.setFullYear(current.getFullYear() + 1);

          cookies.set('authHandshakeCRM', authCookie, {
              path: '/',
              expires: nextYear,
          });

      })
      .catch((error : any) => {
        console.log(error);
      });
      
    },
  },
});
</script>

<style scoped>
.login-page {
  max-width: 400px;
  margin: auto;
  padding: 20px;
}

form {
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 8px;
}

input[type="email"],
input[type="password"] {
  padding: 8px;
  margin-bottom: 16px;
}

button[type="submit"] {
  padding: 8px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  cursor: pointer;
}
</style>