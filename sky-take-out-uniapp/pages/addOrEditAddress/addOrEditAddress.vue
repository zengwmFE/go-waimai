<template>
  <view class="customer-box">
    <nav-bar :title="delId ? '编辑收货地址' : '新增收货地址'" leftIcon="arrowleft" backgroundColor="#333333" @clickLeft="goBack" />
    <view class="add_edit" :style="{ height: 'calc(100% - ' + statusBarHeight + ' - 44px)' }">
      <form class="form_address">
        <view class="uni-form-item form_item">
          <view class="title">联系人:</view>
          <uni-easyinput class="uni-input" placeholder="请填写收货人的姓名" v-model="form.name" :minlength="2" :maxlength="12" />
          <view class="radio">
            <view v-for="item in items" :key="item.value" class="radio-item" @tap.stop="sexChangeHandle(item.value)">
              <image class="radio-img" :src="item.value !== form.sex ? '@/static/icon-radio.png' : '@/static/icon-radio-selected.png'" />
              <text class="radio-label">{{ item.name }}</text>
            </view>
          </view>
        </view>
        <view class="uni-form-item form_item">
          <view class="title">手机号:</view>
          <uni-easyinput class="uni-input" type="number" placeholder="请填写收货人手机号码" v-model="form.phone" :maxlength="11" />
        </view>
        <view class="uni-form-item form_item pad">
          <view class="title">收货地址:</view>
          <view class="update-input">
            <view class="update-adress" @tap.stop="openAddres">
              <text :class="['uni-input', address !== '' ? '' : 'uni-place']">{{ address || '省/市/区' }}</text>
              <text class="addressIcon"><text :class="['icon', showClass ? 'iconOn' : '']" /></text>
            </view>
            <textarea class="detail" placeholder="详细地址（精确到门牌号）" v-model="form.detail" />
          </view>
        </view>
        <view class="uni-form-item form_item tag-box">
          <view class="title">标签:</view>
          <text v-for="item in options" :key="item.type" :class="['tag_text', form.type === item.type ? 'active' : '']" @tap.stop="getTextOption(item.type)">{{ item.name }}</text>
        </view>
      </form>
      <view class="add_address">
        <button class="add_btn" type="primary" plain @tap.stop="addAddressFun">保存地址</button>
        <button v-if="showDel" class="del_btn" type="default" plain @tap.stop="deleteAddressFun">删除地址</button>
      </view>
    </view>
  </view>
</template>
<script>
import { addAddressBook, updateAddressBook, deleteAddressBook, getAddressBookById } from '@/utils/api'
export default {
  data() {
    return {
      showDel: false, showInput: true, showClass: false,
      items: [{ value: '0', name: '先生' }, { value: '1', name: '女士' }],
      options: [{ name: '公司', type: 1 }, { name: '家', type: 2 }, { name: '学校', type: 3 }],
      form: { name: '', phone: '', sex: '0', type: 1, provinceCode: '11', provinceName: '', cityCode: '1101', cityName: '', districtCode: '110102', districtName: '', detail: '' },
      cityPickerValueDefault: [0, 0, 1], address: '', delId: ''
    }
  },
  computed: { statusBarHeight() { return uni.getSystemInfoSync().statusBarHeight + 'px' } },
  methods: {
    goBack() { uni.navigateBack() },
    async queryAddressBookById(id) {
      const res = await getAddressBookById(id)
      if (res.code === 1 && res.data) {
        const d = res.data
        this.form = { name: d.consignee || '', phone: d.phone || '', sex: d.sex || '0', type: d.label || 1, provinceCode: d.provinceCode, provinceName: d.provinceName, cityCode: d.cityCode, cityName: d.cityName, districtCode: d.districtCode, districtName: d.districtName, detail: d.detail || '' }
        this.address = `${d.provinceName}${d.cityName}${d.districtName}`
        this.showDel = true
      }
    },
    getTextOption(type) { this.form.type = type },
    sexChangeHandle(val) { this.form.sex = val },
    openAddres() { /* trigger address picker */ },
    onConfirm(e) { /* address picker callback */ },
    async addAddressFun() {
      const data = {
        consignee: this.form.name, phone: this.form.phone, sex: this.form.sex,
        label: String(this.form.type),
        provinceCode: this.form.provinceCode, provinceName: this.form.provinceName,
        cityCode: this.form.cityCode, cityName: this.form.cityName,
        districtCode: this.form.districtCode, districtName: this.form.districtName,
        detail: this.form.detail, isDefault: 0
      }
      if (!data.consignee || !data.phone || !data.detail) { uni.showToast({ title: '请填写完整信息', icon: 'none' }); return }
      if (this.delId) {
        await updateAddressBook({ ...data, id: this.delId })
      } else {
        await addAddressBook(data)
      }
      uni.navigateBack()
    },
    async deleteAddressFun() {
      if (!this.delId) return
      uni.showModal({
        title: '提示', content: '确定删除该地址吗？',
        success: async (r) => { if (r.confirm) { await deleteAddressBook(this.delId); uni.navigateBack() } }
      })
    }
  },
  async onLoad(options) {
    if (options && options.id) { this.delId = options.id; await this.queryAddressBookById(options.id) }
  }
}
</script>
<style scoped>
.customer-box { height: 100vh; background: #f6f6f6; }
.add_edit { padding: 120rpx 20rpx 40rpx; }
.form_item { background: #fff; border-radius: 12rpx; padding: 30rpx; margin-bottom: 20rpx; }
.title { font-size: 28rpx; color: #333; font-weight: 600; margin-bottom: 16rpx; }
.uni-input { padding: 0; }
.radio { display: flex; margin-top: 20rpx; }
.radio-item { display: flex; align-items: center; margin-right: 40rpx; }
.radio-img { width: 32rpx; height: 32rpx; margin-right: 12rpx; }
.radio-label { font-size: 26rpx; color: #333; }
.update-input { margin-top: 10rpx; }
.update-adress { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #eee; padding-bottom: 16rpx; }
.uni-place { color: #bdbdbd; font-size: 28rpx; }
.addressIcon .icon { width: 32rpx; height: 32rpx; display: inline-block; }
.detail { width: 100%; height: 160rpx; margin-top: 20rpx; padding: 10rpx; font-size: 26rpx; border-radius: 8rpx; background: #f8f8f8; box-sizing: border-box; }
.tag-box { display: flex; flex-wrap: wrap; gap: 20rpx; align-items: center; }
.tag_text { padding: 8rpx 24rpx; border-radius: 20rpx; font-size: 24rpx; background: #f5f5f5; color: #666; }
.tag_text.active { background: #ffc200; color: #333; font-weight: 500; }
.add_address { margin-top: 20rpx; }
.add_btn { width: 100%; height: 86rpx; line-height: 86rpx; border-radius: 8rpx; background: #ffc200; border: 1px solid #ffc200; font-size: 30rpx; font-weight: 600; color: #333; display: flex; align-items: center; justify-content: center; margin-bottom: 20rpx; }
.del_btn { width: 100%; height: 86rpx; line-height: 86rpx; border-radius: 8rpx; background: #fff; border: 1px solid #e94e3c; font-size: 30rpx; color: #e94e3c; display: flex; align-items: center; justify-content: center; }
</style>
