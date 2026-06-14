<template>
  <view class="navBar" :style="{ paddingTop: statusBarHeight + 'px', backgroundColor: backgroundColor, color: color }">
    <view class="leftNav" @tap="handleBack">
      <image v-if="leftIcon" class="back" src="@/static/btn_back.png" mode="aspectFit" />
    </view>
    <view class="centerNav">{{ title }}</view>
    <view class="rightNav"><slot name="right" /></view>
  </view>
</template>
<script>
export default {
  props: {
    title: { type: String, default: '' },
    leftIcon: { type: String, default: '' },
    backgroundColor: { type: String, default: '#333333' },
    color: { type: String, default: '#ffffff' }
  },
  computed: {
    statusBarHeight() { return uni.getSystemInfoSync().statusBarHeight }
  },
  methods: {
    handleBack() { this.$emit('clickLeft') }
  }
}
</script>
<style scoped>
.navBar { position: fixed; top: 0; left: 0; width: 100vw; z-index: 999; display: flex; align-items: center; height: 44px; margin-bottom: 30px;}
.leftNav, .rightNav { width: 44px; height: 44px; display: flex; align-items: center; justify-content: center; flex-shrink: 0; }
.leftNav .back { width: 112px; height: 112px; }
.centerNav { flex: 1; text-align: center; font-size: 34rpx; font-weight: 600; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
</style>
