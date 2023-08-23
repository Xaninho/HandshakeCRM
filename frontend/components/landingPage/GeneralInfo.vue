<template>
    <section class="bg-trading-tools relative max-w-full sm:mx-4 my-20 py-16 shadow rounded-2xl overflow-hidden">
      <div class="relative max-w-screen-xl px-4 sm:px-2 mx-auto grid grid-cols-12 gap-x-6">
        <LandingPageTradingToolImage data-aos="fade-left" class="hidden sm:block" />
        <div data-aos="fade-right" class="col-span-12 lg:col-span-6 space-y-8 sm:space-y-6 px-4 sm:px-6 mt-8">
          <h2 class="text-4xl font-semibold">Next-Level Business Management Solutions</h2>
          <div class="space-y-2">
            <h4 class="text-lg font-medium">Tailored Access, Seamless Functionality</h4>
            <p class="paragraph text-sm xl:text-base">
              Experience the finest in CRM capabilities, catering to businesses and enterprises with a commitment to exceptional service delivery, dependable task execution, and uninterrupted operational performance.
            </p>
          </div>
          <div class="space-y-2">
            <h4 class="text-lg font-medium">Partnered Performance Enhancement</h4>
            <p class="paragraph text-sm xl:text-base">
              Boost your workflow efficiency by seamlessly integrating our CRM with your existing systems. Our dynamic integration solutions are designed to optimize your processes, ensuring a smooth and effective operational environment.
            </p>
          </div>
          <div class="space-y-2">
            <h4 class="text-lg font-medium">Customer-Centric Assistance</h4>
            <p class="paragraph text-sm xl:text-base">
              Enjoy priority 24/7 support accessible globally via phone or email. Assigned account managers dedicated to our valued partners.
            </p>
          </div>
        </div>
      </div>
    </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'GeneralInfo'
})
</script>