<template>
  <el-row :gutter="20">
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">上行</el-col>
          <el-col>{{ bytesToSize(isRunning ? current.up : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">下行</el-col>
          <el-col>{{ bytesToSize(isRunning ? current.down : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">总下载</el-col>
          <el-col>{{ bytesToSize(isRunning ? conn.downloadTotal : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">总上传</el-col>
          <el-col>{{ bytesToSize(isRunning ? conn.uploadTotal : 0) }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="4">
      <el-card shadow="never">
        <el-row :gutter="20">
          <el-col class="label">连接数</el-col>
          <el-col>{{ isRunning ? (conn.connNum || 0) : 0 }}</el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>

  <el-divider />
  <el-button
    :icon="Promotion"
    :type="isSystemProxy ? 'primary' : 'info'"
    @click="setOrUnSetSystemProxy"
  >{{ isSystemProxy ? '取消' : '设置为' }}系统代理</el-button>

  <el-button-group class="ml-4">
    <el-button
      v-if="!isLastConfVer"
      type="warning"
      @click="reloadClashConf"
    >重载配置</el-button>
  </el-button-group>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue'
import { Promotion } from '@element-plus/icons-vue'
import app from '@/util/app'
import { bytesToSize, wsMessage, get } from '@/util'
import { ElMessage } from 'element-plus'
import { useAppStore } from '@/store/app'

const appStore = useAppStore()

const current = ref<Record<string, number>>({})
wsMessage({
  ws: `ws://${window.BaseApi}/traffic`,
  onmessage: (msg) => current.value = msg
})

const conn = computed(() => {
  return appStore.getClashConnInfo()
})

const isRunning = computed(() => {
  return appStore.clashRunning
})

const stopOrStartClash = async() => {
  if (isRunning.value) {
    await app.StopClash()
    appStore.resetClashConnInfo()
    isSystemProxy.value = false
  } else {
    await app.StartClash()
    isSystemProxy.value = true
    const base = await app.GetBaseConf()
    if (window.BaseApi !== base.externalController) {
      window.BaseApi = base.externalController
      location.reload()
    }
  }
}

const isSystemProxy = ref(true)
const setOrUnSetSystemProxy = () => {
  if (isSystemProxy.value) {
    app.UnSetSystemProxy().then(() => {
      isSystemProxy.value = false
    })
  } else {
    app.SetSystemProxy().then(() => {
      isSystemProxy.value = true
    })
  }
}

const isLastConfVer = computed(() => {
  return appStore.getLastConfVer()
})

const reloadClashConf = () => {
  app.StartClash().then(res => {
    ElMessage.success('clash配置应用成功, 但不会更改端口等基础配置')
    appStore.setLastConfVer()
  }).catch(err => {
    ElMessage.success('clash配置应用失败' + err)
    location.reload()
  })
}
</script>

<style lang="scss">
.label {
  font-size: 10px;
}
</style>
