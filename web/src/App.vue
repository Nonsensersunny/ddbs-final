<template>
  <div id="app">
    <el-container>
      <el-aside width="100px" style="background-color: rgb(238, 241, 246)">
        <el-menu default-active="1">
          <el-menu-item index="1">
            <router-link to="/" v-if="!this.$store.getters.loginStatus">主页</router-link>
          </el-menu-item>
          <el-menu-item index="2">
            <router-link to="/order" v-if="this.$store.getters.loginStatus">订单</router-link>
          </el-menu-item>
          <el-menu-item index="3">
            <router-link to="/trace"v-if="this.$store.getters.loginStatus">Trace</router-link>
          </el-menu-item>
          <el-menu-item index="4">
            <router-link to="/about" v-if="this.$store.getters.loginStatus">About</router-link>
          </el-menu-item>
          <el-menu-item index="5">
            <el-link @click="logout" v-if="this.$store.getters.loginStatus">退出</el-link>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-header></el-header>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        isLogin: this.$store.getters.loginStatus,
      }
    },
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
