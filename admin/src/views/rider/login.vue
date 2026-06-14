<template>
  <div class="rider-login">
    <div class="login-card">
      <h2>骑手登录</h2>
      <el-form :model="form" :rules="rules" ref="loginForm">
        <el-form-item prop="phone">
          <el-input v-model="form.phone" placeholder="请输入手机号" size="large" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="large" style="width:100%" :loading="loading" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
      <p class="tip">提示：手机号 13800000001-13800000010 为测试骑手</p>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios'

export default {
  data() {
    return {
      loading: false,
      form: { phone: '' },
      rules: { phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }] }
    }
  },
  methods: {
    async handleLogin() {
      const valid = await (this.$refs.loginForm as any).validate().catch(() => false)
      if (!valid) return
      this.loading = true
      try {
        const res = await axios.post('/api/rider/login', { phone: this.form.phone })
        if (res.data.code === 1) {
          localStorage.setItem('rider_token', res.data.data.token)
          localStorage.setItem('rider_info', JSON.stringify(res.data.data))
          this.$router.push('/rider/orders')
        } else {
          this.$message.error(res.data.msg)
        }
      } catch (e) {
        this.$message.error('登录失败')
      } finally { this.loading = false }
    }
  }
}
</script>

<style scoped>
.rider-login { display: flex; justify-content: center; align-items: center; height: 100vh; background: #f0f2f5; }
.login-card { width: 400px; padding: 40px; background: #fff; border-radius: 8px; box-shadow: 0 2px 12px rgba(0,0,0,.1); }
.login-card h2 { text-align: center; margin-bottom: 30px; color: #303133; }
.tip { text-align: center; color: #999; font-size: 13px; margin-top: 16px; }
</style>
