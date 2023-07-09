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
  
<style lang="">
  /* 样式内容 */
  </style>
  