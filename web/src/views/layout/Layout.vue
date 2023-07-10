<template>
    <div class="common-layout">
        <el-container>
            <el-header>
                <Header></Header>
            </el-header>
            <el-container>
                <el-aside width="200px">
                    <Aside></Aside>
                </el-aside>
                <el-container>
                    <el-main>Main</el-main>
                    <el-footer>Footer</el-footer>
                </el-container>
            </el-container>
        </el-container>
    </div>
</template>
  
<script lang="ts" setup>
import Aside from '@/views/layout/Aside.vue';
import Header from '@/views/layout/Header.vue';


</script>
  
<style lang="css">
  /* 样式内容 */
  .el-aside {
    height: 100vh;
    background-color: #f6f7fa;
  }
  .el-header {
  }
  .el-main {
    background-color: #3754a9;
  }
  </style>
  