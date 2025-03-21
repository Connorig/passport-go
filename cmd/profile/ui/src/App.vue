<template>
  <el-config-provider :size="getGlobalComponentSize" :locale="getGlobalI18n">
    <router-view v-show="setLockScreen"/>
    <Setings ref="setingsRef" v-show="setLockScreen"/>
  </el-config-provider>
</template>

<script setup lang="ts" name="app">
import {computed, defineAsyncComponent, nextTick, onBeforeMount, onMounted, onUnmounted, ref, watch} from 'vue';
import {useRoute} from 'vue-router';
import {useI18n} from 'vue-i18n';
import {storeToRefs} from 'pinia';
import {useTagsViewRoutes} from '/@/stores/tagsViewRoutes';
import {useThemeConfig} from '/@/stores/themeConfig';
import other from '/@/utils/other';
import {Local, Session} from '/@/utils/storage';
import mittBus from '/@/utils/mitt';
import setIntroduction from '/@/utils/setIconfont';
import {getThemeConfigApi} from "/@/api/theme";
// import {getThemeConfigApi} from "/@/api/themeconfig";

// 引入组件
const Setings = defineAsyncComponent(() => import('/@/layout/navBars/topBar/setings.vue'));

// 定义变量内容
const {messages, locale} = useI18n();
const setingsRef = ref();
const route = useRoute();
const stores = useTagsViewRoutes();
const storesThemeConfig = useThemeConfig();
const {themeConfig} = storeToRefs(storesThemeConfig);
// 设置锁屏时组件显示隐藏
const setLockScreen = computed(() => {
  // 防止锁屏后，刷新出现不相关界面
  // https://gitee.com/lyt-top/vue-next-admin/issues/I6AF8P
  return themeConfig.value.isLockScreen ? themeConfig.value.lockScreenTime > 1 : themeConfig.value.lockScreenTime >= 0;
});
// 获取全局组件大小
const getGlobalComponentSize = computed(() => {
  return other.globalComponentSize();
});
// 获取全局 i18n
const getGlobalI18n = computed(() => {
  return messages.value[locale.value];
});
// 设置初始化，防止刷新时恢复默认
onBeforeMount(() => {
  // 设置批量第三方 icon 图标
  setIntroduction.cssCdn();
  // 设置批量第三方 js
  setIntroduction.jsCdn();
});
// 页面加载时
onMounted(() => {
  nextTick(() => {
    getThemeConfig()
    // 监听布局配'置弹窗点击打开
    mittBus.on('openSetingsDrawer', () => {
      setingsRef.value.openDrawer()
    });
    // 获取缓存中的布局配置
    /*		if (Local.get('themeConfig')) {
          console.log("local1:",Local.get('themeConfig'))
          storesThemeConfig.setThemeConfig({ themeConfig: Local.get('themeConfig') });
          document.documentElement.style.cssText = Local.get('themeConfigStyle');
        }*/
    // 获取缓存中的全屏配置
    if (Session.get('isTagsViewCurrenFull')) {
      stores.setCurrenFullscreen(Session.get('isTagsViewCurrenFull'));
    }
  });
});
// 页面销毁时，关闭监听布局配置/i18n监听
onUnmounted(() => {
  mittBus.off('openSetingsDrawer', () => {
  });
});
const getThemeConfig = () => {
  storesThemeConfig.setThemeConfig({themeConfig: themeConfig.value});
  document.documentElement.style.cssText = Local.get('themeConfigStyle');
  // getThemeConfigApi({
  //   skipCache: 1
  // }).then(res => {
  //   if (res && res.code == 200 && Object.keys(res.data).length > 0) {
  //     setConfig(res.data, themeConfig.value)
  //   } else {
  //     storesThemeConfig.setThemeConfig({themeConfig: themeConfig.value});
  //     document.documentElement.style.cssText = Local.get('themeConfigStyle');
  //   }
  // })
}
const setConfig = async (newArray: any, oldArray: any) => {
  let map = new Map()
  for (let key in oldArray) {
    if (newArray[key] != null && newArray[key] != oldArray[key]) {
      map.set(key, newArray[key])
    } else {
      map.set(key, oldArray[key])
    }
  }
  const json = Object.fromEntries(map);
  storesThemeConfig.setThemeConfig({themeConfig: json});
  Local.remove('themeConfig');
  Local.set('themeConfig', themeConfig.value);
  document.documentElement.style.cssText = Local.get('themeConfigStyle');
}
// 监听路由的变化，设置网站标题
watch(
    () => route.path,
    () => {
      other.useTitle();
    },
    {
      deep: true,
    }
);
</script>
