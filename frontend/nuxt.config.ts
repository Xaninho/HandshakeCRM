// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    css: [
        'primevue/resources/themes/saga-blue/theme.css',  //theme
        'primevue/resources/primevue.min.css',            //core css
        'primeicons/primeicons.css',                      //icons
    ],
    build: {
        transpile: ['primevue']
    }
})
