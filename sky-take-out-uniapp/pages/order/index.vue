<template>
  <view class="order_content">
    <nav-bar
      title="提交订单"
      leftIcon="arrowleft"
      backgroundColor="#333333"
      @clickLeft="goBack"
    />
    <scroll-view :enhanced="true" class="order_content_box" scroll-y>
      <view class="new_address">
        <view class="top" @tap.stop="goAddress">
          <view v-if="!address" class="address_name_disabled"
            >请选择收货地址</view
          >
          <view v-else class="address_name">
            <view class="address">
              <text :class="['tag', 'tag' + tagLabel]">{{ addressLabel }}</text>
              <text class="word">{{ address }}</text>
            </view>
            <view class="name"
              ><text class="name_1">{{ nickName }}</text
              ><text class="name_2">{{ phoneNumber }}</text></view
            >
            <view v-if="address" class="infoTip"
              >为减少接触，降低风险，推荐使用无接触配送</view
            >
          </view>
          <view class="address_image"><view class="to_right" /></view>
        </view>
        <view class="bottom">
          <view class="bottomTime _div" @tap.stop="openTimePopuo('bottom')">
            <text class="time_name_disabled">立即送出</text>
            <view class="address_image"
              ><text>{{ arrivalTime }}送达</text><view class="to_right"
            /></view>
          </view>
          <view v-if="address" class="infoTip"
            >因配送订单较多，送达时间可能波动</view
          >
        </view>
      </view>
      <view class="order_list_cont">
        <view class="box order_list">
          <view class="word_text"
            ><text class="word_style">苍穹餐厅</text></view
          >
          <view class="order-type">
            <view
              v-for="(obj, index) in orderDataes"
              :key="index"
              class="type_item"
            >
              <view class="dish_img"
                ><image class="dish_img_url" mode="aspectFill" :src="obj.image"
              /></view>
              <view class="dish_info">
                <view class="dish_name">{{ obj.name }}</view>
                <view v-if="obj.dishFlavor" class="dish_dishFlavor">{{
                  obj.dishFlavor
                }}</view>
                <view class="dish_price"
                  >×<text class="dish_number">{{ obj.number }}</text></view
                >
                <view class="dish_active"
                  ><text>￥</text
                  >{{ (obj.price * obj.number).toFixed(2) }}</view
                >
              </view>
            </view>
            <view
              v-if="orderListDataes.length > 3"
              class="iconUp"
              @tap.stop="showDisplay = !showDisplay"
            >
              <text>{{ showDisplay ? "点击收起" : "展开更多" }}</text>
              <image
                :class="['icon_img', showDisplay ? 'icon_imgDown' : '']"
                src="@/static/toRight.png"
              />
            </view>
            <view class="orderList">
              <view class="orderInfo"
                ><text class="text">打包费</text><text class="may">￥</text
                >{{ packAmount || 0 }}</view
              >
              <view class="orderInfo"
                ><text class="text">配送费</text
                ><text class="may">￥</text>6</view
              >
              <view class="totalMoney"
                >合计<text class="text"
                  ><text>￥</text
                  >{{
                    (orderDishPrice + 6 + (packAmount || 0)).toFixed(2)
                  }}</text
                ></view
              >
            </view>
          </view>
        </view>
      </view>
      <view class="boxPad">
        <view class="box order_list">
          <view class="uniInfo">
            <view @tap.stop="goRemark">
              <uni-list
                ><uni-list-item showArrow title="备注"
                  ><text class="temarkText" slot="footer">{{
                    remark || "推荐使用无接触配送"
                  }}</text></uni-list-item
                ></uni-list
              >
            </view>
            <view @tap.stop="openPopuos('bottom')">
              <uni-list
                ><uni-list-item showArrow title="餐具数量"
                  ><text slot="footer"
                    >已在店选择：{{ tablewareData }}</text
                  ></uni-list-item
                ></uni-list
              >
            </view>
            <view class="invoiceBox">
              <uni-list
                ><uni-list-item title="发票"
                  ><text slot="footer">请联系商家提供</text></uni-list-item
                ></uni-list
              >
            </view>
          </view>
        </view>
      </view>
      <!-- Tableware Popup -->
      <uni-popup ref="popup" type="bottom">
        <view class="popup-content">
          <view class="popupTitle"
            ><text>按政府条例要求：</text
            ><text
              >商家不得主动向您提供一次性餐具，请按需选择餐具数量</text
            ></view
          >
          <view class="popupCon">
            <view class="popupBtn"
              ><text @tap.stop="closePopup">取消</text><text>选择本单餐具</text
              ><text @tap.stop="handlePiker">确定</text></view
            >
            <view class="pickerBox">
              <picker-view
                :value="[0]"
                @change="changeCont"
                style="height: 300rpx"
              >
                <picker-view-column>
                  <view v-for="(item, i) in baseData" :key="i" class="item">{{
                    item
                  }}</view>
                </picker-view-column>
              </picker-view>
            </view>
          </view>
          <view class="popupSet"
            ><view>后续订单餐具设置</view>
            <radio-group @change="handleRadio">
              <label v-for="(item, i) in radioGroup" :key="i"
                ><radio
                  :value="item"
                  color="#FFC200"
                  :checked="item === activeRadio"
                />{{ item }}</label
              >
            </radio-group>
          </view>
        </view>
      </uni-popup>
      <!-- Time Picker Popup -->
      <uni-popup ref="timePopup" type="bottom">
        <view class="popup-content">
          <view class="pickerCon">
            <view class="dayBox">
              <scroll-view
                :enhanced="true"
                scroll-x
                :scroll-into-view="scrollinto"
                :scroll-with-animation="true"
              >
                <view
                  v-for="(item, index) in popleft"
                  :key="index"
                  class="scroll-row-item"
                  :id="'tab' + index"
                  @tap.stop="dateChange(index)"
                >
                  <view :class="tabIndex === index ? 'scroll-row-day' : ''"
                    >{{ item
                    }}<text class="week"
                      >({{
                        ["日", "一", "二", "三", "四", "五", "六"][
                          new Date(Date.now() + index * 86400000).getDay()
                        ]
                      }})</text
                    ></view
                  >
                </view>
              </scroll-view>
            </view>
            <view class="timeBox">
              <scroll-view :enhanced="true" scroll-y style="height: 100%">
                <view
                  v-for="(val, i) in newDateData"
                  :key="i"
                  :class="[
                    'item',
                    selectValue === i ? 'city-column_select' : '',
                  ]"
                  @tap.stop="timeClick(val, i)"
                  >{{ val }}</view
                >
              </scroll-view>
            </view>
          </view>
          <view class="btns" @tap.stop="closeTimePopup">取消</view>
        </view>
      </uni-popup>
    </scroll-view>
    <view class="footer_order_buttom order_form">
      <view class="order_number">
        <image class="order_number_icon" src="@/static/btn_waiter_sel.png" />
        <view class="order_dish_num">{{ orderDishNumber }}</view>
      </view>
      <view class="order_price"
        ><text class="ico">￥</text
        >{{ (orderDishPrice + 6 + (packAmount || 0)).toFixed(2) }}</view
      >
      <view class="order_but">
        <view v-if="isHandlePy" class="order_but_rit">去支付</view>
        <view v-else class="order_but_rit" @tap.stop="payOrderHandle">去支付</view>
      </view>
    </view>
  </view>
