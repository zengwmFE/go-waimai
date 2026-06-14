<template>
  <view class="order_content orderDetail">
    <nav-bar title="订单详情" leftIcon="arrowleft" backgroundColor="#333333" @clickLeft="goBack" />
    <scroll-view :enhanced="true" class="order_content_box" scroll-y>
      <view class="box">
        <view class="orderInfoTip">
          <view class="tit">{{ statusWord(orderDetailsData.status, isOvertime ? 0 : 999) }}</view>
          <view v-if="!timeout && orderDetailsData.status === 1" class="time">
            <view class="timeIcon" />
            等待支付：<text>{{ rocallTime }}</text>
          </view>
          <view class="againBtn">
            <button v-if="!timeout && orderDetailsData.status === 1 || [2,3,4].includes(orderDetailsData.status)" class="new_btn" type="default" @tap.stop="handleCancel('center', orderDetailsData)">取消订单</button>
            <button v-if="!timeout && orderDetailsData.status === 1" class="new_btn btn" type="default" @tap.stop="handlePay(orderDetailsData.id)">立即支付</button>
            <button v-if="orderDetailsData.status === 2" class="new_btn btn" type="default" @tap.stop="handleReminder('center', orderDetailsData.id)">催单</button>
            <button v-if="orderDetailsData.status === 5" class="new_btn" type="default" @tap.stop="handleRefund('center')">申请退款</button>
            <button v-if="orderDetailsData.status !== 7" class="new_btn" type="default" @tap.stop="oneMoreOrder(orderDetailsData.id)">再来一单</button>
          </view>
        </view>
      </view>
      <view v-if="!timeout && orderDetailsData.status === 1" class="box timeTip">
        <view class="icon newIcon" />请在15分钟内完成支付，超时将自动取消。
      </view>
      <view v-if="orderDetailsData.status === 6 && orderDetailsData.payStatus === 2" class="box timeTip">
        <view class="icon moneyIcon" />您的订单已<text style="color:#ffc200">退款成功</text>。
      </view>
      <view class="box order_list_cont">
        <view class="order_list">
          <view class="word_text"><text class="word_style">苍穹餐厅</text></view>
          <view class="order-type">
            <view v-for="(obj, index) in orderDataes" :key="index" class="type_item">
              <view class="dish_img"><image class="dish_img_url" mode="aspectFill" :src="obj.image" /></view>
              <view class="dish_info">
                <view class="dish_name">{{ obj.name }}</view>
                <view class="dish_dishFlavor">{{ obj.dishFlavor || '' }}</view>
                <view class="dish_price">×<text class="dish_number">{{ obj.number }}</text></view>
                <view class="dish_active"><text>￥</text>{{ (obj.amount * obj.number).toFixed(2) }}</view>
              </view>
            </view>
            <view v-if="(orderDetailsData.orderDetailList || []).length > 2" class="iconUp" @tap.stop="showDisplay = !showDisplay">
              <text>{{ showDisplay ? '点击收起' : '展开更多' }}</text>
              <image :class="['icon_img', showDisplay ? 'icon_imgDown' : '']" src="@/static/toRight.png" />
            </view>
            <view class="orderList">
              <view class="orderInfo"><text class="text">打包费</text><text class="may">￥</text>{{ orderDetailsData.packAmount || 0 }}</view>
              <view class="orderInfo"><text class="text">配送费</text><text class="may">￥</text>6</view>
              <view class="totalMoney">合计<text class="text"><text>￥</text>{{ orderDetailsData.amount }}</text></view>
            </view>
          </view>
        </view>
      </view>
      <view class="box contactMerchant" @tap.stop="handlePhone('bottom')"><view class="phoneIcon" />联系商家</view>
      <view class="box"><view class="orderBaseInfo">
        <view><view>期望时间</view><view>{{ orderDetailsData.deliveryStatus === 1 ? '立即送出' : orderDetailsData.estimatedDeliveryTime }}</view></view>
        <view><view>配送地址</view><view><view class="nameInfo"><text>{{ orderDetailsData.consignee }}</text>{{ orderDetailsData.phone }}</view><view>{{ orderDetailsData.address }}</view></view></view>
      </view></view>
      <view class="box"><view class="orderBaseInfo">
        <view><view>订单号码</view><view>{{ orderDetailsData.number }}</view></view>
        <view><view>订单时间</view><view>{{ orderDetailsData.orderTime }}</view></view>
        <view><view>支付方式</view><view>{{ orderDetailsData.payMethod === 1 ? '微信' : '支付宝' }}</view></view>
        <view><view>餐具数量</view><view>{{ orderDetailsData.tablewareNumber }}</view></view>
      </view></view>
    </scroll-view>
  </view>
