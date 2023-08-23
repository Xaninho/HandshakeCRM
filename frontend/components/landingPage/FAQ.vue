<template>
   <section class="w-full my-24">
      <BaseSection>
        <div data-aos="fade-left" data-aos-delay="150" class="col-span-12 lg:col-span-6 px-4 sm:px-6 mt-8">
          <h2 class="text-3xl sm:text-4xl font-semibold mb-10 sm:mb-6">Frequently asked questions</h2>

          <ul class="shadow-box">
            <BaseAccordion v-for="(accordion, index) in accordions" :key="index" :accordion="accordion" />
          </ul>
        </div>
        <div data-aos="fade-right" data-aos-delay="150" class="col-span-12 lg:col-span-6">
          <div class="w-full" style="padding-top: 100px">
            <img src="assets/images/faq.png" class="w-full" alt="" />
          </div>
        </div>
      </BaseSection>
    </section>
  </template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'FAQ',
  data() {
    return {
      accordions: [
        {
          title: 'How does Handshake benefit my business?',
          description:
            "Handshake CRM transforms your business operations by centralizing customer data, automating tasks, and providing insights to make informed decisions. It helps you nurture leads, improve customer satisfaction, and drive growth.",
        },
        {
          title: 'How secure is the data stored in Handshake CRM',
          description:
            "At Handshake, security is paramount. We employ robust encryption protocols and advanced security measures to safeguard your data. Our commitment to data protection ensures your information remains confidential and secure.",
        },
        {
          title: 'How is customer support handled with Handshake CRM',
          description:
            "We take pride in providing excellent customer support. Our dedicated team is available to assist you through various channels, ensuring you receive timely assistance whenever you need it.",
        }
      ]
    }
  }
})
</script>
