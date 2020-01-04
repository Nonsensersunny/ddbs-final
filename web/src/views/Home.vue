<template>
  <div class="home">
    <el-card style="width: 50%; margin: 0 auto">
      <el-form :model="user">
        <el-form-item label="电话" required><el-input v-model="user.phone"></el-input></el-form-item>
        <el-form-item label="地区" required>
          <el-select v-model="user.area" placeholder="请选择地区">
            <el-option
                    v-for="area in this.$store.getters.areas"
                    :key="area"
                    :label="area"
                    :value="area"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="密码" required><el-input type="password" v-model="user.password"></el-input></el-form-item>
      </el-form>
      <div>
        <el-button @click="signup">注册</el-button>
        <el-button @click="signin">登录</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import {User, Error} from "@/assets/js/type";

export default {
  name: 'home',
  data() {
    return {
      user: new User(),
    }
  },
  methods: {
    async signup() {
      try {
        let resp = await this.$store.dispatch("userSignup", this.user);
        if (resp.code !== 0) {
          this.$message.error(Error[resp.code]);
        } else {
          this.$message.success("注册成功");
        }
      } catch (e) {
        console.log(e)
      }
    },
    async signin() {
      try {
        let resp = await this.$store.dispatch("userSignin", this.user);
        if (resp.code !== 0) {
          this.$message.error(Error[resp.code]);
        } else {
          this.$message.success("登录成功");
          this.$store.commit("modifyLoginStatus", true);
          this.$store.commit("modifyArea", this.$cookie.get("area"));
          this.$store.commit("modifyRole", this.$cookie.get("role"));
          await this.$router.push("/order");
        }
      } catch (e) {
        console.log(e)
      }
    }
  },
  created() {
    this.$store.dispatch("getAreas");
  },
  components: {

  }
}
</script>
