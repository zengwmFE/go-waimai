export function statusWord(status, time) {
  const map = { 1: '待付款', 2: '待接单', 3: '已接单', 4: '派送中', 5: '已完成', 6: '已取消', 7: '退款' }
  if (status === 1 && time <= 0) return '已超时'
  return map[status] || '未知'
}

export function getOvertime(time) {
  const orderTime = new Date(time).getTime()
  const now = Date.now()
  const diff = (orderTime + 15 * 60 * 1000 - now) / 1000
  return Math.max(0, Math.floor(diff))
}
