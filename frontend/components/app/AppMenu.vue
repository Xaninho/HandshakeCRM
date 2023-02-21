<template>
    <div class="layout-menu-container">
      <AppSubMenu :items="model" class="layout-menu" :root="true" @menu-item-click="onMenuItemClick" />
    </div>
</template>

<script setup lang='ts'>

    import AppSubMenu from './AppSubMenu.vue'
    const props = defineProps<{
    model: Array<any>
    }>()
    const emit = defineEmits(['menuItemClick'])
    function onMenuItemClick(event: any) {
    emit('menuItemClick', event)
    }

</script>