<template>
    
    <Header />
    <HeroSection />
    <Features />
    <FAQ />
    <ContactUs />
    <Footer />

</template>

<script lang="ts">

definePageMeta({
  layout: 'landing-layout'
});

import { defineComponent } from 'vue'
import Header from '~/components/landingPage/Header.vue'
import HeroSection from '~/components/landingPage/Hero.vue'
import Features from '~/components/landingPage/Features.vue'
import ContactUs from '~/components/landingPage/ContactUs.vue'
import Footer from '~/components/landingPage/Footer.vue'
import FAQ from '~/components/landingPage/FAQ.vue'

export default defineComponent({
  name: 'Index',
  components: {
    Header,
    HeroSection,
    Features,
    ContactUs,
    Footer,
    FAQ
  }
})

</script>