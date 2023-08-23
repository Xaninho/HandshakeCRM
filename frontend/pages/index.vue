<template>
  <div class="w-full">
      
      <BaseNavbar />
      <HeroSection />
      <GeneralInfo />
      <Features />
      <TrustedTechnologies />
      <FAQ />
      <Footer />

  </div>
</template>

<script lang="ts">

definePageMeta({
  layout: 'landing-layout'
});

import { defineComponent } from 'vue'
import Header from '~/components/landingPage/Header.vue'
import HeroSection from '~/components/landingPage/Hero.vue'
import Features from '~/components/landingPage/Features.vue'
import GeneralInfo from '~/components/landingPage/GeneralInfo.vue'
import FAQ from '~/components/landingPage/FAQ.vue'
import TrustedTechnologies from '~/components/landingPage/TrustedTechnologies.vue'
import Footer from '~/components/landingPage/Footer.vue'

export default defineComponent({
  name: 'Index',
  components: {
    Header,
    HeroSection,
    Features,
    GeneralInfo,
    FAQ,
    TrustedTechnologies,
    Footer
  }
})

</script>

<style>

.bg-partner {
  background: url('../assets/images/technologies/background.png');
  background-size: cover;
  background-position: center;
}

</style>