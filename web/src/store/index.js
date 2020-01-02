import Vue from 'vue'
import Vuex from 'vuex'
import {GuestHttp, UserHttp, AdminHttp} from "../assets/js/api";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: false,
    area: "",
    phone: "",
    role: "1",
    areas: []
  },
  getters: {
    areas: (state) => {
      return state.areas;
    },
    loginStatus: (state) => {
      return state.isLogin;
    },
    role: (state) => {
      return state.role;
    },
    area: (state) => {
      return state.area;
    }
  },
  mutations: {
    modifyArea(state, area) {
      state.area = area;
    },
    modifyPhone(state, phone) {
      state.phone = phone;
    },
    modifyLoginStatus(state, status) {
      state.isLogin = status;
    },
    modifyRole(state, role) {
      state.role = role;
    },
    getAreas(state, areas) {
      state.areas = areas;
    }
  },
  actions: {
    async userSignup(context, playload) {
      return await GuestHttp.signup(playload);
    },
    async userSignin(context, playload) {
      return await GuestHttp.signin(playload);
    },
    async userLogout(context) {
      return await UserHttp.logout();
      context.commit("modifyLoginStatus", false);
    },
    async getAreas(context) {
      let resp = await GuestHttp.getAreas();
      context.commit('getAreas', resp.data.areas);
    },
    async createOrder(context, playload) {
      return await UserHttp.createOrder(playload);
    },
    async getAllOrders(context) {
      return await UserHttp.getAllOrders();
    },
    async getOrderById(context, id) {
      return await UserHttp.getOrderById(id);
    },
    async finishOrder(context, id) {
      return await UserHttp.finishOrder(id);
    },
    async getTraceByOrderId(context, id) {
      return await UserHttp.getTraceByOrderId(id);
    },
    async createTrace(context, playload) {
      return await AdminHttp.createTrace(playload);
    },
    async adminGetAllOrders(context) {
      return await AdminHttp.getAllOrders();
    },
    async updateTraceNext(context, {id, next}) {
      return await AdminHttp.updateTraceNext(id, next);
    }
  },
  modules: {
  }
})
