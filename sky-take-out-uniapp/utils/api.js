import store from '@/store'

const BASE_URL = 'http://localhost:8080'

let isRelogin = false
let reloginPromise = null

function doRelogin() {
  if (reloginPromise) return reloginPromise
  reloginPromise = store.dispatch('wechatLogin').then(() => {
    isRelogin = false
    reloginPromise = null
  }).catch(err => {
    isRelogin = false
    reloginPromise = null
    throw err
  })
  return reloginPromise
}

function request(url, method = 'GET', data = {}, _retry = true) {
  return new Promise((resolve, reject) => {
    uni.request({
      url: BASE_URL + url,
      method,
      data,
      header: {
        'Content-Type': 'application/json',
        'authentication': uni.getStorageSync('token') || ''
      },
      success(res) {
        if (res.statusCode === 200) {
          const d = res.data
          if (d.code === 0 && (d.msg === 'TOKEN_INVALID' || d.msg === 'NOT_LOGIN') && _retry) {
            doRelogin().then(() => {
              request(url, method, data, false).then(resolve).catch(reject)
            }).catch(() => {
              uni.reLaunch({ url: '/pages/index/index' })
              reject(d)
            })
            return
          }
          resolve(d)
        } else {
          reject(res.data)
        }
      },
      fail(err) { reject(err) }
    })
  })
}

// 供外部手动触发登录
export { doRelogin }

// User
export function wxLogin(data) { return request('/user/user/login', 'POST', data) }

// Category
export function getCategoryList(data) { return request('/user/category/list', 'GET', data) }

// Dish
export function getDishList(data) { return request('/user/dish/list', 'GET', data) }

// Setmeal
export function getSetmealList(data) { return request('/user/setmeal/list', 'GET', data) }
export function getDishBySetmealId(id) { return request(`/user/setmeal/dish/${id}`, 'GET') }

// Shopping Cart
export function getShoppingCartList() { return request('/user/shoppingCart/list', 'GET') }
export function addShoppingCart(data) { return request('/user/shoppingCart/add', 'POST', data) }
export function subShoppingCart(data) { return request('/user/shoppingCart/sub', 'POST', data) }
export function cleanShoppingCart() { return request('/user/shoppingCart/clean', 'DELETE') }

// Order
export function submitOrder(data) { return request('/user/order/submit', 'POST', data) }
export function getOrderPage(data) { return request('/user/order/historyOrders', 'GET', data) }
export function getOrderDetail(id) { return request(`/user/order/orderDetail/${id}`, 'GET') }
export function cancelOrder(id) { return request(`/user/order/cancel/${id}`, 'PUT') }
export function repetitionOrder(id) { return request(`/user/order/repetition/${id}`, 'POST') }
export function reminder(id) { return request(`/user/order/reminder/${id}`, 'GET') }
export function payment(data) { return request('/user/order/payment', 'PUT', data) }

// Address Book
export function getAddressList() { return request('/user/addressBook/list', 'GET') }
export function getAddressBookDefault() { return request('/user/addressBook/default', 'GET') }
export function getAddressBookById(id) { return request(`/user/addressBook/${id}`, 'GET') }
export function addAddressBook(data) { return request('/user/addressBook', 'POST', data) }
export function updateAddressBook(data) { return request('/user/addressBook', 'PUT', data) }
export function deleteAddressBook(id) { return request(`/user/addressBook/${id}`, 'DELETE') }
export function setDefaultAddress(data) { return request('/user/addressBook/default', 'PUT', data) }

// Shop
export function getShopStatus() { return request('/user/shop/status', 'GET') }
