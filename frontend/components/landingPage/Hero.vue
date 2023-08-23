<template>
  <section id="hero" class="w-full pb-24">
    <BaseSection>
      <div class="col-span-12 lg:col-span-6 mt-12 xl:mt-10 space-y-4 sm:space-y-6 px-6 text-center sm:text-left" style="padding-top: 30px;">
        <h1
          data-aos="fade-right"
          data-aos-once="true"
          class="text-[2.5rem] sm:text-5xl xl:text-6xl font-bold leading-tight capitalize sm:pr-8 xl:pr-10"
        >
          Revolutionize Business Management
        </h1>
        <p data-aos="fade-down" data-aos-once="true" data-aos-delay="300" class="paragraph hidden sm:block">
          Streamline tasks, enhance customer relationships, and supercharge growth with our intuitive CRM solution designed to elevate your business success.
        </p>
        <div
          data-aos="fade-up"
          data-aos-once="true"
          data-aos-delay="700"
          class="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4 mt-2"
        >
          <BaseButton
            class="max-w-full px-8 py-4 bg-gradient-to-r from-[#00D4B0] to-[#00E54B] border border-[#00D4B0 text-white"
            @click="navigateToDashboard"
            >
            Enter App
          </BaseButton>
        </div>
      </div>
      <div class="hidden sm:block col-span-12 lg:col-span-6">
        <div class="w-full">
          <img class="w-5/6 h-5/6 sm:w-3/4 sm:h-3/4 md:w-full md:h-full" src="~/assets/images/hero-bg.png" alt="" />
        </div>
      </div>
    </BaseSection>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'Hero',
  data() {
    return {
      router: useRouter(),
    }
  },
  methods: {
    navigateToDashboard () {
      this.router.push("/dashboard");
    },
  }
})
</script>
