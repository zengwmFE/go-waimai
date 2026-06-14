# 苍穹外卖 API 接口文档

> 基于小程序前端 + 管理后台前端源码提取，共 **76 个接口**

## 通用说明

- **Base URL**: `http://localhost:8080`
- **管理后台代理**: 开发环境 `/api` → `http://localhost:8080/admin`
- **统一响应格式**:
  ```json
  { "code": 1, "msg": "success", "data": {} }
  ```
  - `code=1` 成功，`code=0` 失败
- **认证方式**:
  - 管理端: Header `token` 携带 JWT
  - 用户端: Header `authentication` 携带 JWT

---

## 一、管理端接口 /admin

### 1.1 员工管理 /admin/employee

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/admin/employee/login` | 员工登录 | `{ username, password }` |
| POST | `/admin/employee/logout` | 员工退出 | - |
| GET | `/admin/employee/page` | 员工分页查询 | `?page=&pageSize=&name=` |
| POST | `/admin/employee` | 新增员工 | `{ name, username, phone, sex, idNumber }` |
| PUT | `/admin/employee` | 编辑员工 | `{ id, name, username, phone, sex, idNumber }` |
| GET | `/admin/employee/:id` | 查询员工详情 | 路径参数 id |
| POST | `/admin/employee/status/:status` | 启用/禁用 | `?id=` (status: 1启用 0禁用) |
| PUT | `/admin/employee/editPassword` | 修改密码 | `{ empId, oldPassword, newPassword }` |

### 1.2 分类管理 /admin/category

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/admin/category` | 新增分类 | `{ name, type, sort }` |
| GET | `/admin/category/page` | 分类分页 | `?page=&pageSize=&name=&type=` |
| DELETE | `/admin/category/:id` | 删除分类 | 路径参数 id |
| PUT | `/admin/category` | 修改分类 | `{ id, name, type, sort }` |
| POST | `/admin/category/status/:status` | 启用/禁用 | `?id=` |
| GET | `/admin/category/list` | 分类列表 | `?type=` (1菜品分类 2套餐分类) |

### 1.3 菜品管理 /admin/dish

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/admin/dish` | 新增菜品 | `{ name, categoryId, price, image, description, flavors: [{name, value}] }` |
| GET | `/admin/dish/page` | 菜品分页 | `?page=&pageSize=&name=&categoryId=&status=` |
| DELETE | `/admin/dish` | 批量删除 | `?ids=1,2,3` |
| GET | `/admin/dish/:id` | 菜品详情 | 路径参数 id |
| PUT | `/admin/dish` | 修改菜品 | `{ id, name, categoryId, price, image, description, flavors }` |
| POST | `/admin/dish/status/:status` | 起售/停售 | `?id=` (status: 1起售 0停售) |
| GET | `/admin/dish/list` | 菜品列表 | `?categoryId=` |

### 1.4 套餐管理 /admin/setmeal

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/admin/setmeal` | 新增套餐 | `{ name, categoryId, price, image, description, setmealDishes: [{dishId, copies}] }` |
| GET | `/admin/setmeal/page` | 套餐分页 | `?page=&pageSize=&name=&categoryId=&status=` |
| DELETE | `/admin/setmeal` | 批量删除 | `?ids=1,2,3` |
| GET | `/admin/setmeal/:id` | 套餐详情 | 路径参数 id |
| PUT | `/admin/setmeal` | 修改套餐 | `{ id, name, categoryId, price, image, description, setmealDishes }` |
| POST | `/admin/setmeal/status/:status` | 起售/停售 | `?id=` |
| GET | `/admin/setmeal/list` | 套餐列表 | `?categoryId=` |
| GET | `/admin/setmeal/dish/:id` | 套餐内菜品 | 路径参数 id (套餐ID) |

### 1.5 订单管理 /admin/order

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/admin/order/conditionSearch` | 订单搜索 | `?page=&pageSize=&number=&phone=&status=&beginTime=&endTime=` |
| GET | `/admin/order/details/:orderId` | 订单详情 | 路径参数 orderId |
| GET | `/admin/order/statistics` | 订单统计 | - (返回各状态数量) |
| PUT | `/admin/order/confirm` | 接单 | `{ id }` |
| PUT | `/admin/order/rejection` | 拒单 | `{ id, rejectionReason }` |
| PUT | `/admin/order/cancel` | 取消订单 | `{ id, cancelReason }` |
| PUT | `/admin/order/delivery/:id` | 派送订单 | 路径参数 id |
| PUT | `/admin/order/complete/:id` | 完成订单 | 路径参数 id |

### 1.6 店铺管理 /admin/shop

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/admin/shop/status` | 获取营业状态 | - |
| PUT | `/admin/shop/:status` | 设置营业状态 | 路径参数 status (1营业 0打烊) |

### 1.7 公共接口 /admin/common

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/admin/common/upload` | 文件上传 | FormData: `file` |
| GET | `/admin/common/download` | 文件下载 | `?name=` |

### 1.8 工作台 /admin/workspace

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/admin/workspace/businessData` | 今日营业数据 | - |
| GET | `/admin/workspace/overviewOrders` | 订单概览 | - |
| GET | `/admin/workspace/overviewDishes` | 菜品概览 | - |
| GET | `/admin/workspace/overviewSetmeals` | 套餐概览 | - |

