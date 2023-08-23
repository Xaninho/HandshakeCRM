<template>
    <section class="bg-partner relative max-w-full sm:mx-6 my-24 shadow sm:rounded-2xl overflow-hidden">
      <div class="w-full px-6 sm:px-0 py-16 flex flex-col items-center justify-center space-y-4 text-center">
        <h3 data-aos="flip-down" class="text-2xl text-neutral-800 font-semibold">Trusted Technologies</h3>
        <p data-aos="flip-down" class="paragraph">Crafted Utilizing Leading-Edge Technologies, Programming Languages, and Frameworks</p>
        <div data-aos="fade-up" class="flex flex-wrap items-center justify-center">
          <LandingPageTechImage
            v-for="img in ['nuxt.png', 'typescript.png', 'netlify.png', 'go.png', 'gin.png']"
            :key="img"
            :img="img"
          />
        </div>
      </div>
    </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'TrustedTechnologies'
})
</script>