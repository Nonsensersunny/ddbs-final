<template>
    <div id="order">
        <el-table
            :data="orders"
            :row-class-name="tableRowClassName"
            :default-sort="{prop: 'finished', order: 'ascending'}"
            :highlight-current-row="true">
            <el-table-column prop="time" label="创建时间" sortable>
                <template slot-scope="scope"><el-icon class="el-icon-time" /> {{ scope.row.time | timeFilter }}</template>
            </el-table-column>
            <el-table-column prop="start" label="始发地"></el-table-column>
            <el-table-column prop="dest" label="目的地"></el-table-column>
            <el-table-column prop="finished" label="状态" sortable>
                <template slot-scope="scope">
                    <el-tag v-if="scope.row.finished" type="success">已完成</el-tag>
                    <el-tag v-else type="danger">运输中</el-tag>
                </template>
            </el-table-column>
            <el-table-column align="right">
                <template slot="header" slot-scope="scope">
                    <el-button size="mini" type="primary" @click="getAllOrders"><el-icon class="el-icon-refresh"></el-icon></el-button>
                    <el-button v-if="role !== '2'" size="mini" type="success" @click="createOrderDialogVisible = true">新订单</el-button>
                </template>
                <template slot-scope="scope">
                    <el-button v-if="role === '2'" type="primary" size="mini" plain :disabled="scope.row.finished" @click="getOrderById(scope.row.id)">编辑</el-button>
                    <el-button v-if="role !== '2'" type="primary" size="mini" plain @click="getOrderById(scope.row.id)">查看</el-button>
                    <el-button v-if="role !== '2'" type="danger" size="mini" :disabled="scope.row.finished" @click="finishOrder(scope.row.id)" plain>结束</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-dialog
            title="创建订单"
            :visible.sync="createOrderDialogVisible">
            <el-form :model="newOrder">
                <el-form-item label="始发地">
                    <el-select v-model="newOrder.start" placeholder="请选择始发地">
                        <el-option
                                v-for="area in this.$store.getters.areas"
                                :key="area"
                                :label="area"
                                :value="area"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="目的地">
                    <el-select v-model="newOrder.dest" placeholder="请选择目的地">
                        <el-option
                                v-for="area in this.$store.getters.areas"
                                :key="area"
                                :label="area"
                                :value="area"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <el-button type="primary" @click="createOrder">确定</el-button>
            <el-button type="info" @click="createOrderDialogVisible = false">取消</el-button>
        </el-dialog>
        <el-dialog
                title="详细信息"
                :visible.sync="orderDetailDialogVisible">
            <p>订单ID：{{ currentOrder.id }}</p>
            <p>创建时间：{{ currentOrder.time | timeFilter }}</p>
            <el-steps :active="traces.length + 1" finish-status="success">
                <el-step :title="currentOrder.start" description="始发地"></el-step>
                <el-step v-for="trace in traces" :key="trace.id" :title="trace.current"></el-step>
                <el-step title="..." v-if="hasNextDest"></el-step>
                <el-step :title="currentOrder.dest" description="目的地"></el-step>
            </el-steps>
            <div v-if="role === '2'">
                <el-form :model="newTrace" status-icon>
                    <el-form-item label="订单ID" ><el-input v-model="currentOrder.id" disabled></el-input></el-form-item>
                    <el-form-item label="当前站"><el-input v-model="area" disabled></el-input></el-form-item>
                    <el-form-item label="下一站" required>
                        <el-select v-model="newTrace.next" placeholder="请选择地区">
                            <el-option
                                    v-for="area in this.$store.getters.areas"
                                    :key="area"
                                    :label="area"
                                    :value="area"></el-option>
                        </el-select>
                    </el-form-item>
                </el-form>
                <el-button type="primary" @click="createTrace">确定</el-button>
                <el-button type="info" @click="orderDetailDialogVisible = false">取消</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {Order, Trace} from "@/assets/js/type";

    export default {
        name: "Order",
        filters: {
            timeFilter(date) {
                return new Date(date).toLocaleString("zh");
            }
        },
        data() {
            return {
                orders: [],
                createOrderDialogVisible: false,
                newOrder: new Order(),
                orderDetailDialogVisible: false,
                currentOrder: {
                    id: "",
                    start: "",
                    dest: "",
                    time: "",
                    finished: false,
                },
                traces: [],
                role: this.$store.getters.role,
                area: this.$store.getters.area,
                newTrace: new Trace(),
                hasNextDest: true,
            }
        },
        methods: {
            async getAllOrders() {
                try {
                    let resp;
                    if (this.$store.getters.role === "2") {
                        resp = await this.$store.dispatch("adminGetAllOrders");
                    } else {
                        resp = await this.$store.dispatch("getAllOrders");
                    }
                    if (resp.code !== 0) {
                        this.$message.error(Error[resp.code]);
                    } else {
                        this.orders = resp.data.orders;
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            async createOrder() {
                try {
                    let resp = await this.$store.dispatch("createOrder", this.newOrder);
                    if (resp.code !== 0) {
                        this.$message.error(Error[resp.code]);
                    } else {
                        this.$message.success("创建订单成功!");
                        this.createOrderDialogVisible = false;
                        await this.getAllOrders();
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            async finishOrder(id) {
                try {
                    this.$confirm("确定完成此订单？", "提示", {
                        confirmButtonText: "确定",
                        cancelButtonText: "取消",
                        type: "warning"
                    }).then(async () => {
                        let resp = await this.$store.dispatch("finishOrder", id);
                        if (resp.code !== 0) {
                            this.$message.error(Error[resp.code]);
                        } else {
                            this.$message.success("订单已完成！");
                            await this.getAllOrders();
                        }
                    })
                } catch (e) {
                    console.log(e);
                }
            },
            async getOrderById(id) {
                try {
                    await this.getTraceByOrderId(id);
                    let resp = await this.$store.dispatch("getOrderById", id);
                    await this.getTraceByOrderId(id);
                    if (resp.code !== 0) {
                        this.$message.error(Error[resp.code]);
                        this.orderDetailDialogVisible = false;
                    } else {
                        this.orderDetailDialogVisible = true;
                        this.currentOrder = resp.data.order;
                    }
                } catch (e) {
                    console.log(e);
                }
            },
            async getTraceByOrderId(id) {
                try {
                    let resp = await this.$store.dispatch("getTraceByOrderId", id);
                    if (resp.code !== 0) {
                        this.$message.error(Error[resp.code]);
                        this.orderDetailDialogVisible = false;
                    } else {
                        this.traces = resp.data.traces;
                    }
                } catch (e) {
                    console.log(e);
                }
            },
            async createTrace() {
                try {
                    let that = this;
                    let startTrace = this.traces.filter(t => {
                        return t.current === that.area;
                    })[0];
                    let id = startTrace.id;
                    let next = that.newTrace.next;
                    let updateResp = await this.$store.dispatch("updateTraceNext", {id, next});
                    if (updateResp.code !== 0) {
                        this.$message.error(Error[updateResp.code]);
                        this.orderDetailDialogVisible = false;
                    } else {
                        this.$message.success("物流信息更新成功");
                        this.orderDetailDialogVisible = false;
                    }

                    // this.newTrace.oid = this.currentOrder.id;
                    // this.newTrace.current = this.area;
                    // let resp = await this.$store.dispatch("createTrace", this.newTrace);
                    // if (resp.code !== 0) {
                    //     this.$message.error(Error[resp.code]);
                    //     this.orderDetailDialogVisible = false;
                    // } else {
                    //     this.$message.success("物流信息更新成功");
                    //     this.orderDetailDialogVisible = false;
                    // }
                } catch (e) {
                    console.log(e);
                }
            },
            tableRowClassName({row}) {
                return row.finished ? "finished" : "";
            }
        },
        created() {
            this.getAllOrders();
        }
    }
</script>

<style scoped>
.el-table .finished {
    background: red;
}
</style>