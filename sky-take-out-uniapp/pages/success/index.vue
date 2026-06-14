<template>
  <view class="success_content">
    <view class="success_info">
      <image class="success_icon" src="@/static/success.png" mode="aspectFit" />
      <view class="success_title">下单成功</view>
      <view class="word-box">
        <text class="word_bottom">预计<text class="word_date">{{ arrivalTime }}</text>送达</text>
      </view>
      <view class="success_desc">后厨已开始疯狂备餐中, 请耐心等待~</view>
      <view class="btns">
        <view class="go_dish defaultBtn" @tap.stop="goIndex">返回首页</view>
        <view class="go_dish" @tap.stop="goOrder">查看订单</view>
      </view>
    </view>
  </view>
</template>
<script>
import { mapState } from 'vuex'
export default {
  data() { return { arrivalTime: '', orderId: null } },
  computed: { ...mapState(['shopInfo', 'arrivals']) },
  methods: {
    goIndex() { uni.switchTab({ url: '/pages/index/index' }) },
    goOrder() { uni.navigateTo({ url: '/pages/historyOrder/historyOrder' }) },
    getHarfAnOur() { this.arrivalTime = this.arrivals[0]?.time || '立即送出' }
  },
  onLoad(options) {
    this.getHarfAnOur()
    if (options && options.orderId) this.orderId = options.orderId
  }
}
</script>
<style scoped>
.success_content { height: 100vh; display: flex; align-items: center; justify-content: center; background: #fff; }
.success_info { text-align: center; }
.success_icon { width: 160rpx; height: 160rpx; margin-bottom: 30rpx; }
.success_title { font-size: 40rpx; font-weight: 600; color: #333; margin-bottom: 20rpx; }
.word-box { margin-bottom: 16rpx; }
.word_bottom { font-size: 28rpx; color: #666; }
.word_date { color: #f58c21; }
.success_desc { font-size: 26rpx; color: #999; margin-bottom: 60rpx; }
.btns { display: flex; justify-content: center; gap: 30rpx; }
.go_dish { width: 220rpx; height: 80rpx; line-height: 80rpx; background: #ffc200; border-radius: 40rpx; color: #333; font-size: 28rpx; text-align: center; }
.defaultBtn { background: #f5f5f5; color: #666; }
</style>
