<template>
  <div class="surface-0 text-center">
    <div class="mb-3 font-bold text-4xl">
      <span class="text-green-600">PrimeVue, </span>
      <span class="text-blue-400">TypeScript</span>
      <span class="text-green-400"> & Nuxt 3</span>
    </div>
    <h2 class="text-2xl pb-6">
      HandShake CRM in early development!
      <MyIcon/>
      
    </h2>
    <div class="grid grid-cols-2 lg:grid-cols-3 gap-2">
    </div>
  </div>
</template>

<script setup lang='ts'>

/// <reference types="vite-svg-loader" />
import MyIcon from '@/assets/NuxtIllustration.svg?component'

</script>

<style>

div > svg {
  height: 50px;
  width: 50px;
}

</style>