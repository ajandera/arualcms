<template>
  <aside>
    <ul>
      <li v-for="(item, index) in menu" class="item mt-2" v-bind:key="index">
        <a href="#">{{ item }}</a>
      </li>
    </ul>
  </aside>
</template>

<script>

export default {
  name: 'Menu',
  data() {
    return {
      menu: ["Posts", "Texts"]
    }
  }
}
</script>

<style lang="css" scoped>
aside ul {
  list-style: none;
}

</style>
