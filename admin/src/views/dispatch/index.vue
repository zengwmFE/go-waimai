<template>
  <div class="dispatch-container">
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>骑手派单调度</span>
      </div>

      <!-- 待派发订单列表 -->
      <div class="queue-section">
        <h3>待派单队列</h3>
        <el-table :data="pendingOrders" border style="width: 100%">
          <el-table-column prop="number" label="订单号" width="200" />
          <el-table-column prop="address" label="配送地址" />
          <el-table-column prop="amount" label="金额" width="120" />
          <el-table-column prop="orderTime" label="下单时间" width="180" />
          <el-table-column label="操作" width="120">
            <template slot-scope="scope">
              <el-button type="text" @click="grabOrder(scope.row)">抢单</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- 模拟骑手批量抢单 -->
      <div class="simulate-section">
        <h3>模拟骑手派单</h3>
        <div class="simulate-controls">
          <el-input-number v-model="riderCount" :min="1" :max="20" label="骑手数量" />
          <el-button type="primary" @click="simulateDispatch" style="margin-left: 16px">
            模拟 {{ riderCount }} 个骑手同时抢单
          </el-button>
        </div>
        <div v-if="dispatchResults.length" class="results">
          <h4>抢单结果</h4>
          <div v-for="(item, i) in dispatchResults" :key="i" class="result-item">
            {{ item }}
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts">
import {
  getOrderDetailPage
} from '@/api/order'
export default {
  data() {
    return {
      pendingOrders: [],
      riderCount: 5,
      dispatchResults: [],
      pageSize: 10,
      page: 1,
    }
  },
  created() {
    this.init()
  },
  methods: {
    async   init() {
    const params = {
      page: this.page,
      pageSize: this.pageSize,
      status: 3,
    }
    getOrderDetailPage({ ...params })
      .then((res) => {
        if (res.data.code === 1) {
          this.pendingOrders = res.data.data.records
          this.orderStatus = 3
          this.counts = Number(res.data.data.total)
        } else {
          this.$message.error(res.data.msg)
        }
      })
      .catch((err) => {
        this.$message.error('请求出错了：' + err.message)
      })
  },
    async grabOrder(order) {},
    async simulateDispatch() {}
  }
}
</script>

<style scoped>
.dispatch-container { padding: 20px; }
.queue-section { margin-bottom: 30px; }
.simulate-controls { display: flex; align-items: center; margin-bottom: 16px; }
.results { background: #f5f7fa; padding: 16px; border-radius: 4px; margin-top: 16px; }
.result-item { padding: 4px 0; font-family: monospace; }
</style>
