<template>
  <view class="my-center">
    <nav-bar title="我的" leftIcon="arrowleft" backgroundColor="#ffc200" color="#fff" @clickLeft="goBack" />
    <view class="my_info">
      <view class="head">
        <button class="avatar-btn" open-type="chooseAvatar" @chooseavatar="onChooseAvatar">
          <image class="head_image" :src="psersonUrl || '/static/btn_waiter_sel.png'" />
        </button>
      </view>
      <view class="phone_name">
        <view class="name">
          <input class="name_input" type="nickname" v-model="nickName" placeholder="点击设置昵称" @blur="saveNickName" />
          <image v-if="gender == '2'" class="name_type" src="@/static/girl.png" />
          <image v-if="gender == '1'" class="name_type" src="@/static/boy.png" />
        </view>
        <view class="phone"><text class="phone_text">{{ phoneNumber }}</text></view>
      </view>
    </view>
    <view class="container">
      <view class="box address_order">
        <view class="address" @tap.stop="goAddress">
          <image class="location" src="@/static/address.png" />
          <text class="address_word">地址管理</text>
          <image class="to_right" src="@/static/toRight.png" mode="aspectFit" />
        </view>
        <view class="order" @tap.stop="goOrder">
          <image class="location" src="@/static/order.png" />
          <text class="order_word">历史订单</text>
          <image class="to_right" src="@/static/toRight.png" mode="aspectFit" />
        </view>
      </view>
      <view v-if="recentOrdersList && recentOrdersList.length > 0" class="recent">
        <text class="order_line">最近订单</text>
      </view>
      <scroll-view :enhanced="true" :style="{ height: scrollH + 'px' }" scroll-y @scrolltolower="lower">
        <view class="main recent_orders">
          <view v-for="(item, index) in recentOrdersList" :key="index" class="box order_lists">
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
              <button class="new_btn" type="default" @tap.stop="oneOrderFun(item.id)">再来一单</button>
              <button v-if="item.status === 1 && getOvertime(item.orderTime) > 0" class="new_btn btn" type="default" @tap.stop="goDetail(item.id)">去支付</button>
            </view>
          </view>
        </view>
        <view v-if="loading" class="loading-tip">{{ loadingText }}</view>
      </scroll-view>
    </view>
  </view>
