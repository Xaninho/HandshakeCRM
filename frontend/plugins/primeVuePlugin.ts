import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import Chart from 'primevue/chart';

export default defineNuxtPlugin(nuxtApp => {
    nuxtApp.vueApp.use(PrimeVue, {ripple: true});
    nuxtApp.vueApp.component('Button', Button);
    nuxtApp.vueApp.component('Chart', Chart);
})