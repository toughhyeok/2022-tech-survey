<template>
  <MyContents />
  <MyFooter />
</template>

<script>
import MyContents from './components/MyContents.vue'
import MyFooter from './components/MyFooter.vue'

export default {
  name: 'App',
  components: {
    MyContents,
    MyFooter
  }
}
</script>

<style>
#app {
  font-family: Inter,Helvetica Neue,Arial,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #666;
  margin-top: 60px;
}
</style>
