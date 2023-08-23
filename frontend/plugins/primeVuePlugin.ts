import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import Chart from 'primevue/chart';
import Card from 'primevue/card';
import ToastService from 'primevue/toastservice';

export default defineNuxtPlugin(nuxtApp => {
    nuxtApp.vueApp.use(PrimeVue, {ripple: true});
    nuxtApp.vueApp.component('Button', Button);
    nuxtApp.vueApp.component('Chart', Chart);
    nuxtApp.vueApp.component('Card', Card);
    nuxtApp.vueApp.use(ToastService);
})