<template>
  <div class="rider-container">
    <div class="rider-header">
      <span class="rider-name">🚀 {{ riderName }}</span>
      <el-button type="danger" size="small" plain @click="logout">退出</el-button>
    </div>

    <el-tabs v-model="activeTab">
      <el-tab-pane label="待抢订单" name="pending">
        <el-table :data="pendingOrders" border stripe v-loading="loading" empty-text="暂无待派送订单">
          <el-table-column prop="number" label="订单号" width="200" />
          <el-table-column prop="address" label="配送地址" show-overflow-tooltip />
          <el-table-column prop="consignee" label="收货人" width="100" />
          <el-table-column prop="amount" label="金额" width="100">
            <template slot-scope="s">¥{{ s.row.amount.toFixed(2) }}</template>
          </el-table-column>
          <el-table-column label="操作" width="100">
            <template slot-scope="s">
              <el-button type="warning" size="small" @click="grab(s.row.id)">抢单</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <el-tab-pane label="我的订单" name="my">
        <el-table :data="myOrders" border stripe v-loading="loadingMy" empty-text="还没有抢到订单">
          <el-table-column prop="number" label="订单号" width="200" />
          <el-table-column prop="address" label="配送地址" show-overflow-tooltip />
          <el-table-column prop="consignee" label="收货人" width="100" />
          <el-table-column prop="amount" label="金额" width="100">
            <template slot-scope="s">¥{{ s.row.amount.toFixed(2) }}</template>
          </el-table-column>
          <el-table-column label="状态" width="100">
            <template slot-scope="s">
              <el-tag :type="s.row.status === 4 ? 'primary' : 'success'">
                {{ s.row.status === 4 ? '配送中' : '已完成' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="100">
            <template slot-scope="s">
              <el-button v-if="s.row.status === 4" type="success" size="small" @click="complete(s.row.id)">送达</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script lang="ts">
import axios from 'axios'

export default {
  data() {
    return {
      activeTab: 'pending',
      riderName: '',
      pendingOrders: [] as any[],
      myOrders: [] as any[],
      loading: false,
      loadingMy: false
    }
  },
  created() {
    const info = localStorage.getItem('rider_info')
    if (!info) { this.$router.push('/rider/login'); return }
    this.riderName = JSON.parse(info).name
    this.fetchPending()
  },
  methods: {
    getToken() { return localStorage.getItem('rider_token') || '' },
    authHeader() { return { headers: { token: this.getToken() } } },
    async fetchPending() {
      this.loading = true
      try {
        const res = await axios.get('/api/rider/orders/pending', this.authHeader())
        if (res.data.code === 1) this.pendingOrders = res.data.data || []
      } finally { this.loading = false }
    },
    async fetchMy() {
      this.loadingMy = true
      try {
        const res = await axios.get('/api/rider/orders/my', this.authHeader())
        if (res.data.code === 1) this.myOrders = res.data.data || []
      } finally { this.loadingMy = false }
    },
    async grab(orderId: number) {
      try {
        const res = await axios.get(`/api/rider/orders/grab?id=${orderId}`, this.authHeader())
        if (res.data.code === 1) {
          this.$message.success('抢单成功！')
          this.fetchPending()
          this.fetchMy()
        } else {
          this.$message.warning(res.data.msg)
          this.fetchPending()
        }
      } catch { this.$message.error('抢单失败') }
    },
    async complete(orderId: number) {
      try {
        const res = await axios.get(`/api/rider/orders/complete?id=${orderId}`, this.authHeader())
        if (res.data.code === 1) {
          this.$message.success('已送达')
          this.fetchMy()
        } else {
          this.$message.error(res.data.msg)
        }
      } catch { this.$message.error('操作失败') }
    },
    logout() {
      localStorage.removeItem('rider_token')
      localStorage.removeItem('rider_info')
      this.$router.push('/rider/login')
    }
  },
  watch: {
    activeTab(val) { if (val === 'my') this.fetchMy() }
  }
}
</script>

<style scoped>
.rider-container { padding: 20px; min-height: 100vh; background: #f0f2f5; }
.rider-header { display: flex; justify-content: space-between; align-items: center; padding: 16px 24px; background: #fff; border-radius: 8px; margin-bottom: 16px; }
.rider-name { font-size: 20px; font-weight: 600; }
</style>
