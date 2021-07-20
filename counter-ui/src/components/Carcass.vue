<template>
  <el-tabs id="h" type="border-card">
    <el-tab-pane label="Project info"><Info></Info></el-tab-pane>
    <el-tab-pane label="Other">Nothing there</el-tab-pane>
  </el-tabs>
</template>

<script>
import Info from './ShowCPUMEM'
export default {
  name: "List",
  components: {
    Info,
  }
}
</script>

<style>
#h {
  width: 60%;
  margin: 100px auto 0;
}
</style>