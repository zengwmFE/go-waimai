<template>
  <view class="nonet_content">
    <view class="success_info">
      <image class="success_icon" src="@/static/noNet.png" mode="aspectFit" />
      <view class="success_title">网络无法连接</view>
      <view class="go_dish" @tap.stop="goIndex">点击刷新</view>
    </view>
  </view>
</template>
<script>
import { mapState } from 'vuex'
export default {
  computed: { ...mapState(['shopInfo']) },
  methods: { goIndex() { uni.navigateTo({ url: '/pages/index/index' }) } }
}
</script>
<style scoped>
.nonet_content { height: 100vh; display: flex; align-items: center; justify-content: center; background: #fff; }
.success_info { text-align: center; }
.success_icon { width: 240rpx; height: 240rpx; margin-bottom: 40rpx; }
.success_title { font-size: 32rpx; color: #333; margin-bottom: 40rpx; }
.go_dish { width: 300rpx; height: 80rpx; line-height: 80rpx; background: #ffc200; border-radius: 40rpx; color: #333; font-size: 28rpx; text-align: center; margin: 0 auto; }
</style>
