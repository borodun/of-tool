<template>
  <div id="all">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark"><h1 id="ecsList">CPU USAGE PER FUNCTION</h1></div>
      </el-col>
    </el-row>
    <el-form ref="form" label-width="200px">
      <el-form-item label="Choose function">
        <el-select v-model="selectedFunction" paceholder="Select" v-on:focus="ShowFunctions">
          <el-option
              v-for="item in functions"
              :key="item"
              :label="item"
              :value="item">
            <span style="float: left">{{ item }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Request body">
        <el-input
            type="textarea"
            :rows="3"
            placeholder="Request body"
            v-model="requestBody"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" v-on:click="MeasureCPU">INVOKE</el-button>
      </el-form-item>
      <el-form-item>
        <span> Cpu time used on function: {{ cpuTimeUsed }}s</span>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from "axios"

const axios_faas = axios.create({
  baseURL: process.env.VUE_APP_FAAS_API,
});

const axios_prometheus = axios.create({
  baseURL: process.env.VUE_APP_CPUMEM_API,
});

export default {
  name: "CPUPerFunction",

  data() {
    return {
      namespace: "openfaas-fn",
      selectedFunction: "",

      cpu: 0,

      functions: [],

      requestBody: "",
      cpuTimeUsed: 0,
    }
  },
  methods: {
    Sleep(ms) {
      return new Promise(resolve => setTimeout(resolve, ms));
    },
    async MeasureCPU() {
      this.GetCPU()
      let before = this.cpu
      console.log("Measured CPU before INVOKE: " + before)
      this.InvokeFunction()
      await this.Sleep(5000)
      this.GetCPU()
      let after = this.cpu
      console.log("Measured CPU after INVOKE: " + after)

      this.cpuTimeUsed = after - before
    },
    ShowFunctions() {
      axios_faas.get("/list-functions").then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          this.functions = response.data.functions
          console.log(this.pods)
        }
      }).catch(function (error) {
        console.log(error);
      })
    },
    GetCPU() {
      axios_prometheus.get("/get-cpu/" + this.selectedFunction).then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          this.cpu = parseFloat(response.data.rs)
        }
      }).catch(function (error) {
        console.log(error);
      })
    },
    InvokeFunction() {
      axios_faas.get("/invoke-function/" + this.selectedFunction + "/" + this.requestBody).then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          console.log(response.data.rs)
        }
      }).catch(function (error) {
        console.log(error);
      })
    }
  }
}
</script>

<style scoped>
#all {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #303133;
}
</style>