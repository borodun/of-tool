<template>
  <el-tabs id="h" type="border-card">
    <el-tab-pane label="Prometheus"><Info></Info></el-tab-pane>
    <el-tab-pane label="CPU time per function"><CPUPerFunction></CPUPerFunction></el-tab-pane>
  </el-tabs>
</template>

<script>
import Info from './ShowCPUMEM'
import CPUPerFunction from "./CPUPerFunction";
export default {
  name: "List",
  components: {
    Info,
    CPUPerFunction
  }
}
</script>

<style>
#h {
  width: 60%;
  margin: 100px auto 0;
}
</style>