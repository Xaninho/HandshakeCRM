<template>
    <section class="bg-bookmark-white py-20">
      <div class="container">
        <!-- Heading -->
        <div class="sm:w-3/4 lg:w-5/12 mx-auto px-2">
          <h1 class="text-3xl text-center text-bookmark-blue">Frequently Asked Questions</h1>
          <p class="text-center text-bookmark-grey mt-4">
            Here are some of our FAQs. If you have any other questions you’d like answered please feel free to email us.
          </p>
        </div>
        <!-- FAQ Items -->
        <div class="flex flex-col sm:w-3/4 lg:w-5/12 mt-12 mx-auto">
          <div class="flex items-center border-b py-4">
            <span class="flex-1">What is a Bookmark?</span>
            <i class="text-bookmark-purple fas fa-chevron-down"></i>
          </div>
          <div class="flex items-center border-b py-4">
            <span class="flex-1">How can I request a new browser?</span>
            <i class="text-bookmark-purple fas fa-chevron-down"></i>
          </div>
          <div class="flex items-center border-b py-4">
            <span class="flex-1">Is there a mobile app?</span>
            <i class="text-bookmark-purple fas fa-chevron-down"></i>
          </div>
          <div class="flex items-center border-b py-4">
            <span class="flex-1">What about other Chromium browsers?</span>
            <i class="text-bookmark-purple fas fa-chevron-down"></i>
          </div>
          <button type="button" class="mt-12 flex self-center btn btn-purple hover:bg-bookmark-white hover:text-black">
            More Info
          </button>
        </div>
      </div>
    </section>
  </template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'FAQ'
})
</script>
