<template>
  <div id="app">
    <div id="nav">
      <router-link to="/" v-if="!this.$store.getters.loginStatus">Home</router-link>
      <router-link to="/order" v-if="this.$store.getters.loginStatus">Order |</router-link>
      <router-link to="/trace"v-if="this.$store.getters.loginStatus">Trace |</router-link>
      <router-link to="/about" v-if="this.$store.getters.loginStatus">About |</router-link>
      <el-link @click="logout" v-if="this.$store.getters.loginStatus">退出</el-link>
    </div>
    <router-view/>
  </div>
</template>

<script>
  export default {
    methods: {
      async logout() {
        try {
          let resp = await this.$store.dispatch("userLogout");
          if (resp.code !== 0) {
            this.$message.error(Error[resp.code]);
          } else {
            this.$message.success("退出成功");
            await this.$router.push("/");
          }
        } catch (e) {
          console.log(e)
        }
      }
    }
  }
</script>

<style lang="stylus">
/*#app*/
/*  font-family 'Avenir', Helvetica, Arial, sans-serif*/
/*  -webkit-font-smoothing antialiased*/
/*  -moz-osx-font-smoothing grayscale*/
/*  text-align center*/
/*  color #2c3e50*/
/*  margin-top 60px*/
</style>
