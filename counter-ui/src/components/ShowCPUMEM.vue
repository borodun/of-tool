<template>
  <div id="all">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark"><h1 id="ecsList">CPU MEM INFO</h1></div>
      </el-col>
    </el-row>
    <el-form ref="form" label-width="200px">
      <el-form-item label="Choose namespace">
        <el-select v-model="namespace" placeholder="Select" v-on:focus="ShowNamespaces">
          <el-option
              v-for="item in namespaces"
              :key="item.namespace"
              :label="item.namespace"
              :value="item.namespace">
            <span style="float: left">{{ item.namespace }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Choose pod">
        <el-select v-model="pod" placeholder="Select" v-on:focus="ShowPods" v-on:change="GetCPUMEM">
          <el-option
              v-for="item in pods"
              :key="item.pod"
              :label="item.pod"
              :value="item.pod">
            <span style="float: left">{{ item.pod }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-table
            :data="info"
            border
            style="width: 100%">
          <el-table-column
              prop="cpu"
              label="CPU"
              width="180">
          </el-table-column>
          <el-table-column
              prop="mem"
              label="MEM"
              width="180">
          </el-table-column>
        </el-table>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from "axios"

const axios_instance = axios.create({
  baseURL: process.env.VUE_APP_CPUMEM_API,
});

export default {
  name: "Info",

  data() {
    return {
      namespace: "",
      pod: "",

      info: [{
        cpu: "1",
        mem: "1",
      }],

      namespaces: [{
        namespace: "",
      }],
      pods: [{
        pod: "",
      }]

    }
  },
  methods: {
    GetCPUMEM(){
      this.GetCPU()
      this.GetMEM()
    },
    ShowNamespaces() {
      axios_instance.get("/get-namespaces").then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          this.namespaces = response.data.namespaces
          console.log(this.namespaces)
        }
      }).catch(function (error) {
        console.log(error);
      })
    },
    ShowPods() {
      axios_instance.get("/get-pods/" + this.namespace).then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          this.pods = response.data.pods
          console.log(this.pods)
        }
      }).catch(function (error) {
        console.log(error);
      })
    },
    GetCPU() {
      axios_instance.get("/get-cpu/" + this.pod).then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          this.info[0].cpu = response.data.rs
          console.log(this.pods)
        }
      }).catch(function (error) {
        console.log(error);
      })
    },
    GetMEM() {
      axios_instance.get("/get-mem/" + this.pod).then((response) => {
        console.log(response.data)
        if (response.data.err !== null) {
          console.log("error")
        } else {
          this.info[0].mem = response.data.rs
          console.log(this.pods)
        }
      }).catch(function (error) {
        console.log(error);
      })
    }
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
}

#all {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #303133;
}

#ecsList {
  font-weight: 30;
  font-size: 25pt;
  text-align: center;
}

h2 {
  font-weight: 15;
  font-size: 15pt;
  margin: 30px 30px;
}
</style>