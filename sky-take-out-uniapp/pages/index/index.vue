<template>
  <view class="home_content">
    <view class="restaurant_info_box">
      <view class="restaurant_info">
        <view class="info_top">
          <view class="info_top_left">
            <image class="logo_ruiji" src="@/static/logo_ruiji.png" />
          </view>
          <view class="info_top_right">
            <view class="right_title">
              <text>苍穹外卖</text>
              <view v-if="shopStatus === 1" class="businessStatus">营业中</view>
              <view v-else class="businessStatus close">休息中</view>
            </view>
            <view class="right_details">
              <view class="details_flex"
                ><image class="top_icon" src="@/static/length.png" /><text
                  class="icon_text"
                  >距离1.5km</text
                ></view
              >
              <view class="details_flex"
                ><image class="top_icon" src="@/static/money.png" /><text
                  class="icon_text"
                  >配送费6元</text
                ></view
              >
              <view class="details_flex test"
                ><image class="top_icon" src="@/static/time.png" /><text
                  class="icon_text"
                  >预计时长12min</text
                ></view
              >
            </view>
          </view>
        </view>
        <view class="info_bottom">
          <view>
            <view class="word">苍穹餐厅为顾客打造专业的大众化美食外送餐饮</view>
            <view class="address"
              ><text>北京市朝阳区新街大道一号楼8层</text></view
            >
          </view>
          <view
            ><view class="phone" @tap.stop="handlePhone('bottom')"
              ><text class="phoneIcon" /></view
          ></view>
        </view>
      </view>
    </view>
    <view class="restaurant_menu_list">
      <view class="type_list">
        <scroll-view
          :enhanced="true"
          class="menu-scroll-view"
          scroll-y
          :scroll-with-animation="true"
          :scroll-top="scrollTop + 100"
          :scroll-into-view="itemId"
        >
          <view
            v-for="(item, index) in typeListData"
            :key="index"
            :class="['type_item', typeIndex == index ? 'active' : '']"
            :id="'target'"
            @tap.stop="swichMenu(index)"
          >
            <view :class="['item', item.name.length > 5 ? 'allLine' : '']">{{
              item.name
            }}</view>
          </view>
          <view class="seize_seat" />
        </scroll-view>
      </view>
      <scroll-view
        :enhanced="true"
        v-if="dishListItems && dishListItems.length > 0"
        class="vegetable_order_list"
        scroll-y
      >
        <view
          v-for="(item, index) in dishListItems"
          :key="index"
          class="type_item"
        >
          <view class="dish_img" @tap.stop="openDetailHandle(item)">
            <image
              class="dish_img_url"
              mode="aspectFill"
              :src="getNewImage(item.image)"
            />
          </view>
          <view class="dish_info">
            <view class="dish_name" @tap.stop="openDetailHandle(item)">{{
              item.name
            }}</view>
            <view class="dish_label" @tap.stop="openDetailHandle(item)">{{
              item.description || item.name
            }}</view>
            <view class="dish_label" @tap.stop="openDetailHandle(item)"
              >月销量0</view
            >
            <view class="dish_price"
              ><text class="ico">￥</text>{{ item.price.toFixed(2) }}</view
            >
            <view
              v-if="!item.flavors || item.flavors.length === 0"
              class="dish_active"
            >
              <image
                v-if="item.dishNumber >= 1"
                class="dish_red"
                src="@/static/btn_red.png"
                @tap.stop="redDishAction('dish', '普通', item)"
              />
              <text v-if="item.dishNumber > 0" class="dish_number">{{
                item.dishNumber
              }}</text>
              <image
                class="dish_add"
                src="@/static/btn_add.png"
                @tap.stop="addDishAction('dish', '普通', item)"
              />
            </view>
            <view v-else class="dish_active_btn">
              <view class="check_but" @tap.stop="moreNormDataesHandle(item)"
                >选择规格</view
              >
            </view>
          </view>
        </view>
        <view class="seize_seat" />
      </scroll-view>
      <view v-else class="no_dish"
        ><view v-if="typeListData.length > 0">该分类下暂无菜品</view></view
      >
    </view>
    <!-- Bottom Cart Bar -->
    <view v-if="orderListData.length === 0" class="footer_order_buttom">
      <view class="order_number"
        ><image
          class="order_number_icon"
          src="@/static/btn_waiter_nor.png"
          mode="aspectFit"
      /></view>
      <view class="order_price"><text class="ico">￥</text>0</view>
      <view class="order_but">去结算</view>
    </view>
    <view v-else class="footer_order_buttom order_form">
      <view class="orderCar" @tap.stop="openOrderCartList = !openOrderCartList">
        <view class="order_number">
          <image
            class="order_number_icon"
            src="@/static/btn_waiter_sel.png"
            mode="aspectFit"
          />
          <view class="order_dish_num">{{ orderDishNumber }}</view>
        </view>
        <view class="order_price"
          ><text class="ico">￥</text>{{ orderDishPrice.toFixed(2) }}</view
        >
      </view>
      <view class="order_but" @tap.stop="goOrder">去结算</view>
    </view>
    <!-- Dish Detail Popup -->
    <view v-show="openDetailPop" class="pop_mask">
      <view v-if="dishDetailes.type == 1" class="dish_detail_pop">
        <image
          class="div_big_image"
          mode="aspectFill"
          :src="dishDetailes.image"
        />
        <view class="title">{{ dishDetailes.name }}</view>
        <view class="desc">{{ dishDetailes.description }}</view>
        <view class="but_item">
          <view class="price"
            ><text class="ico">￥</text
            >{{ dishDetailes.price.toFixed(2) }}</view
          >
          <view
            v-if="
              dishDetailes.flavors &&
              dishDetailes.flavors.length === 0 &&
              dishDetailes.dishNumber > 0
            "
            class="active"
          >
            <image
              class="dish_red"
              src="@/static/btn_red.png"
              @tap.stop="redDishAction('detail', '普通', dishDetailes)"
            />
            <text class="dish_number">{{ dishDetailes.dishNumber }}</text>
            <image
              class="dish_add"
              src="@/static/btn_add.png"
              @tap.stop="addDishAction('detail', '普通', dishDetailes)"
            />
          </view>
          <view
            v-if="dishDetailes.flavors && dishDetailes.flavors.length > 0"
            class="active"
          >
            <view
              class="dish_card_add"
              @tap.stop="moreNormDataesHandle(dishDetailes)"
              >选择规格</view
            >
          </view>
          <view
            v-if="
              dishDetailes.dishNumber === 0 &&
              (!dishDetailes.flavors || dishDetailes.flavors.length === 0)
            "
            class="active"
          >
            <view
              class="dish_card_add"
              @tap.stop="addDishAction('detail', '普通', dishDetailes)"
              >加入购物车</view
            >
          </view>
        </view>
        <view class="close" @tap.stop="openDetailPop = false"
          ><image class="close_img" src="@/static/but_close.png"
        /></view>
      </view>
      <view v-else class="dish_detail_pop">
        <!-- <scroll-view :enhanced="true" class="dish_items" scroll-y> -->
        <view
          v-for="(item, index) in dishMealData"
          :key="index"
          class="dish_item"
        >
          <image class="div_big_image" :src="item.image" />
          <view class="title"
            >{{ item.name }}<text>{{ "X" + item.copies }}</text></view
          >
          <view class="desc">{{ item.description }}</view>
        </view>
        <!-- </scroll-view> -->
        <view class="but_item">
          <view class="price"
            ><text class="ico">￥</text>{{ dishDetailes.price }}</view
          >
          <view
            v-if="dishDetailes.dishNumber && dishDetailes.dishNumber > 0"
            class="active"
          >
            <image
              class="dish_red"
              src="@/static/btn_red.png"
              @tap.stop="redDishAction('detail', '普通', dishDetailes)"
            />
            <text class="dish_number">{{ dishDetailes.dishNumber }}</text>
            <image
              class="dish_add"
              src="@/static/btn_add.png"
              @tap.stop="addDishAction('detail', '普通', dishDetailes)"
            />
          </view>
          <view v-else
            ><view
              class="dish_card_add"
              @tap.stop="addDishAction('detail', '普通', dishDetailes)"
              >加入购物车</view
            ></view
          >
        </view>
        <view class="close" @tap.stop="openDetailPop = false"
          ><image class="close_img" src="@/static/but_close.png"
        /></view>
      </view>
    </view>
    <!-- Spec Selector Popup -->
    <view v-if="openMoreNormPop" class="pop_mask">
      <view class="more_norm_pop">
        <view class="title">{{ moreNormDishdata.name }}</view>
        <scroll-view :enhanced="true" class="items_cont" scroll-y>
        <view
          v-for="(obj, index) in moreNormdata"
          :key="index"
          class="item_row"
        >
          <view class="flavor_name">{{ obj.name }}</view>
          <view class="flavor_item">
            <view
              v-for="(item, ind) in obj.value"
              :key="ind"
              :class="[
                true ? 'item' : '',
                checkMoreNormPop(obj, item) ? 'act' : '',
              ]"
              @tap.stop="checkMoreNormSelect(index, ind)"
              >{{ item }}</view
            >
          </view>
        </view>
        </scroll-view>
        <view class="but_item">
          <view class="price"
            ><text class="ico">￥</text>{{ moreNormDishdata.price }}</view
          >
          <view class="active"
            ><view
              class="dish_card_add"
              @tap.stop="addShop('detail', '普通', moreNormDishdata)"
              >加入购物车</view
            ></view
          >
        </view>
        <view class="close" @tap.stop="closeMoreNorm(moreNormDishdata)"
          ><image class="close_img" src="@/static/but_close.png"
        /></view>
      </view>
    </view>
    <!-- Cart Popup -->
    <view
      v-show="openOrderCartList"
      class="pop_mask"
      @tap.stop="openOrderCartList = false"
    >
      <view class="cart_pop" @tap.stop>
        <view class="top_title">
          <view class="tit">购物车</view>
          <view class="clear" @tap.stop="clearCardOrder"
            ><image class="clear_icon" src="@/static/clear.png" /><text
              class="clear-des"
              >清空</text
            ></view
          >
        </view>
        <scroll-view :enhanced="true" class="card_order_list" scroll-y>
          <view
            v-for="(obj, index) in orderListData"
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
                ><text class="ico">￥</text>{{ obj.amount }}</view
              >
              <view class="dish_active">
                <image
                  v-if="obj.number > 0"
                  class="dish_red"
                  src="@/static/btn_red.png"
                  @tap.stop="redDishAction('cart', '购物车', obj)"
                />
                <text v-if="obj.number > 0" class="dish_number">{{
                  obj.number
                }}</text>
                <image
                  class="dish_add"
                  src="@/static/btn_add.png"
                  @tap.stop="addDishAction('cart', '购物车', obj)"
                />
              </view>
            </view>
          </view>
          <view class="seize_seat" />
        </scroll-view>
      </view>
    </view>
    <view v-show="loaddingSt" class="pop_mask"
      ><view class="lodding"
        ><image class="lodding_ico" src="@/static/lodding.gif" /></view
    ></view>
  </view>
