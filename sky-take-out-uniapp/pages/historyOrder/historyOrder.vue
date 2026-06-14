<template>
  <view class="history_order">
    <nav-bar title="历史订单" leftIcon="arrowleft" backgroundColor="#333333" @clickLeft="goBack" />
    <scroll-view :enhanced="true" class="scroll-row" scroll-x :scroll-into-view="scrollinto" :scroll-with-animation="true">
      <view v-for="(item, index) in tabBars" :key="index" class="scroll-row-item" :id="'tab' + index" @tap.stop="changeTab(index)">
        <view :class="tabIndex == index ? 'scroll-row-item-act' : ''">{{ item }}</view>
      </view>
    </scroll-view>
    <swiper :style="{ height: scrollH + 'px' }" :current="tabIndex" @change="onChangeSwiperTab">
      <swiper-item v-for="(_, index) in tabBars" :key="index">
        <scroll-view :enhanced="true" :style="{ height: scrollH + 'px' }" scroll-y @scrolltolower="lower">
          <view v-if="recentOrdersList && recentOrdersList.length > 0" class="main recent_orders">
            <view v-for="(item, idx) in recentOrdersList" :key="idx" class="box order_lists">
              <view class="date_type">
                <text class="time">{{ item.orderTime }}</text>
                <text :class="['type', 'status', item.status == 2 ? 'status' : '']">{{ statusWord(item.status, getOvertime(item.orderTime)) }}</text>
              </view>
              <view class="orderBox" @tap.stop="goDetail(item.id)">
                <view class="food_num">
                  <scroll-view :enhanced="true" class="pic" scroll-x>
                    <view v-for="(num, y) in item.orderDetailList" :key="y" class="food_num_item">
                      <view class="img"><image :src="num.image" /></view>
                      <view class="food">{{ num.name }}</view>
                    </view>
                  </scroll-view>
                </view>
                <view class="numAndAum">
                  <view><text>{{ '￥' + numes(item.orderDetailList).total.toFixed(2) }}</text></view>
                  <view><text>{{ '共' + numes(item.orderDetailList).count + '件' }}</text></view>
                </view>
              </view>
              <view class="againBtn">
                <button class="new_btn" type="default" @tap.stop="oneMoreOrder(item.id)">再来一单</button>
                <button v-if="item.status === 1 && getOvertime(item.orderTime) > 0" class="new_btn btn" type="default" @tap.stop="goDetail(item.id)">去支付</button>
                <button v-if="item.status === 2" class="new_btn btn" type="default" @tap.stop="handleReminder('center', item.id)">催单</button>
              </view>
            </view>
          </view>
        </scroll-view>
      </swiper-item>
    </swiper>
  </view>
</template>
<script>
import { getOrderPage, repetitionOrder, reminder } from '@/utils/api'
import { statusWord, getOvertime } from '@/utils/index'
import { mapMutations } from 'vuex'
export default {
  components: { 'nav-bar': () => import('@/components/nav-bar/nav-bar.vue') },
  data() {
    return {
      recentOrdersList: [], pageInfo: { page: 1, pageSize: 10, total: 0 },
      tabIndex: 0, tabBars: ['全部订单', '待付款', '已取消'],
      scrollH: 0, scrollinto: 'tab0', loadingType: 0
    }
  },
  methods: {
    ...mapMutations(['setAddressBackUrl']),
    numes(list) {
      let count = 0, total = 0
      if (list && list.length > 0) { list.forEach(obj => { count += Number(obj.number); total += Number(obj.number) * Number(obj.amount) }) }
      return { count, total }
    },
    statusWord, getOvertime,
    async getList() {
      let params = { pageSize: 10, page: this.pageInfo.page }
      if (this.tabIndex === 1) params.status = 1
      if (this.tabIndex === 2) params.status = 6
      const res = await getOrderPage(params)
      if (res.code === 1) { this.recentOrdersList = this.recentOrdersList.concat(res.data.records); this.pageInfo.total = res.data.total }
    },
    async oneMoreOrder(id) { await repetitionOrder(id); uni.switchTab({ url: '/pages/index/index' }) },
    changeTab(index) { this.tabIndex = index; this.recentOrdersList = []; this.pageInfo.page = 1; this.getList() },
    onChangeSwiperTab(e) { this.changeTab(e.detail.current) },
    dataAdd() {
      const pages = Math.ceil(this.pageInfo.total / 10)
      if (this.pageInfo.page < pages) { this.pageInfo.page++; this.getList() }
    },
    lower() { this.dataAdd() },
    goDetail(id) { this.setAddressBackUrl('/pages/historyOrder/historyOrder'); uni.redirectTo({ url: '/pages/details/index?orderId=' + id }) },
    handleReminder(type, id) { reminder(id); uni.showToast({ title: '已催单', icon: 'success' }) },
    goBack() { uni.navigateBack() }
  },
  onLoad() { this.getList() },
  onReady() { uni.getSystemInfo({ success: res => { this.scrollH = res.windowHeight - uni.upx2px(100) } }) }
}
</script>
<style scoped>
.history_order { height: 100vh; background: #f6f6f6; }
.scroll-row { white-space: nowrap; background: #fff; padding: 100rpx 0 20rpx; display: flex; }
.scroll-row-item { display: inline-block; text-align: center; padding: 0 30rpx; }
.scroll-row-item view { padding: 10rpx 0; font-size: 28rpx; color: #666; }
.scroll-row-item-act { color: #333; font-weight: 600; border-bottom: 4rpx solid #ffc200; }
.recent_orders { padding: 20rpx; }
.order_lists { background: #fff; border-radius: 12rpx; padding: 20rpx; margin-bottom: 20rpx; }
.date_type { display: flex; justify-content: space-between; margin-bottom: 16rpx; }
.time { font-size: 24rpx; color: #999; }
.type.status { color: #f58c21; }
.orderBox { margin-bottom: 16rpx; }
.food_num { margin-bottom: 12rpx; }
.pic { white-space: nowrap; }
.food_num_item { display: inline-block; margin-right: 20rpx; text-align: center; }
.food_num_item .img image { width: 80rpx; height: 80rpx; border-radius: 8rpx; }
.food { font-size: 20rpx; color: #666; margin-top: 6rpx; }
.numAndAum { display: flex; justify-content: space-between; font-size: 26rpx; color: #333; }
.againBtn { display: flex; gap: 20rpx; }
.new_btn { flex: 1; height: 60rpx; line-height: 60rpx; border-radius: 30rpx; background: #f5f5f5; font-size: 24rpx; color: #333; text-align: center; }
.new_btn.btn { background: #ffc200; }
</style>
