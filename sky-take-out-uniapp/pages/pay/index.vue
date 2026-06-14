<template>
  <view class="customer-box">
    <view class="wrap">
      <view class="contion">
        <view class="orderPay">
          <view>
            <view v-if="timeout">订单已超时</view>
            <view v-else
              >支付剩余时间<text>{{ rocallTime }}</text></view
            >
          </view>
          <view class="money"
            >￥<text>{{ orderDataInfo.orderAmount }}</text></view
          >
          <view>{{ "苍穹餐厅-" + orderDataInfo.orderNumber }}</view>
        </view>
      </view>
      <view class="box payBox">
        <view class="contion">
          <view class="example-body">
            <radio-group class="uni-list" @change="styleChange">
              <view class="uni-list-item">
                <view
                  v-for="(item, index) in payMethodList"
                  :key="index"
                  class="uni-list-item__container"
                >
                  <view class="uni-list-item__content">
                    <view class="wechatIcon" />
                    <text class="uni-list-item__content-title">{{ item }}</text>
                  </view>
                  <view class="uni-list-item__extra">
                    <radio
                      class="radioIcon"
                      :value="item"
                      color="#FFC200"
                      :checked="index === activeRadio"
                    />
                  </view>
                </view>
              </view>
            </radio-group>
          </view>
        </view>
      </view>
      <view class="bottomBox btnBox">
        <button class="add_btn" type="primary" plain @tap.stop="handleSave">
          确认支付
        </button>
      </view>
    </view>
  </view>
</template>
<script>
import { mapState } from "vuex";
import { payment } from "@/utils/api";
export default {
  data() {
    return {
      timeout: false,
      rocallTime: "",
      orderId: null,
      orderDataInfo: {},
      activeRadio: 0,
      payMethodList: ["微信支付"],
      times: null,
    };
  },
  computed: { ...mapState(["orderData"]) },
  methods: {
    async handleSave() {
      try {
        const res = await payment({
          orderNumber: this.orderDataInfo.orderNumber,
          payMethod: 1,
        });
        if (res.code === 1) {
          uni.redirectTo({
            url: "/pages/success/index?orderId=" + this.orderId,
          });
        }
      } catch (e) {
        console.error(e);
      }
    },
    runTimeBack() {
      this.times = setInterval(() => {
        const t = this.orderDataInfo.orderTime;
        if (!t) return;
        const orderTime = new Date(t).getTime();
        const remaining = Math.max(
          0,
          Math.floor((orderTime + 15 * 60 * 1000 - Date.now()) / 1000),
        );
        if (remaining <= 0) {
          this.timeout = true;
          clearInterval(this.times);
          return;
        }
        const m = Math.floor(remaining / 60);
        const s = remaining % 60;
        this.rocallTime = `${m < 10 ? "0" + m : m}:${s < 10 ? "0" + s : s}`;
      }, 1000);
    },
  },
  created() {
    console.log(this.orderData, 'this.orderData')
    if (this.orderData) this.orderDataInfo = this.orderData;
  },
  mounted() {
    this.runTimeBack();
  },
  onLoad(options) {
    if (options && options.orderId) this.orderId = options.orderId;
  },
  beforeDestroy() {
    if (this.times) clearInterval(this.times);
  },
};
</script>
<style scoped>
.customer-box {
  height: 100vh;
  background: #f6f6f6;
}
.wrap {
  width: 730rpx;
  margin: 0 auto;
  padding-top: 60rpx;
}
.contion {
  padding: 0 20rpx;
}
.orderPay {
  text-align: center;
  font-size: 26rpx;
  line-height: 36rpx;
  color: #666;
  padding: 98rpx 0 60rpx;
}
.orderPay > view {
  padding: 6rpx 0;
}
.orderPay .money {
  font-size: 36rpx;
  line-height: 60rpx;
  color: #333;
  padding-top: 24rpx;
}
.orderPay .money text {
  font-size: 70rpx;
  font-weight: 600;
}
.box {
  background: #fff;
  border-radius: 12rpx;
  margin-bottom: 20rpx;
}
.example-body {
  display: flex;
  padding: 0;
}
.uni-list {
  flex: 1;
}
.uni-list-item {
  display: flex;
  flex: 1;
  background: #fff;
}
.uni-list-item__container {
  padding: 16rpx 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.uni-list-item__content {
  display: flex;
  align-items: center;
}
.uni-list-item__content-title {
  font-size: 28rpx;
  color: #333;
}
.uni-list-item__extra {
  padding-right: 20rpx;
}
.wechatIcon {
  width: 48rpx;
  height: 48rpx;
  background: #1dc779;
  border-radius: 8rpx;
  margin-right: 20rpx;
  display: inline-block;
}
.bottomBox {
  position: fixed;
  bottom: 40rpx;
  left: 20rpx;
  right: 20rpx;
}
.add_btn {
  width: 100%;
  height: 86rpx;
  line-height: 86rpx;
  border-radius: 8rpx;
  background: #ffc200;
  border: 1px solid #ffc200;
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
