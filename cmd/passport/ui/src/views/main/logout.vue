<template>
  <div>

  </div>
</template>
<script setup lang="ts">
import { ref,reactive,onMounted} from 'vue'
import { useRoute } from 'vue-router';
import {Local} from "/@/utils/storage";
import router from "/@/router";
const route = useRoute();
const state = reactive({
  originUrl:""
})
onMounted(() => {
  logoutFunc()
})

function logoutFunc(){
  Local.remove('token')
  if(route.query.url===""||route.query.url===undefined){
    state.originUrl="http://user.thingple.io/"//默认返回地址
    // Local.remove('token')
    // ElMessage.error("路径错误")
  }else{
    state.originUrl=encodeURIComponent(route.query.url)
    console.log('originUrl',state.originUrl)
  }
  //登录
  router.push({
    path:'/login',
    query:{url:state.originUrl}
  })
}
</script>
<style scoped lang="scss">

</style>
