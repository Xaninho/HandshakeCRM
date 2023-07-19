<template>
    <section class="bg-bookmark-purple text-white py-20">
        <div class="container">
        <div class="sm:w-3/4 lg:w-2/4 mx-auto">
            <p class="font-light uppercase text-center mb-8">35,000+ ALREADY JOINED</p>
            <h1 class="text-3xl text-center">Stay up-to-date with what we’re doing</h1>
            <div class="flex flex-col sm:flex-row gap-6 mt-8">
            <input
                type="text"
                placeholder="Enter your email address"
                class="focus:outline-none flex-1 px-2 py-3 rounded-md text-black"
            />
            <button type="button" class="btn bg-bookmark-red hover:bg-bookmark-white hover:text-black">
                Contact Us
            </button>
            </div>
        </div>
        </div>
    </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'ContactUs'
})
</script>