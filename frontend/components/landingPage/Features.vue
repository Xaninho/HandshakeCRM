<template>
    <section class="bg-trading-tools relative max-w-full sm:mx-4 xl:mx-10 my-24 shadow sm:rounded-2xl overflow-hidden">
      <div class="w-full py-16 flex flex-col items-center">
        <h2 data-aos="flip-down" class="text-3xl sm:text-4xl font-semibold text-center">
          Seamless Onboarding in Minutes
        </h2>
        <div
          data-aos="fade-up"
          class="relative w-full flex flex-col lg:flex-row items-center justify-between space-y-12 lg:space-y-0 px-4 xl:px-10 mt-16 sm:mt-8"
        >
        <LandingPageStep v-for="step in steps" :key="step.title" :step="step" />
        </div>
      </div>
    </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Features',
  data() {
    return {
      steps: [
        {
          img: 'login.png',
          title: 'Quick Registration',
          description:
            "Get started with your CRM journey by registering for a free account. Whether you're accessing it through the web, iOS, or Android, our intuitive setup process will guide you through creating your personalized profile.",
        },
        {
          img: 'collaborate.png',
          title: 'Effortless Funding',
          description:
            "Select your desired funding method, be it through bank transfers or credit cards, to easily finance your CRM account. We offer hassle-free options for topping up your CRM balance, putting you in control of your financial resources.",
        },
        {
          img: 'metrics.png',
          title: 'Streamlined Asset Management',
          description:
           "Initiate transactions effortlessly as you manage your valuable assets within the CRM. Our feature empowers you to buy, track, and securely store your chosen assets. Additionally, enjoy the convenience of seamless transfers to your associates, promoting effective collaboration.",
        },
      ]
    }
  }
})
</script>
