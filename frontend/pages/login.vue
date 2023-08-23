<template>
  <div class="flex flex-col items-center justify-center min-h-screen">
    <img src="~assets/images/handshake.png" class="w-32 h-32 mb-4" alt="Handshake Logo" />
    <h1 class="text-center text-3xl font-bold mb-4">Login</h1>
    <div class="w-full max-w-md bg-white p-6 rounded-lg shadow-md">
      
      <label for="email" class="block text-gray-700 font-medium mb-2 w-full">Email</label>
      <InputText id="email" v-model="email" class="w-full border rounded-md py-2 px-3 focus:outline-none focus:border-green-500" />
      
      <label for="password" class="block text-gray-700 font-medium mb-2 w-full">Password</label>
      <Password v-model="password" :feedback="false" class="w-full border rounded-md py-2 px-3 focus:outline-none focus:border-green-500" />

      <Button
        type="submit"
        @click="login"
        class="w-full mt-4 bg-green-500 hover:bg-green-600 text-white font-bold py-2 rounded-full flex items-center justify-center"
      >
        Login
      </Button>
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

</style>