</template>
<script>
import {
  getCategoryList,
  getDishList,
  getShopStatus,
  addShoppingCart,
  subShoppingCart,
  cleanShoppingCart,
} from "@/utils/api";
import { mapState, mapMutations } from "vuex";

export default {
  components: {
    "nav-bar": () => import("@/components/nav-bar/nav-bar.vue"),
    phone: () => import("@/components/phone/phone.vue"),
  },
  data() {
    return {
      openOrderCartList: false,
      typeListData: [],
      dishListItems: [],
      dishDetailes: {},
      openDetailPop: false,
      openMoreNormPop: false,
      moreNormDishdata: null,
      moreNormdata: null,
      dishMealData: null,
      flavorDataes: [],
      selectedFlavorIndices: [],
      typeIndex: 0,
      orderDishNumber: 0,
      orderDishPrice: 0,
      phoneData: "",
      shopStatus: null,
      scrollTop: 0,
      itemId: "",
    };
  },
  computed: {
    ...mapState(["shopInfo", "shopPhone", "orderListData", "lodding", "token"]),
    loaddingSt() {
      return this.lodding;
    },
    ht() {
      return uni.getSystemInfoSync().statusBarHeight + 44 + "px";
    },
  },
  methods: {
    ...mapMutations([
      "setShopStatus",
      "initdishListMut",
      "setLodding",
      "setOrderData",
    ]),
    async getData() {
      try {
        const catRes = await getCategoryList({ type: 1 });
        if (catRes.code === 1 && catRes.data) {
          this.typeListData = catRes.data;
          if (catRes.data.length > 0) this.swichMenu(0);
        }
        const statusRes = await getShopStatus();
        if (statusRes.data !== undefined) {
          this.shopStatus = statusRes.data;
          this.setShopStatus(this.shopStatus);
        }
      } catch (e) {
        console.error(e);
      }
    },
    async swichMenu(index) {
      this.typeIndex = index;
      const cat = this.typeListData[index];
      if (!cat) return;
      try {
        const res = await getDishList({ categoryId: cat.id });
        if (res.code === 1 && res.data) {
          this.dishListItems = res.data.map((d) => ({ ...d, dishNumber: 0 }));
        }
      } catch (e) {
        console.error(e);
      }
    },
    getNewImage(image) {
      return image || "";
    },
    addDishAction(source, type, item) {
      item.dishNumber = (item.dishNumber || 0) + 1;
      this.computOrderInfo();
      const existing = this.orderListData.find((d) => d.id === item.id);
      if (existing) {
        existing.number = item.dishNumber;
      } else {
        this.initdishListMut([...this.orderListData, { ...item, number: 1 }]);
      }
      addShoppingCart({
        dishId: item.id,
        dishFlavor: item.dishFlavor || "",
      }).catch(() => {});
    },
    redDishAction(source, type, item) {
      if (item.dishNumber > 0) {
        item.dishNumber--;
        this.computOrderInfo();
      }
      subShoppingCart({
        dishId: item.id,
        dishFlavor: item.dishFlavor || "",
      }).catch(() => {});
    },
    addShop(source, type, item) {
      item.dishNumber = (item.dishNumber || 0) + 1;
      this.computOrderInfo();
      this.openMoreNormPop = false;
      const existing = this.orderListData.find((d) => d.id === item.id);
      if (existing) {
        existing.number = item.dishNumber;
      } else {
        this.initdishListMut([...this.orderListData, { ...item, number: 1 }]);
      }
      addShoppingCart({
        dishId: item.id,
        dishFlavor: item.dishFlavor || "",
      }).catch(() => {});
    },
    clearCardOrder() {
      this.initdishListMut([]);
      this.dishListItems.forEach((d) => (d.dishNumber = 0));
      this.computOrderInfo();
      cleanShoppingCart();
    },
    openDetailHandle(item) {
      this.dishDetailes = item;
      this.openDetailPop = true;
    },
    moreNormDataesHandle(item) {
      this.moreNormDishdata = item;
      this.moreNormdata = item.flavors || [];
  
      this.moreNormdata = this.moreNormdata.map((item, index) => {
        return {
          ...item,
          value: JSON.parse(item.value)
        }
      })
      console.log(this.moreNormdata, "this.moreNormdata")
      this.selectedFlavorIndices = (item.flavors || []).map(() => -1);
      this.openMoreNormPop = true;
    },
    checkMoreNormPop(obj, item) {
      const idx = this.moreNormdata.indexOf(obj);
      if (idx < 0) return false;
      const valIdx = obj.value.indexOf(item);
      return this.selectedFlavorIndices[idx] === valIdx;
    },
    checkMoreNormSelect(objIdx, valIdx) {
      console.log("checkMoreNormSelect");
      const selected = [...this.selectedFlavorIndices];
      selected[objIdx] = selected[objIdx] === valIdx ? -1 : valIdx;
      this.selectedFlavorIndices = selected;
    },
    closeMoreNorm(item) {
      this.openMoreNormPop = false;
    },
    computOrderInfo() {
      let num = 0,
        price = 0;
      this.dishListItems.forEach((d) => {
        if (d.dishNumber > 0) {
          num += d.dishNumber;
          price += d.dishNumber * d.price;
        }
      });
      this.orderDishNumber = num;
      this.orderDishPrice = price;
    },
    goOrder() {
      if (this.orderListData.length === 0) {
        uni.showToast({ title: "请先添加商品", icon: "none" });
        return;
      }
      this.setOrderData(this.orderListData);
      uni.navigateTo({ url: "/pages/order/index" });
    },
    handlePhone() {
      this.phoneData = this.shopPhone;
    },
    disabledScroll() {},
  },
  onLoad() {
    this.getData();
  },
  onShow() {
    if (this.token) this.getData();
  },
};
</script>
<style scoped>
.home_content {
  min-height: 100vh;
  background: #f6f6f6;
  padding-bottom: 120rpx;
}
.restaurant_info_box {
  position: relative;
  width: 100%;
  background: linear-gradient(
    184deg,
    rgba(0, 0, 0, 0.35) 25%,
    rgba(51, 51, 51, 0) 96%
  );
  padding-top: 80rpx;
}
.restaurant_info {
  margin: 0 30rpx;
  background: #fff;
  box-shadow: 0 4rpx 10rpx rgba(69, 69, 69, 0.1);
  border-radius: 8rpx;
  padding: 14rpx 18rpx 22rpx 16rpx;
}
.info_top {
  display: flex;
  padding-bottom: 10rpx;
  border-bottom: 1px dashed #ebebeb;
}
.info_top_left {
  margin-right: 20rpx;
  padding-top: 10rpx;
}
.info_top_left image {
  width: 86rpx;
  height: 86rpx;
}
.info_top_right {
  flex: 1;
}
.right_title {
  display: flex;
  align-items: center;
}
.right_title text {
  font-size: 36rpx;
  font-weight: 500;
  color: #20232a;
  line-height: 50rpx;
}
.right_details {
  display: flex;
}
.details_flex {
  white-space: nowrap;
}
.details_flex .top_icon {
  width: 28rpx;
  height: 28rpx;
  display: inline-block;
  vertical-align: middle;
  margin-right: 6rpx;
}
.details_flex .icon_text {
  font-size: 24rpx;
  color: #333;
  line-height: 36rpx;
  padding-right: 20rpx;
}
.test {
  flex: 1;
}
.info_bottom {
  margin-top: 16rpx;
  display: flex;
  font-size: 24rpx;
  color: #9b9b9b;
}
.info_bottom > view:first-child {
  flex: 1;
}
.word {
  display: block;
  padding-bottom: 20rpx;
}
.phone {
  padding: 10rpx 20rpx 10rpx 40rpx;
  border-left: 1px solid rgba(219, 219, 219, 0.45);
}
.businessStatus {
  display: inline-block;
  width: 92rpx;
  height: 36rpx;
  background: #1dc779;
  border-radius: 8rpx;
  color: #fff;
  font-size: 24rpx;
  line-height: 36rpx;
  text-align: center;
  margin-left: 10rpx;
}
.businessStatus.close {
  background: #999;
}
.restaurant_menu_list {
  display: flex;
  width: 100%;
  height: calc(100vh - 200rpx);
}
.type_list {
  background: #f3f4f7;
  width: 168rpx;
}
.type_list .type_item {
  padding: 20rpx 26rpx 20rpx 20rpx;
  text-align: center;
  font-size: 26rpx;
  color: #666;
}
.type_list .active {
  color: #333;
  background: #fff;
  font-weight: 500;
}
.seize_seat {
  height: 160rpx;
}
.vegetable_order_list {
  background: #fff;
  flex: 1;
  padding: 0 20rpx;
}
.vegetable_order_list .type_item {
  display: flex;
  padding: 20rpx 0;
}
.dish_img {
  width: 172rpx;
}
.dish_img_url {
  width: 172rpx;
  height: 172rpx;
  border-radius: 8rpx;
  display: block;
}
.dish_info {
  position: relative;
  flex: 1;
  padding-left: 20rpx;
  min-height: 172rpx;
}
.dish_name {
  font-size: 32rpx;
  font-weight: 600;
  color: #333;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.dish_label {
  font-size: 22rpx;
  line-height: 40rpx;
  color: #666;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.dish_price {
  height: 72rpx;
  line-height: 72rpx;
  font-size: 36rpx;
  color: #e94e3c;
  font-weight: 500;
}
.dish_price .ico {
  font-size: 24rpx;
}
.dish_active {
  position: absolute;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
}
.dish_add,
.dish_red {
  width: 72rpx;
  height: 72rpx;
  display: block;
}
.dish_number {
  padding: 0 10rpx;
  line-height: 72rpx;
  font-size: 30rpx;
  font-weight: 500;
}
.dish_active_btn {
  position: absolute;
  right: 0;
  bottom: 15rpx;
}
.check_but {
  width: 128rpx;
  height: 48rpx;
  line-height: 48rpx;
  text-align: center;
  background: #ffc200;
  border-radius: 24rpx;
  font-size: 24rpx;
  font-weight: 500;
  color: #333;
}
.no_dish {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9b9b9b;
  font-size: 26rpx;
  background: #fff;
}
.footer_order_buttom {
  position: fixed;
  bottom: 48rpx;
  left: 30rpx;
  right: 30rpx;
  display: flex;
  height: 88rpx;
  background: rgba(0, 0, 0, 0.9);
  border-radius: 50rpx;
  box-shadow: 0 6rpx 10rpx rgba(0, 0, 0, 0.25);
  z-index: 99;
  padding: 0 10rpx;
}
.orderCar {
  flex: 1;
  display: flex;
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
  font-weight: 500;
}
.order_price .ico {
  font-size: 24rpx;
}
.order_but {
  width: 204rpx;
  height: 72rpx;
  line-height: 72rpx;
  border-radius: 72rpx;
  color: #333;
  text-align: center;
  margin-top: 8rpx;
  background: #ffc200;
  font-weight: 500;
}
.pop_mask {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  z-index: 99;
  background: rgba(0, 0, 0, 0.4);
}
.cart_pop {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  /* height: 60vh; */
  background: #fff;
  border-radius: 8rpx 8rpx 0 0;
  padding: 20rpx 30rpx 30rpx;
  box-sizing: border-box;
}
.cart_pop .top_title {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #ebeef5;
  padding-bottom: 20rpx;
}
.cart_pop .top_title .tit {
  font-size: 40rpx;
  font-weight: bold;
  color: #20232a;
}
.cart_pop .top_title .clear {
  color: #999;
  font-size: 28rpx;
  display: flex;
  align-items: center;
}
.clear_icon {
  width: 30rpx;
  height: 30rpx;
  margin-right: 8rpx;
}
.card_order_list {
  height: calc(100% - 70rpx);
}
.dish_detail_pop,
.more_norm_pop {
  width: calc(100vw - 160rpx);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 40rpx;
  background: #fff;
  border-radius: 20rpx;
  /* max-height: 80vh; */
  /* overflow-y: auto; */
}
.div_big_image {
  width: 100%;
  height: 320rpx;
  border-radius: 10rpx;
}
.dish_detail_pop .title,
.more_norm_pop .title {
  font-size: 40rpx;
  line-height: 80rpx;
  text-align: center;
  font-weight: bold;
}
.dish_detail_pop .desc {
  font-size: 24rpx;
  color: #666;
  text-align: center;
}
.dish_items {
  height: 60vh;
}
.dish_item {
  margin-bottom: 20rpx;
}
.dish_item image {
  width: 100%;
  height: 200rpx;
  border-radius: 10rpx;
}
.dish_item .title {
  font-size: 28rpx;
}
.but_item {
  display: flex;
  position: relative;
  margin-top: 20rpx;
}
.but_item .price {
  color: #e94e3c;
  line-height: 88rpx;
  font-size: 48rpx;
  font-weight: bold;
}
.but_item .price .ico {
  font-size: 28rpx;
}
.but_item .active {
  position: absolute;
  right: 0;
  bottom: 20rpx;
  display: flex;
  align-items: center;
}
.dish_card_add {
  width: 200rpx;
  height: 60rpx;
  line-height: 60rpx;
  text-align: center;
  font-weight: 500;
  font-size: 28rpx;
  background: #ffc200;
  border-radius: 30rpx;
}
.items_cont {
  max-height: 50vh;
}
.item_row {
  margin-bottom: 20rpx;
}
.flavor_name {
  font-size: 28rpx;
  color: #666;
  padding-left: 10rpx;
  padding-top: 20rpx;
}
.flavor_item {
  display: flex;
  flex-wrap: wrap;
}
.flavor_item .item {
  border: 1px solid #ffb302;
  border-radius: 12rpx;
  margin: 20rpx 10rpx;
  padding: 0 26rpx;
  height: 60rpx;
  line-height: 60rpx;
  color: #333;
  font-size: 26rpx;
}
.flavor_item .act {
  background: #ffc200;
  border-color: #ffc200;
  font-weight: 500;
}
.close {
  position: absolute;
  bottom: -180rpx;
  left: 50%;
  transform: translateX(-50%);
}
.close_img {
  width: 88rpx;
  height: 88rpx;
}
.lodding {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.lodding_ico {
  width: 160rpx;
  height: 160rpx;
  border-radius: 100%;
}
</style>
