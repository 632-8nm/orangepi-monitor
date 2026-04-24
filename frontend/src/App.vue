<template>
  <header>
    <h1>Orange Pi 系统监控</h1>
    <p>{{ statusText }}</p>
  </header>

  <div class="container">
    <div class="card">
      <div class="card-title">CPU 核心状态</div>
      <div class="card-value"><span>{{ stats.cpu_usage.toFixed(1) }}</span><span class="card-unit">%</span></div>
      <p class="card-subtitle">
        频率: <span>{{ Math.round(stats.cpu_freq) }}</span> MHz |
        温度: <span :style="{ color: cpuTempColor }">{{ stats.cpu_temp }}</span>
      </p>
      <div class="progress-container">
        <div class="progress-bar" :style="{ width: stats.cpu_usage + '%' }"></div>
      </div>
    </div>

    <div class="card">
      <div class="card-title">内存资源</div>
      <div class="card-value"><span>{{ stats.mem_usage.toFixed(1) }}</span><span class="card-unit">%</span></div>
      <p class="card-subtitle">{{ stats.mem_summary }}</p>
      <div class="progress-container">
        <div class="progress-bar" :style="{ width: stats.mem_usage + '%' }"></div>
      </div>
    </div>

    <div class="card">
      <div class="card-title">系统运行信息</div>
      <div class="info-row"><span>OS:</span> <span>{{ stats.os_info }}</span></div>
      <div class="info-row"><span>运行时间:</span> <span>{{ formattedUptime }}</span></div>
      <div class="status-badge">系统状态: 正常</div>
    </div>

    <div class="card">
      <div class="card-title">实时网络传输</div>
      <div class="net-item">
        <span class="net-label">下载 (↓)</span>
        <span class="net-speed">{{ stats.net_down.toFixed(1) }}</span> <small>KB/s</small>
      </div>
      <div class="net-item">
        <span class="net-label">上传 (↑)</span>
        <span class="net-speed">{{ stats.net_up.toFixed(1) }}</span> <small>KB/s</small>
      </div>
    </div>
  </div>

  <div class="status-footer">
    <span class="status-dot"></span> 数据每秒实时更新
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, onMounted, reactive, ref } from "vue";

const stats = reactive({
  cpu_usage: 0,
  cpu_freq: 0,
  cpu_temp: "--",
  mem_usage: 0,
  mem_summary: "0 / 0 GB",
  os_info: "--",
  uptime: 0,
  net_down: 0,
  net_up: 0,
});

const lastUpdate = ref("");
const failCount = ref(0);
let timer = null;

const formattedUptime = computed(() => {
  const seconds = Number(stats.uptime) || 0;
  const d = Math.floor(seconds / (3600 * 24));
  const h = Math.floor((seconds % (3600 * 24)) / 3600);
  const m = Math.floor((seconds % 3600) / 60);
  return `${d}天 ${h}时 ${m}分`;
});

const statusText = computed(() => {
  if (failCount.value > 0) return `连接后端失败 (重试次数: ${failCount.value})...`;
  return lastUpdate.value ? `系统状态正常 | 最后更新: ${lastUpdate.value}` : "正在连接服务器...";
});

const cpuTempColor = computed(() => (parseFloat(stats.cpu_temp) > 60 ? "#ef4444" : "#f8fafc"));

const fetchStats = async () => {
  try {
    const response = await fetch("/api/stats");
    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
    const data = await response.json();
    Object.assign(stats, data);
    lastUpdate.value = new Date().toLocaleTimeString();
    failCount.value = 0;
  } catch (error) {
    failCount.value += 1;
    console.error("Failed to fetch stats:", error);
  }
};

onMounted(() => {
  fetchStats();
  timer = setInterval(fetchStats, 1000);
});

onBeforeUnmount(() => {
  if (timer) clearInterval(timer);
});
</script>
