<template>
    <el-menu
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
    >
      <el-menu-item index="0">Groot</el-menu-item>
      <div class="flex-grow" ></div>
      <el-menu-item @click="scan">Scan</el-menu-item>
      <el-sub-menu index="2">
        <template #title>Workspace</template>
        <el-menu-item index="2-1">item one</el-menu-item>
        <el-menu-item index="2-2">item two</el-menu-item>
        <el-menu-item index="2-3">item three</el-menu-item>
        <el-sub-menu index="2-4">
          <template #title>item four</template>
          <el-menu-item index="2-4-1">item one</el-menu-item>
          <el-menu-item index="2-4-2">item two</el-menu-item>
          <el-menu-item index="2-4-3">item three</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
    </el-menu>

    <el-dialog v-model="centerDialogVisible" title="Create Scan" width="50%" center>
      <el-form :model="form" label-width="120px">
        <el-form-item label="Project name">
      <el-input v-model="form.name" />
    </el-form-item>
    <el-form-item label="Domains">
      <el-input v-model="form.domains" type="textarea" />
    </el-form-item>
      </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="centerDialogVisible = false">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>
  </template>
  <script lang="ts" setup>
  import { ref } from 'vue'
  import { reactive } from 'vue'
  let centerDialogVisible = ref(false);
  function scan() {
    centerDialogVisible.value = true
  }
  const form = reactive({
  name: '',
  domains: ''
})
  
  </script>
  
  <style>
  .flex-grow {
    flex-grow: 1;
  }
  </style>