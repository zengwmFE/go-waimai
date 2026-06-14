<template>
  <view class="customer-box">
    <nav-bar title="订单备注" leftIcon="arrowleft" backgroundColor="#333333" @clickLeft="goBack" />
    <view class="wrap">
      <view class="box">
        <view class="contion">
          <view class="order_list">
            <view class="uni-textarea">
              <textarea
                class="beizhu_text"
                placeholder="无接触配送，将商品挂家门口或放前台，地址封闭管理时请电话联系"
                v-model="remark"
                @input="monitorInput"
                :maxlength="50"
              />
              <text class="numText">
                <text :class="numVal === 0 ? 'tip' : ''">{{ numVal }}</text>/50
              </text>
            </view>
          </view>
        </view>
      </view>
      <view class="btnBox">
        <button class="add_btn" type="primary" plain @tap.stop="handleSaveRemark">完成</button>
      </view>
    </view>
  </view>
</template>
<script>
import { mapMutations, mapState } from 'vuex'
export default {
  data() { return { remark: '', numVal: 0 } },
  computed: {
    ...mapState(['remarkData']),
    getVal() { const v = this.remark || ''; return v.length <= 50 ? v : v.substring(0, 50) }
  },
  methods: {
    ...mapMutations(['setRemark']),
    goBack() { uni.navigateBack() },
    handleSaveRemark() {
      this.setRemark(this.remark)
      uni.navigateBack()
    },
    monitorInput(e) { this.numVal = (e.detail.value || '').length },
    validateTextLength(value) { return value ? value.substring(0, 50) : '' }
  },
  onLoad() {
    const saved = this.remarkData
    this.remark = saved || ''
    this.numVal = (saved || '').length
  }
}
</script>
<style scoped>
.customer-box { height: 100vh; background: #f6f6f6; }
.wrap { width: 730rpx; margin: 0 auto; padding-top: 120rpx; }
.box { background: #fff; border-radius: 12rpx; }
.contion { padding: 20rpx; }
.uni-textarea { position: relative; }
.beizhu_text { width: 100%; height: 320rpx; line-height: 60rpx; font-size: 26rpx; background: #fff; border-radius: 12rpx; padding: 20rpx; box-sizing: border-box; }
.numText { position: absolute; right: 20rpx; bottom: 20rpx; font-size: 26rpx; color: #333; }
.numText .tip { color: #bdbdbd; }
.btnBox { margin-top: 36rpx; }
.add_btn { width: 100%; height: 86rpx; line-height: 86rpx; border-radius: 8rpx; background: #ffc200; border: 1px solid #ffc200; font-size: 30rpx; font-weight: 600; color: #333; display: flex; align-items: center; justify-content: center; }
</style>
