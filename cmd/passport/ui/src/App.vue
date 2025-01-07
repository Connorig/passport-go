<template>
<!--	<el-config-provider :locale="i18nLocale">-->
  <div class="main-class">
    <router-view v-show="getThemeConfig.lockScreenTime !== 0" />
  </div>

<!--		<LockScreen v-if="getThemeConfig.isLockScreen" />-->
<!--		<Setings ref="setingsRef" v-show="getThemeConfig.lockScreenTime !== 0" />-->
<!--		<CloseFull />-->
<!--	</el-config-provider>-->
</template>

<script lang="ts">
import { computed, ref, getCurrentInstance, onBeforeMount, onMounted, onUnmounted, nextTick, defineComponent, watch, reactive, toRefs } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from '/@/store/index';
import { useTitle } from '/@/utils/setWebTitle';
import { Local, Session } from '/@/utils/storage';
import setIntroduction from '/@/utils/setIconfont';
export default defineComponent({
	name: 'app',
	setup() {
		const { proxy } = getCurrentInstance() as any;
		const setingsRef = ref();
		const route = useRoute();
		const store = useStore();
		const title = useTitle();
		const state = reactive({
			i18nLocale: null,
		});
		// 获取布局配置信息
		const getThemeConfig = computed(() => {
			return store.state.themeConfig.themeConfig;
		});
		// 布局配置弹窗打开
		const openSetingsDrawer = () => {
			setingsRef.value.openDrawer();
		};
    const updateThemeConfig = () => {
      document.documentElement.style.cssText = Local.get('themeConfigStyle');
    };
    const setConfig = async (newArray: any) => {
      let map = new Map()
      let oldArray = getThemeConfig.value
      for (let key in oldArray) {
        if (newArray[key] != null && newArray[key] != oldArray[key]) {
          map.set(key, newArray[key])
        } else {
          map.set(key, oldArray[key])
        }
      }
      const json = Object.fromEntries(map);
      store.dispatch('themeConfig/setThemeConfig', json);
      Local.remove('themeConfig');
      Local.set('themeConfig', getThemeConfig.value);
      console.log("[setConfig]",getThemeConfig.value)
      document.documentElement.style.cssText = Local.get('themeConfigStyle');
    }
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
        updateThemeConfig()
				// 监听布局配置弹窗点击打开
				proxy.mittBus.on('openSetingsDrawer', () => {
					openSetingsDrawer();
				});
				// 设置 i18n，App.vue 中的 el-config-provider
				proxy.mittBus.on('getI18nConfig', (locale: string) => {
					state.i18nLocale = locale;
				});
				// 获取缓存中的布局配置
				if (Local.get('themeConfig')) {
					store.dispatch('themeConfig/setThemeConfig', Local.get('themeConfig'));
					document.documentElement.style.cssText = Local.get('themeConfigStyle');
				}
				// 获取缓存中的全屏配置
				if (Session.get('isTagsViewCurrenFull')) {
					store.dispatch('tagsViewRoutes/setCurrenFullscreen', Session.get('isTagsViewCurrenFull'));
				}
			});
		});
		// 页面销毁时，关闭监听布局配置/i18n监听
		onUnmounted(() => {
			proxy.mittBus.off('openSetingsDrawer', () => {});
			proxy.mittBus.off('getI18nConfig', () => {});
		});
		// 监听路由的变化，设置网站标题
		watch(
			() => route.path,
			() => {
				title();
			}
		);
		return {
			setingsRef,
			getThemeConfig,
			...toRefs(state),
		};
	},
});
</script>
<style scoped lang="scss">
.main-class{
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

</style>
