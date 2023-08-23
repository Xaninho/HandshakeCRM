<template>
  <nav id="navbar" class="relative z-10 w-full text-neutral-800">
    <div class="flex flex-col max-w-screen-xl px-8 mx-auto lg:items-center lg:justify-between lg:flex-row py-4">
      <div class="flex flex-col lg:flex-row items-center space-x-4 xl:space-x-8">
        <div class="w-full flex flex-row items-center justify-between py-6">
          <div>
            <img src="~/assets/images/handshake.png" class="w-24 xl:w-28" alt="Handshake CRM" style="width: 60px;" />
          </div>

        </div>
      </div>
      <div :class="[open ? 'flex' : 'hidden lg:flex']" class="space-x-3">
        <base-button
          class="px-8 xl:px-10 py-3 mt-2 bg-inherit text-gradient border border-[#00E54B]"
          @click="navigateToLogin"
        >
          Login
        </base-button>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { useRouter } from 'vue-router';

export default {
  name: 'BaseNavbar',
  data() {
    return {
      open: false,
      dropdownNavbar: false,
      router: useRouter(),
    }
  },
  methods: {
    dropdownToggler () {
      this.dropdownNavbar = !this.dropdownNavbar
    },
    navigateToLogin () {
      this.router.push("/login");
    },
  }
}
</script>
