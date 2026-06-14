import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: uni.getStorageSync('token') || '',
    sessionId: '',
    baseUserInfo: uni.getStorageSync('baseUserInfo') || null,
    shopInfo: {
      name: '苍穹餐厅',
      description: '苍穹餐厅为顾客打造专业的大众化美食外送餐饮',
      address: '北京市朝阳区新街大道一号楼8层',
      distance: '1.5km',
      deliveryFee: 6,
      deliveryTime: '12min'
    },
    shopPhone: '18500557668',
    shopStatus: 1,
    orderListData: [],
    orderData: null,
    arrivals: [{ time: '立即送出' }],
    remarkData: '',
    addressData: null,
    addressBackUrl: '/pages/my/my',
    lodding: false,
    isLogin: false
  },
  mutations: {
    setToken(state, token) {
      state.token = token
      uni.setStorageSync('token', token)
    },
    setSessionId(state, id) { state.sessionId = id },
    setBaseUserInfo(state, info) {
      state.baseUserInfo = info
      uni.setStorageSync('baseUserInfo', info)
    },
    setShopInfo(state, info) { state.shopInfo = { ...state.shopInfo, ...info } },
    setShopPhone(state, phone) { state.shopPhone = phone },
    setShopStatus(state, status) { state.shopStatus = status },
    setOrderData(state, data) { state.orderData = data },
    setArrivalTime(state, time) { const a = [...state.arrivals]; a[0] = time; state.arrivals = a },
    setRemark(state, val) { state.remarkData = val },
    setAddress(state, addr) { state.addressData = addr },
    setAddressBackUrl(state, url) { state.addressBackUrl = url },
    setLodding(state, val) { state.lodding = val },
    initdishListMut(state, list) { state.orderListData = list },
    setIsLogin(state, val) { state.isLogin = val }
  },
  actions: {
    wechatLogin({ commit }) {
      return new Promise((resolve, reject) => {
        uni.login({
          provider: 'weixin',
          success: (loginRes) => {
            uni.request({
              url: 'http://localhost:8080/user/user/login',
              method: 'POST',
              data: { code: loginRes.code },
              header: { 'Content-Type': 'application/json' },
              success: (res) => {
                if (res.statusCode === 200 && res.data && res.data.data) {
                  const user = res.data.data
                  commit('setToken', user.token)
                  commit('setBaseUserInfo', {
                    id: user.id,
                    nickName: user.name || '',
                    avatarUrl: user.avatar || ''
                  })
                  commit('setIsLogin', true)
                  resolve(user)
                } else {
                  reject(res.data)
                }
              },
              fail: (err) => reject(err)
            })
          },
          fail: (err) => reject(err)
        })
      })
    }
  }
})
export default store
