
<template>
    <div class="layout-topbar">
      <NuxtLink to="/" class="layout-topbar-logo">
        <span>Handshake CRM</span>
      </NuxtLink>
      <button class="p-link layout-menu-button layout-topbar-button" @click="onMenuToggle">
        <i class="pi pi-bars" />
      </button>
  
      <button
        v-styleclass="{
          selector: '@next',
          enterClass: 'hidden',
          enterActiveClass: 'scalein',
          toggleClass: 'hidden',
          leaveActiveClass: 'fadeout',
          hideOnOutsideClick: true,
        }"
        class="p-link layout-topbar-menu-button layout-topbar-button"
      >
        <i class="pi pi-ellipsis-v" />
      </button>
  
      <ul class="layout-topbar-menu hidden lg:flex origin-top">
        <li>
          <button class="p-link layout-topbar-button">
            <i class="pi pi-user" />
            <span>Profile</span>
          </button>
        </li>
        <li>
          <button class="p-link layout-topbar-button" @click="toggle">
            <i class="pi pi-cog" />
            <span>Settings</span>
          </button>
        </li>
      </ul>

      <client-only>
        <OverlayPanel id="overlay_panel" ref="op" append-to="body" style="width: 200px">
          <div class="field-radiobutton">
            <RadioButton
              id="saga" v-model="themeStore.themeName" name="layoutColorMode" value="saga"
              @change="themeStore.setLight()"
            />
            <label>Light</label>
          </div>
          <div class="field-radiobutton">
            <RadioButton
              id="arya" v-model="themeStore.themeName" name="layoutColorMode" value="arya"
              @change="themeStore.setDark()"
            />
            <label>Dark</label>
          </div>
        </OverlayPanel>
      </client-only>
    </div>
</template>

<script setup lang='ts'>
import { useThemeStore } from '@/stores'

const emit = defineEmits(['menuToggle'])
const themeStore = useThemeStore()
const op = ref<any>(null)
function toggle(event: any) {
  op.value.toggle(event)
}
function redirectToGithub(event: any) {
  window.open('https://github.com/sfxcode/nuxt3-primevue-starter', '_blank')
}

function onMenuToggle(event: any) {
  emit('menuToggle', event)
}

</script>

<style scoped>
.hidden {
display: none !important;
}
@media (min-width: 1024px) {
.lg\:flex {
    display: flex !important;
}
}
</style>