### 1.9 数据统计 /admin/report

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/admin/report/turnoverStatistics` | 营业额统计 | `?begin=&end=` |
| GET | `/admin/report/userStatistics` | 用户统计 | `?begin=&end=` |
| GET | `/admin/report/ordersStatistics` | 订单统计 | `?begin=&end=` |
| GET | `/admin/report/top10` | 销量排名Top10 | `?begin=&end=` |
| GET | `/admin/report/export` | 导出报表 | - (返回 Excel Blob) |

### 1.10 消息通知 /admin/messages

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/admin/messages/page` | 消息列表 | `?page=&pageSize=` |
| GET | `/admin/messages/countUnread` | 未读消息数 | - |
| PUT | `/admin/messages/batch` | 全部标记已读 | `{ ids }` |
| PUT | `/admin/messages/:id` | 标记单条已读 | 路径参数 id |

---

## 二、用户端接口 /user

### 2.1 用户认证 /user/user

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/user/user/login` | 微信登录 | `{ code }` (wx.login的code) |

- **响应**: `{ token, id, name }`
- **说明**: 首次登录自动注册

### 2.2 分类 /user/category

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/category/list` | 获取分类列表 | `?type=` (1菜品 2套餐) |

### 2.3 菜品 /user/dish

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/dish/list` | 按分类查菜品 | `?categoryId=` |
| GET | `/user/dish/flavor/:dishId` | 菜品口味 | 路径参数 dishId |

### 2.4 套餐 /user/setmeal

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/setmeal/list` | 按分类查套餐 | `?categoryId=` |
| GET | `/user/setmeal/dish/:id` | 套餐内菜品 | 路径参数 id (套餐ID) |

### 2.5 店铺 /user/shop

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/shop/status` | 获取营业状态 | - |

### 2.6 购物车 /user/shoppingCart

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/shoppingCart/list` | 查看购物车 | - |
| POST | `/user/shoppingCart/add` | 添加商品 | `{ dishId?, setmealId?, dishFlavor? }` |
| POST | `/user/shoppingCart/sub` | 减少商品 | `{ dishId?, setmealId?, dishFlavor? }` |
| DELETE | `/user/shoppingCart/clean` | 清空购物车 | - |

### 2.7 订单 /user/order

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| POST | `/user/order/submit` | 提交订单 | `{ addressBookId, payMethod, remark, amount, estimatedDeliveryTime, deliveryStatus, tablewareNumber, tablewareStatus, packAmount }` |
| PUT | `/user/order/payment` | 支付订单 | `{ orderNumber, payMethod }` |
| GET | `/user/order/historyOrders` | 历史订单 | `?page=&pageSize=&status=` (status: ''全部 '1'待付款 '6'已取消) |
| GET | `/user/order/orderDetail/:id` | 订单详情 | 路径参数 id |
| PUT | `/user/order/cancel/:id` | 取消订单 | 路径参数 id |
| POST | `/user/order/repetition/:id` | 再来一单 | 路径参数 id |
| GET | `/user/order/reminder/:id` | 催单 | 路径参数 id |

- **支付响应**: `{ nonceStr, packageStr, paySign, timeStamp, signType }` → 调起微信支付

### 2.8 地址簿 /user/addressBook

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/addressBook/list` | 地址列表 | - |
| GET | `/user/addressBook/default` | 默认地址 | - |
| PUT | `/user/addressBook/default` | 设为默认 | `{ id }` |
| POST | `/user/addressBook` | 新增地址 | `{ consignee, phone, sex, label, provinceCode, cityCode, districtCode, provinceName, cityName, districtName, detail }` |
| PUT | `/user/addressBook` | 编辑地址 | `{ id, consignee, phone, sex, label, ... }` |
| DELETE | `/user/addressBook` | 删除地址 | `?id=` |
| GET | `/user/addressBook/:id` | 地址详情 | 路径参数 id |

### 2.9 文件下载 /user/common

| 方法 | 路径 | 说明 | 请求参数 |
|------|------|------|----------|
| GET | `/user/common/download` | 图片下载 | `?name=` |

---

## 三、WebSocket

| 路径 | 说明 |
|------|------|
| `ws://localhost:8080/ws/:sid` | 管理端订单提醒 (sid为客户端唯一标识) |

**消息格式**:
```json
{ "type": 1, "orderId": 12345, "content": "订单号：xxx" }
```
- `type=1` 新订单提醒
- `type=2` 催单提醒

---

## 四、订单状态码

| 状态码 | 含义 |
|--------|------|
| 1 | 待付款 |
| 2 | 待接单 |
| 3 | 已接单 |
| 4 | 派送中 |
| 5 | 已完成 |
| 6 | 已取消 |

## 五、支付状态码

| 状态码 | 含义 |
|--------|------|
| 0 | 未支付 |
| 1 | 已支付 |
| 2 | 退款 |

---

## 六、关键响应示例

### 管理端登录成功
```json
{
  "code": 1,
  "msg": "success",
  "data": {
    "id": 1,
    "name": "管理员",
    "userName": "admin",
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### 分页查询
```json
{
  "code": 1,
  "msg": "success",
  "data": {
    "total": 100,
    "records": [...]
  }
}
```

### 微信登录成功
```json
{
  "code": 1,
  "msg": "success",
  "data": {
    "id": 1,
    "name": "微信用户",
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```
