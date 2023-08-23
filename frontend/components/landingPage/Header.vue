<template>
    <header>
        <nav class="container relative flex flex-wrap items-center justify-between p-8 mx-auto lg:justify-between xl:px-0">
        <div class="py-1"><img src="~/assets/images/logo-bookmark.svg" /></div>
        <ul class="hidden sm:flex flex-1 justify-end items-center gap-12 text-bookmark-blue uppercase text-xs">
            <li class="cursor-pointer">Features</li>
            <li class="cursor-pointer">Pricing</li>
            <li class="cursor-pointer">Contact</li>
            <button type="button" class="bg-bookmark-red text-white rounded-md px-7 py-3 uppercase" @click="redirectToLoginPage">Login</button>
        </ul>
        <div class="flex sm:hidden flex-1 justify-end">
            <i class="text-2xl fas fa-bars"></i>
        </div>
        </nav>
    </header>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Header',
  // create a data section with an array
  data() {
    return {
      navigation: [
        "Product",
        "Features",
        "Pricing",
        "Company",
        "Blog",
      ]
    }
  },
  methods: {
    redirectToLoginPage() : void {
      this.$router.push('/login')
    }
  }
})
</script>