</template>
<script>
import { getAddressBookDefault, submitOrder } from "@/utils/api";
import { mapState, mapMutations } from "vuex";
export default {
  components: { "nav-bar": () => import("@/components/nav-bar/nav-bar.vue") },
  data() {
    return {
      orderDishPrice: 0,
      nickName: "",
      phoneNumber: "",
      address: "",
      gender: "0",
      remark: "",
      arrivalTime: "立即送出",
      addressBookId: "",
      addressLabel: "",
      tagLabel: "",
      orderDishNumber: 0,
      showDisplay: false,
      tablewareData: "无需餐具",
      packAmount: 0,
      isHandlePy: false,
      popleft: ["今天", "明天"],
      tabIndex: 0,
      scrollinto: "tab0",
      baseData: ["无需餐具", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"],
      activeRadio: "无需餐具",
      radioGroup: ["依据餐量提供", "无需餐具"],
      newDateData: [],
      selectValue: 0,
    };
  },
  computed: {
    ...mapState([
      "shopInfo",
      "orderListData",
      "arrivals",
      "remarkData",
      "addressData",
    ]),
    orderListDataes() {
      return this.orderListData;
    },
    orderDataes() {
      console.log(this.orderListData, "this.orderListData");
      return this.showDisplay
        ? this.orderListData
        : this.orderListData.slice(0, 3);
    },
  },
  methods: {
    ...mapMutations([
      "setAddressBackUrl",
      "setOrderData",
      "setArrivalTime",
      "setRemark",
    ]),
    goBack() {
      uni.navigateBack();
    },
    goAddress() {
      this.setAddressBackUrl("/pages/order/index");
      uni.navigateTo({ url: "/pages/address/address?form=order" });
    },
    goRemark() {
      uni.navigateTo({ url: "/pages/remark/index" });
    },
    async getAddressBookDefault() {
      try {
        const res = await getAddressBookDefault();
        if (res.code === 1 && res.data) {
          const addr = res.data;
          this.address = `${addr.provinceName}${addr.cityName}${addr.districtName}${addr.detail}`;
          this.nickName = addr.consignee;
          this.phoneNumber = addr.phone;
          this.addressBookId = addr.id;
          this.addressLabel =
            addr.label === "1"
              ? "公司"
              : addr.label === "2"
              ? "家"
              : addr.label === "3"
              ? "学校"
              : "";
          this.tagLabel = addr.label;
        }
      } catch (e) {
        console.error(e);
      }
    },
    nowString() {
      const d = new Date();
      const pad = (n) => String(n).padStart(2, "0");
      return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(
        d.getDate(),
      )} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
    },
    async payOrderHandle() {
      if (!this.addressBookId) {
        uni.showToast({ title: "请选择收货地址", icon: "none" });
        return;
      }
      this.isHandlePy = true;
      try {
        const res = await submitOrder({
          addressBookId: this.addressBookId,
          amount: this.orderDishPrice + 6 + this.packAmount,
          remark: this.remark,
          tablewareNumber:
            this.tablewareData === "无需餐具" ? 0 : Number(this.tablewareData),
          tablewareStatus: this.tablewareData === "无需餐具" ? 0 : 1,
          deliveryStatus: this.arrivalTime === "立即送出" ? 1 : 2,
          estimatedDeliveryTime:
            this.arrivalTime === "立即送出"
              ? this.nowString()
              : this.arrivalTime,
        });
        if (res.code === 1) {
          this.setOrderData({
            orderAmount: this.orderDishPrice + 6 + this.packAmount,
            orderNumber: res.data.number,
            orderId: res.data.id,
            orderTime: res.data.orderTime,
          });
          uni.redirectTo({ url: "/pages/pay/index" });
        }
      } catch (e) {
        console.error(e);
      }
      this.isHandlePy = false;
    },
    computOrderInfo() {
      let num = 0,
        price = 0;
      this.orderListData.forEach((d) => {
        num += Number(d.number);
        price += Number(d.number) * Number(d.price);
      });
      this.orderDishNumber = num;
      this.orderDishPrice = price;
    },
    openPopuos() {
      this.$refs.popup && this.$refs.popup.open();
    },
    closePopup() {
      this.$refs.popup && this.$refs.popup.close();
    },
    handlePiker() {
      this.$refs.popup && this.$refs.popup.close();
    },
    changeCont(e) {
      const idx = e.detail.value[0];
      this.tablewareData = this.baseData[idx];
    },
    handleRadio(e) {
      this.activeRadio = e.detail.value;
    },
    openTimePopuo() {
      this.$refs.timePopup && this.$refs.timePopup.open();
    },
    closeTimePopup() {
      this.$refs.timePopup && this.$refs.timePopup.close();
    },
    dateChange(index) {
      this.tabIndex = index;
      this.selectValue = 0;
      if (index === 0) {
        this.newDateData = this.buildTimeSlots(0);
      } else {
        this.newDateData = this.buildTimeSlots(1);
      }
    },
    buildTimeSlots(dayOffset) {
      const now = new Date();
      const base = new Date(
        now.getFullYear(),
        now.getMonth(),
        now.getDate() + dayOffset,
      );
      const hours = dayOffset === 0 ? now.getHours() : 0;
      const slots = ["立即派送"];
      for (let h = Math.max(hours, 8); h < 22; h++) {
        for (let m of ["00", "30"]) {
          const t = `${String(h).padStart(2, "0")}:${m}`;
          if (dayOffset === 0) {
            const slotTime = new Date(
              base.getFullYear(),
              base.getMonth(),
              base.getDate(),
              h,
              parseInt(m),
            );
            if (slotTime > now) slots.push(t);
          } else {
            slots.push(t);
          }
        }
      }
      return slots;
    },
    timeClick(val, i) {
      this.selectValue = i;
      if (val === "立即派送") {
        this.arrivalTime = "立即送出";
      } else {
        const now = new Date();
        const date = new Date(
          now.getFullYear(),
          now.getMonth(),
          now.getDate() + this.tabIndex,
        );
        const y = date.getFullYear();
        const mo = String(date.getMonth() + 1).padStart(2, "0");
        const d = String(date.getDate()).padStart(2, "0");
        this.arrivalTime = `${y}-${mo}-${d} ${val}:00`;
      }
      this.closeTimePopup();
    },
    init() {
      this.newDateData = this.buildTimeSlots(0);
      this.computOrderInfo();
      const info = this.$store.state.baseUserInfo;
      if (info) {
        this.nickName = info.nickName || "";
        this.gender = info.gender || "0";
      }
      this.phoneNumber = this.$store.state.shopPhone || "";
      this.remark = this.remarkData || "";
      this.getAddressBookDefault();
    },
  },
  onLoad() {
    this.init();
  },
  onShow() {
    this.init();
  },
};
</script>
<style scoped>
.order_content {
  height: 100vh;
  background: #f6f6f6;
}
.order_content_box {
  height: 100%;
  padding: 100rpx 18rpx 20rpx;
}
.box {
  background: #fff;
  border-radius: 8rpx;
  margin-bottom: 20rpx;
}
.new_address {
  background: #fff;
  border-radius: 8rpx;
  margin-bottom: 20rpx;
  padding-top: 32rpx;
}
.new_address .top {
  margin: 0 20rpx 10rpx;
  display: flex;
  position: relative;
  align-items: center;
}
.address_name {
  flex: 1;
  overflow: hidden;
}
.address_name .address {
  height: 50rpx;
  line-height: 50rpx;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  position: relative;
  padding-left: 76rpx;
}
.address .tag {
  position: absolute;
  left: 0;
  top: 8rpx;
  padding: 2rpx 8rpx;
  border-radius: 4rpx;
  font-size: 18rpx;
  color: #fff;
}
.tag1 {
  background: #f58c21;
}
.tag2 {
  background: #1dc779;
}
.tag3 {
  background: #4a90d9;
}
.address .word {
  font-size: 34rpx;
  font-weight: 550;
  color: #20232a;
}
.address_name .name {
  margin: 20rpx 0 10rpx;
  color: #666;
}
.name_1,
.name_2 {
  font-size: 26rpx;
  color: #333;
}
.name_2 {
  margin-left: 10rpx;
}
.address_image {
  position: absolute;
  right: 0;
  top: 8rpx;
  display: flex;
  align-items: center;
}
.address_image text {
  font-size: 24rpx;
  color: #f58c21;
}
.to_right {
  width: 32rpx;
  height: 32rpx;
  background: #ccc;
  border-radius: 4rpx;
  margin-left: 10rpx;
}
.address_name_disabled {
  flex: 1;
  font-size: 36rpx;
  color: #f58c21;
}
.bottom {
  margin: 18rpx 20rpx 28rpx;
}
.bottomTime {
  display: flex;
}
.time_name_disabled {
  flex: 1;
  font-size: 28rpx;
  font-weight: 600;
  color: #333;
}
.infoTip {
  color: #f58c21;
  font-size: 24rpx;
  padding-right: 30rpx;
  margin-top: 10rpx;
}
.order_list .order-type {
  padding: 40rpx 0 10rpx;
}
.type_item {
  display: flex;
  margin-bottom: 40rpx;
}
.dish_img {
  width: 90rpx;
  margin: 0 24rpx 0 20rpx;
}
.dish_img_url {
  width: 90rpx;
  height: 90rpx;
  border-radius: 8rpx;
  display: block;
}
.dish_info {
  position: relative;
  flex: 1;
}
.dish_name {
  font-size: 26rpx;
  color: #20232a;
}
.dish_price {
  font-size: 28rpx;
  color: #818693;
  margin-top: 10rpx;
}
.dish_number {
  padding: 0 10rpx;
}
.dish_active {
  position: absolute;
  right: 0;
  top: 0;
  font-size: 28rpx;
  color: #333;
  font-weight: 600;
  display: flex;
  align-items: center;
}
.dish_active text {
  font-size: 24rpx;
}
.iconUp {
  text-align: center;
  font-size: 24rpx;
  color: #666;
  padding: 8rpx 0;
}
.icon_img {
  width: 30rpx;
  height: 30rpx;
  transform: rotate(90deg);
  vertical-align: middle;
}
.icon_imgDown {
  transform: rotate(-90deg);
  margin-top: -5px;
}
.orderList {
  margin: 0 20rpx 20rpx;
}
.orderInfo {
  display: flex;
  font-size: 28rpx;
  padding: 10rpx 0 16rpx;
  font-weight: 600;
  color: #333;
}
.orderInfo .text {
  flex: 1;
  font-size: 26rpx;
  font-weight: normal;
}
.totalMoney {
  border-top: 1px solid #efefef;
  text-align: right;
  margin-top: 20rpx;
  padding-top: 20rpx;
  font-size: 24rpx;
}
.totalMoney .text {
  padding-left: 8rpx;
  font-weight: 600;
  color: #333;
  font-size: 32rpx;
}
.totalMoney .text text {
  font-size: 26rpx;
}
.word_text {
  padding: 16rpx 20rpx;
  border-bottom: 1px solid #efefef;
}
.word_text .word_style {
  font-size: 28rpx;
  font-weight: 500;
  color: #333;
}
.boxPad {
  padding-bottom: 170rpx;
}
.uniInfo {
  padding: 20rpx;
}
.temarkText {
  width: 480rpx;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-align: right;
  font-size: 28rpx;
  color: #666;
}
.uniInfo ::v-deep .uni-list-item__content-title {
  font-size: 30rpx;
}
.uniInfo ::v-deep .uni-list-item__extra {
  font-size: 28rpx;
}
.footer_order_buttom {
  position: fixed;
  bottom: 48rpx;
  padding-bottom: env(safe-area-inset-bottom);
  left: 30rpx;
  right: 30rpx;
  display: flex;
  height: 88rpx;
  background: rgba(0, 0, 0, 0.9);
  border-radius: 50rpx;
  z-index: 99;
  padding: 0 10rpx;
}
.order_number {
  position: relative;
  width: 120rpx;
}
.order_number_icon {
  position: absolute;
  width: 120rpx;
  height: 118rpx;
  left: 12rpx;
  bottom: 0;
  display: block;
}
.order_dish_num {
  position: absolute;
  z-index: 9;
  min-width: 12rpx;
  height: 36rpx;
  line-height: 36rpx;
  padding: 0 12rpx;
  left: 92rpx;
  font-size: 24rpx;
  top: -8rpx;
  border-radius: 20rpx;
  background: #e94e3c;
  color: #fff;
  font-weight: 500;
}
.order_price {
  flex: 1;
  text-align: left;
  color: #fff;
  line-height: 88rpx;
  padding-left: 34rpx;
  font-size: 36rpx;
  font-weight: bold;
}
.order_price .ico {
  font-size: 24rpx;
}
.order_but {
  width: 200rpx;
  height: 72rpx;
  line-height: 72rpx;
  border-radius: 72rpx;
  text-align: center;
  margin-top: 8rpx;
}
.order_but_rit {
  width: 100%;
  border-radius: 72rpx;
  background: #ffc200;
  font-size: 30rpx;
  font-weight: 500;
  color: #333;
  height: 72rpx;
  line-height: 72rpx;
}
.popup-content {
  padding: 30rpx;
  background-color: #fff;
}
.popupTitle {
  font-size: 24rpx;
  color: #666;
  margin-bottom: 20rpx;
  line-height: 36rpx;
}
.pickerBox {
  height: 300rpx;
}
.pickerBox .item {
  text-align: center;
  padding: 20rpx;
}
.popupBtn {
  display: flex;
  justify-content: space-between;
  font-size: 28rpx;
  margin-bottom: 20rpx;
}
.popupBtn text:first-child,
.popupBtn text:last-child {
  color: #f58c21;
}
.popupSet {
  margin-top: 20rpx;
  font-size: 26rpx;
}
.popupSet label {
  display: block;
  margin: 10rpx 0;
}
.pickerCon {
  display: flex;
  height: 30vh;
  color: #666;
  font-size: 28rpx;
}
.dayBox {
  background: #f6f6f6;
  width: 45%;
}
.dayBox .scroll-row-item {
  height: 88rpx;
  line-height: 88rpx;
  text-align: center;
}
.scroll-row-day {
  background: #fff;
  color: #333;
}
.week {
  padding-left: 20rpx;
}
.timeBox {
  flex: 1;
  height: 100%;
  overflow: hidden;
}
.timeBox .item {
  padding: 0 22rpx 0 30rpx;
  height: 88rpx;
  line-height: 88rpx;
}
.city-column_select {
  color: #f5932f;
}
.btns {
  width: 100%;
  height: 98rpx;
  line-height: 98rpx;
  background: #fff;
  font-size: 28rpx;
  color: #333;
  text-align: center;
  border-top: 2rpx solid #efefef;
}
</style>