</template>
<script>
import { getOrderPage, repetitionOrder } from '@/utils/api'
import { mapMutations } from 'vuex'
import { statusWord, getOvertime } from '@/utils/index'
export default {
  data() {
    return {
      psersonUrl: '', nickName: '', gender: '0', phoneNumber: '18500557668',
      recentOrdersList: [], scrollH: 0,
      pageInfo: { page: 1, pageSize: 10, total: 0 }, loadingText: '', loading: false
    }
  },
  methods: {
    ...mapMutations(['setAddressBackUrl']),
    numes(list) {
      let count = 0, total = 0
      if (list && list.length > 0) {
        list.forEach(obj => { count += Number(obj.number); total += Number(obj.number) * Number(obj.amount) })
      }
      return { count, total }
    },
    statusWord, getOvertime,
    async getList() {
      const res = await getOrderPage({ pageSize: 10, page: this.pageInfo.page })
      if (res.code === 1) {
        this.recentOrdersList = this.recentOrdersList.concat(res.data.records)
        this.pageInfo.total = res.data.total
        this.loadingText = ''
        this.loading = false
      }
    },
    goAddress() {
      this.setAddressBackUrl('/pages/my/my')
      uni.redirectTo({ url: '/pages/address/address?form=my' })
    },
    goOrder() { uni.navigateTo({ url: '/pages/historyOrder/historyOrder' }) },
    async oneOrderFun(id) {
      await repetitionOrder(id)
      uni.switchTab({ url: '/pages/index/index' })
    },
    goDetail(id) {
      this.setAddressBackUrl('/pages/my/my')
      uni.redirectTo({ url: '/pages/details/index?orderId=' + id })
    },
    dataAdd() {
      const pages = Math.ceil(this.pageInfo.total / 10)
      if (this.pageInfo.page >= pages) { this.loadingText = '没有更多了'; this.loading = true }
      else { this.pageInfo.page++; this.getList() }
    },
    lower() { this.loadingText = '数据加载中...'; this.loading = true; this.dataAdd() },
    onChooseAvatar(e) {
      const { avatarUrl } = e.detail
      this.psersonUrl = avatarUrl
      this.uploadAvatar(avatarUrl)
    },
    uploadAvatar(tempPath) {
      uni.uploadFile({
        url: 'http://localhost:8080/user/user/avatar',
        filePath: tempPath,
        name: 'file',
        header: { 'Authorization': this.$store.state.token || '' },
        success: (res) => {
          const data = JSON.parse(res.data)
          if (data.code === 1) {
            this.$store.commit('setBaseUserInfo', { ...this.$store.state.baseUserInfo, avatarUrl: data.data })
          }
        }
      })
    },
    saveNickName() {
      if (!this.nickName) return
      uni.request({
        url: 'http://localhost:8080/user/user/profile',
        method: 'PUT',
        data: { nickName: this.nickName },
        header: { 'Content-Type': 'application/json', 'Authorization': this.$store.state.token || '' },
        success: () => {
          this.$store.commit('setBaseUserInfo', { ...this.$store.state.baseUserInfo, nickName: this.nickName })
        }
      })
    },
    goBack() { uni.redirectTo({ url: '/pages/index/index' }) }
  },
  onLoad() {
    const info = this.$store.state.baseUserInfo
    if (info) {
      this.psersonUrl = info.avatarUrl || ''
      this.nickName = info.nickName || ''
      this.gender = info.gender || '0'
    }
    this.phoneNumber = this.$store.state.shopPhone || '18500557668'
    this.getList()
  },
  onReady() {
    uni.getSystemInfo({ success: res => { this.scrollH = res.windowHeight - uni.upx2px(100) } })
  }
}
</script>
<style scoped>
.my-center { background: #f6f6f6; height: 100%; }
.my_info { height: 172rpx; width: 750rpx; background-color: #ffc200; display: flex; padding-top: 80rpx; }
.head { width: 172rpx; height: 172rpx; display: flex; align-items: center; justify-content: center; }
.avatar-btn { width: 116rpx; height: 116rpx; padding: 0; margin: 0; border: none; background: transparent; line-height: 1; }
.avatar-btn::after { border: none; }
.head_image { width: 116rpx; height: 116rpx; border-radius: 50%; background-color: #fff; display: block; }
.name_input { font-size: 32rpx; font-weight: 550; color: #333; height: 44rpx; line-height: 44rpx; padding: 0; background: transparent; }
.phone_name { flex: 1; display: flex; flex-direction: column; justify-content: center; }
.name { display: flex; align-items: center; margin-bottom: 8rpx; }
.name_text { font-size: 32rpx; font-weight: 550; color: #333; margin-right: 12rpx; }
.name_type { width: 32rpx; height: 32rpx; }
.phone_text { font-size: 28rpx; color: #333; }
.container { margin-top: 20rpx; }
.address_order { width: 710rpx; margin: 0 auto; background: #fff; border-radius: 12rpx; padding: 0 30rpx; }
.address, .order { display: flex; align-items: center; height: 100rpx; position: relative; }
.order { border-top: 1px dashed #ebebeb; }
.location { width: 34rpx; height: 36rpx; margin-right: 12rpx; }
.address_word, .order_word { font-size: 28rpx; color: #333; flex: 1; }
.to_right { width: 30rpx; height: 30rpx; position: absolute; right: 0; }
.recent { padding: 30rpx 22rpx 10rpx; }
.order_line { font-size: 28rpx; font-weight: 550; color: #333; }
.order_lists { background: #fff; border-radius: 12rpx; padding: 20rpx; margin: 20rpx 22rpx; }
.date_type { display: flex; justify-content: space-between; margin-bottom: 16rpx; }
.time { font-size: 24rpx; color: #999; }
.type { font-size: 24rpx; }
.type.status { color: #f58c21; }
.orderBox { margin-bottom: 16rpx; }
.food_num { margin-bottom: 12rpx; }
.pic { white-space: nowrap; }
.food_num_item { display: inline-block; margin-right: 20rpx; text-align: center; }
.food_num_item image { width: 80rpx; height: 80rpx; border-radius: 8rpx; }
.food { font-size: 20rpx; color: #666; margin-top: 6rpx; }
.numAndAum { display: flex; justify-content: space-between; font-size: 26rpx; color: #333; }
.againBtn { display: flex; gap: 20rpx; }
.new_btn { flex: 1; height: 60rpx; line-height: 60rpx; border-radius: 30rpx; background: #f5f5f5; font-size: 24rpx; color: #333; text-align: center; }
.new_btn.btn { background: #ffc200; }
.loading-tip { text-align: center; padding: 20rpx; color: #999; font-size: 24rpx; }
</style>