</template>
<script>
import { getOrderDetail, cancelOrder, repetitionOrder, reminder } from '@/utils/api'
import { statusWord, getOvertime } from '@/utils/index'
import { mapState, mapMutations } from 'vuex'
export default {
  components: { 'nav-bar': () => import('@/components/nav-bar/nav-bar.vue') },
  data() { return { showDisplay: false, rocallTime: '', orderDetailsData: {}, timeout: false, isOvertime: false } },
  computed: { ...mapState(['shopInfo', 'shopPhone']), phone() { return this.shopPhone }, orderDataes() { const list = this.orderDetailsData.orderDetailList || []; return this.showDisplay ? list : list.slice(0, 2) } },
  methods: {
    ...mapMutations(['setOrderData']),
    statusWord,
    async getBaseData(id) { const res = await getOrderDetail(id); if (res.code === 1) { this.orderDetailsData = res.data; this.runTimeBack(res.data.orderTime) } },
    handleCancel(type, obj) {
      uni.showModal({ title: '提示', content: '确定取消订单吗？', success: async (r) => { if (r.confirm) { await cancelOrder(obj.id); uni.navigateBack() } } })
    },
    async oneMoreOrder(id) { await repetitionOrder(id); uni.switchTab({ url: '/pages/index/index' }) },
    handleReminder(type, id) { reminder(id); uni.showToast({ title: '已催单', icon: 'success' }) },
    handlePay(id) { uni.redirectTo({ url: '/pages/pay/index?orderId=' + id }) },
    handlePhone() { uni.makePhoneCall({ phoneNumber: this.shopPhone }) },
    handleRefund() { uni.showToast({ title: '请联系商家退款', icon: 'none' }) },
    goBack() { uni.navigateBack() },
    runTimeBack(time) {
      if (this.orderDetailsData.status !== 1) return
      const orderTime = new Date(time).getTime()
      const remaining = Math.max(0, Math.floor((orderTime + 15 * 60 * 1000 - Date.now()) / 1000))
      if (remaining <= 0) { this.timeout = true; this.isOvertime = true; return }
      const m = Math.floor(remaining / 60)
      this.rocallTime = `${m < 10 ? '0' + m : m}:${remaining % 60 < 10 ? '0' + remaining % 60 : remaining % 60}`
    }
  },
  onLoad(options) { if (options && options.orderId) this.getBaseData(options.orderId) }
}
</script>
<style scoped>
.order_content { height: 100vh; background: #f6f6f6; }
.order_content_box { height: 100%; padding: 100rpx 18rpx 20rpx; }
.box { background: #fff; border-radius: 8rpx; margin-bottom: 20rpx; }
.orderInfoTip { text-align: center; padding: 30rpx 0 0; }
.orderInfoTip .tit { font-weight: 600; font-size: 36rpx; color: #333; padding-bottom: 12rpx; }
.orderInfoTip .time { padding: 0 0 14rpx; color: #666; font-size: 28rpx; display: flex; align-items: center; justify-content: center; }
.orderInfoTip .time text { color: #f58c21; }
.againBtn { display: flex; justify-content: center; gap: 20rpx; padding: 0 20rpx 26rpx; }
.new_btn { height: 60rpx; line-height: 60rpx; border-radius: 30rpx; background: #f5f5f5; font-size: 26rpx; padding: 0 30rpx; }
.new_btn.btn { background: #ffc200; }
.timeTip { padding: 26rpx; font-size: 26rpx; color: #666; display: flex; align-items: center; }
.timeTip text { color: #ffc200; }
.icon { width: 48rpx; height: 48rpx; margin-right: 12rpx; display: inline-block; }
.newIcon { background: #1dc779; border-radius: 8rpx; }
.moneyIcon { background: #f58c21; border-radius: 8rpx; }
.timeIcon { width: 28rpx; height: 28rpx; display: inline-block; margin-right: 12rpx; background: #f58c21; border-radius: 4rpx; }
.order-type { padding: 40rpx 0 10rpx; }
.type_item { display: flex; margin-bottom: 40rpx; }
.dish_img { width: 90rpx; margin: 0 24rpx 0 20rpx; }
.dish_img_url { width: 90rpx; height: 90rpx; border-radius: 8rpx; display: block; }
.dish_info { position: relative; flex: 1; }
.dish_name { font-size: 26rpx; color: #20232a; }
.dish_price { font-size: 28rpx; color: #818693; margin-top: 10rpx; }
.dish_number { padding: 0 10rpx; }
.dish_active { position: absolute; right: 0; top: 0; font-size: 28rpx; color: #333; font-weight: 600; }
.dish_active text { font-size: 24rpx; }
.iconUp { text-align: center; font-size: 24rpx; color: #666; padding: 8rpx 0; }
.icon_img { width: 30rpx; height: 30rpx; transform: rotate(90deg); vertical-align: middle; }
.icon_imgDown { transform: rotate(-90deg); margin-top: -5px; }
.orderList { margin: 0 20rpx 20rpx; }
.orderInfo { display: flex; font-size: 28rpx; padding: 10rpx 0 16rpx; font-weight: 600; color: #333; }
.orderInfo .text { flex: 1; font-size: 26rpx; font-weight: normal; }
.totalMoney { border-top: 1px solid #efefef; text-align: right; margin-top: 20rpx; padding-top: 20rpx; font-size: 24rpx; }
.totalMoney .text { padding-left: 8rpx; font-weight: 600; color: #333; font-size: 32rpx; }
.totalMoney .text text { font-size: 26rpx; }
.word_text { padding: 16rpx 20rpx; border-bottom: 1px solid #efefef; }
.word_text .word_style { font-size: 28rpx; font-weight: 500; color: #333; }
.contactMerchant { padding: 26rpx; text-align: center; font-size: 28rpx; color: #333; font-weight: 600; }
.phoneIcon { width: 32rpx; height: 32rpx; background: #1dc779; border-radius: 4rpx; display: inline-block; vertical-align: middle; margin-right: 6rpx; }
.orderBaseInfo { padding: 38rpx 24rpx; font-size: 28rpx; color: #333; }
.orderBaseInfo .nameInfo text { padding-right: 12rpx; }
.orderBaseInfo > view { display: flex; padding-bottom: 32rpx; line-height: 40rpx; }
.orderBaseInfo > view:last-child { padding-bottom: 0; }
.orderBaseInfo > view > view:first-child { padding-right: 40rpx; width: 110rpx; color: #666; }
</style>
