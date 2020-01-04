import axios from "axios";
import {server} from "@/assets/js/config";

export const GuestHttp = {
    client: axios.create({
        baseURL: `http://${server.host}:${server.port}`,
        withCredentials: true
    }),
    async signup(user) {
        let res = await this.client.put('user/', user);
        return res.data;
    },
    async signin(user) {
        let res = await this.client.post('user/', user);
        return res.data;
    },
    async getAreas() {
        let res = await this.client.get('user/area');
        return res.data;
    }
}

export const UserHttp = {
    client: axios.create({
        baseURL: `http://${server.host}:${server.port}`,
        withCredentials: true
    }),
    async logout() {
        let res = await this.client.delete('user/');
        return res.data;
    },
    async createOrder(order) {
        let res = await this.client.put('order/', order);
        return res.data;
    },
    async getAllOrders() {
        let res = await this.client.get('order/');
        return res.data;
    },
    async getOrderById(id) {
        let res = await this.client.get(`order/?id=${id}`);
        return res.data;
    },
    async finishOrder(id) {
        let res = await this.client.delete(`order/?id=${id}`);
        return res.data;
    },
    async getTraceByOrderId(id) {
        let res = await this.client.get(`trace/?oid=${id}`);
        return res.data;
    }
}

export const AdminHttp = {
    client: axios.create({
        baseURL: `http://${server.host}:${server.port}`,
        withCredentials: true
    }),
    async createTrace(trace) {
        let res = await this.client.put('admin/trace', trace);
        return res.data;
    },
    async getAllOrders() {
        let res = await this.client.get('admin/order');
        return res.data;
    },
    async updateTraceNext(id, next) {
        let formData = new FormData();
        formData.append("id", id);
        formData.append("next", next);
        let res = await this.client.post(`admin/trace`, {
            id: id,
            next: next,
        });
        return res.data;
    }
}