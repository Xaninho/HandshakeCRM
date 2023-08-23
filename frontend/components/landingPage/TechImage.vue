<template>
  <div>
    <img style="max-width: 80px;
    width: 100%;
    max-height: 100px;
    margin: 20px;" :src="`assets/images/technologies/${img}`" class="sm:w-1/2 lg:w-72 mx-auto" alt="" />
  </div>
</template>
<script>
export default {
  name: 'TechImage',
  props: {
    img: {
      type: String,
      default: '',
    },
  },
}
</script>
