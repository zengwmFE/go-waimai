<template>
  <view class="pop_mask" v-show="visible" @tap="closePopup">
    <view class="popup-content" @tap.stop>
      <view class="phoneNum">{{ phoneData }}</view>
      <view class="btn">
        <view @tap="call">呼叫</view>
        <view @tap="closePopup">取消</view>
      </view>
    </view>
  </view>
</template>
<script>
export default {
  props: { phoneData: { type: String, default: '' } },
  data() { return { visible: false } },
  methods: {
    show() { this.visible = true },
    closePopup() { this.visible = false },
    call() { uni.makePhoneCall({ phoneNumber: this.phoneData }); this.closePopup() }
  }
}
</script>
<style scoped>
.pop_mask { position: fixed; top: 0; left: 0; width: 100vw; height: 100vh; background: rgba(0,0,0,0.4); z-index: 9999; display: flex; align-items: center; justify-content: center; }
.popup-content { background: #fff; border-radius: 12px; padding: 30px; width: 280px; text-align: center; }
.phoneNum { font-size: 20px; margin-bottom: 20px; }
.btn { display: flex; justify-content: space-around; }
.btn view { padding: 8px 30px; border-radius: 20px; }
.btn view:first-child { background: #ffc200; color: #333; }
.btn view:last-child { background: #f5f5f5; }
</style>
