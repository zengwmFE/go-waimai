<script>
export default {
  onLaunch() {
    console.log('App Launch')
    this.doWechatLogin()
  },
  onShow() {
    console.log('App Show')
    if (!this.$store.state.token) {
      this.doWechatLogin()
    }
  },
  onHide() { console.log('App Hide') },
  methods: {
    doWechatLogin() {
      this.$store.dispatch('wechatLogin').then(user => {
        console.log('登录成功', user)
        // 新用户没有设置昵称，提示去设置
        if (!user.name) {
          uni.showToast({ title: '请前往「我的」设置昵称和头像', icon: 'none', duration: 2000 })
        }
      }).catch(err => {
        console.error('登录失败', err)
        uni.showToast({ title: '登录失败，请重试', icon: 'none' })
      })
    }
  }
}
</script>
<style>
page { background-color: #f6f6f6; }
</style>
