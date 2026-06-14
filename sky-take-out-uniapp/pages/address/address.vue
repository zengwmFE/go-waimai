<template>
  <view class="customer-box">
    <nav-bar :title="addressList && addressList.length > 0 ? '地址管理' : ''" leftIcon="arrowleft" backgroundColor="#333333" @clickLeft="goBack" />
    <view class="address" :style="{ height: 'calc(100% - 136rpx - ' + statusBarHeight + ' - 44px - 20rpx)' }">
      <view v-if="addressList && addressList.length > 0" class="address_content">
        <view v-for="(item, index) in addressList" :key="index" class="address_liests">
          <view class="list_item_top" @tap.stop="choseAddress(index, item)">
            <view class="item_left">
              <view class="details">
                <text :class="['tag', 'tag' + item.label]">{{ getLableVal(item) }}</text>
                <text class="address_word">{{ item.provinceName + item.cityName + item.districtName + item.detail }}</text>
              </view>
              <view class="sale">
                <text class="name">{{ item.sex === '0' ? item.consignee + ' 男士' : item.consignee + ' 女士' }}</text>
                <text class="num">{{ item.phone }}</text>
              </view>
            </view>
            <view class="item_right">
              <image class="edit" src="@/static/edit.png" @tap.stop="addOrEdit('编辑', item)" />
            </view>
          </view>
          <view class="list_item_bottom">
            <label class="radio" @tap.stop="getRadio(index, item)">
              <radio v-if="testValue" class="item_radio" color="#FFC200" :value="item.id" :checked="item.isDefault === 1 && isActive === index" />
              设为默认地址
            </label>
          </view>
        </view>
      </view>
      <view v-if="isEmpty" class="no_address">
        <text class="no_word">一个地址都没有哦</text>
      </view>
      <view class="add_address">
        <button class="add_btn" type="primary" plain @tap.stop="addOrEdit('新增')">
          <text class="add-icon">+</text>新增收货地址
        </button>
      </view>
    </view>
  </view>
</template>
<script>
import { getAddressList as getAddrList, setDefaultAddress } from '@/utils/api'
import { mapState, mapMutations } from 'vuex'
export default {
  components: { NavBar: () => import('@/components/nav-bar/nav-bar.vue') },
  data() { return { testValue: true, addressList: [], formRouter: '', isActive: null, isEmpty: false } },
  computed: {
    ...mapState(['addressBackUrl']),
    statusBarHeight() { return uni.getSystemInfoSync().statusBarHeight + 'px' }
  },
  methods: {
    ...mapMutations(['setAddress']),
    goBack() { uni.navigateBack() },
    getLableVal(item) { const map = { '0': '', '1': '公司', '2': '家', '3': '学校' }; return map[item.label] || '' },
    async getAddressList() {
      const res = await getAddrList()
      if (res.code === 1) {
        this.addressList = res.data || []
        this.isEmpty = this.addressList.length === 0
      }
    },
    addOrEdit(type, item) {
      const url = type === '新增' ? '/pages/addOrEditAddress/addOrEditAddress' : `/pages/addOrEditAddress/addOrEditAddress?id=${item.id}`
      uni.navigateTo({ url })
    },
    choseAddress(index, item) {
      this.setAddress(item)
      uni.navigateBack()
    },
    async getRadio(index, item) {
      this.isActive = index
      if (item.isDefault === 1) return
      await setDefaultAddress({ id: item.id })
      this.getAddressList()
    }
  },
  onShow() { this.getAddressList() }
}
</script>
<style scoped>
.customer-box { height: 100vh; background: #f6f6f6; }
.address { width: 750rpx; overflow-y: auto; }
.address_content { margin: 0 20rpx; padding-top: 100rpx; padding-bottom: 20rpx; }
.address_liests { background: #fff; border-radius: 12rpx; display: flex; flex-direction: column; margin-top: 20rpx; padding: 0 28rpx 0 12rpx; box-sizing: border-box; }
.list_item_top { flex: 1; display: flex; padding-top: 20rpx; }
.item_left { flex: 1; overflow: hidden; margin-left: 12rpx; }
.details { margin-top: 10rpx; display: flex; height: 40rpx; line-height: 40rpx; }
.address_word { flex: 1; font-size: 28rpx; color: #333; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
.tag { padding: 0 8rpx; border-radius: 4rpx; font-size: 22rpx; color: #fff; margin-right: 10rpx; line-height: 32rpx; height: 32rpx; }
.tag1 { background: #f58c21; }
.tag2 { background: #1dc779; }
.tag3 { background: #4a90d9; }
.sale { margin-top: 20rpx; padding-bottom: 20rpx; }
.sale .name, .sale .num { font-size: 28rpx; color: #999; line-height: 40rpx; }
.sale .num { margin-left: 20rpx; }
.item_right { width: 100rpx; display: flex; align-items: center; justify-content: center; }
.item_right .edit { width: 32rpx; height: 32rpx; }
.list_item_bottom { height: 80rpx; line-height: 80rpx; border-top: 1px solid #efefef; }
.list_item_bottom .radio { font-size: 26rpx; color: #333; }
.item_radio { transform: scale(0.7); }
.no_address { text-align: center; font-size: 28rpx; color: #999; padding-top: 40rpx; }
.add_address { position: fixed; bottom: 40rpx; left: 20rpx; right: 20rpx; }
.add_btn { width: 100%; height: 86rpx; line-height: 86rpx; border-radius: 8rpx; background: #ffc200; border: 1px solid #ffc200; font-size: 30rpx; font-weight: 600; color: #333; display: flex; align-items: center; justify-content: center; }
.add-icon { font-size: 32rpx; margin-right: 8rpx; }
</style>
