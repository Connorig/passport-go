import { Module } from 'vuex';
import { Session, Local } from '/@/utils/storage';
// 此处加上 `.ts` 后缀报错，具体原因不详
import { UserInfosState, RootStateTypes } from '/@/store/interface/index';
import { getUserInfoByTokenApi, clearTokenApi } from '/@/api/passport/index'
import { ElMessage } from 'element-plus';

const userInfosModule: Module<UserInfosState, RootStateTypes> = {
	namespaced: true,
	state: {
		userInfos: {},
	},
	mutations: {
		// 设置用户信息
		getUserInfos(state: any, data: object) {
			state.userInfos = data;
		},
	},
	actions: {
		// 设置用户信息
		async setUserInfos({ commit }, data: object) {
			var data = {}
			let auth: Array<string> = ['base', 'admin', 'test']
			data.authPageList = auth
			data.type = Local.get("type")
			commit('getUserInfos', data)
		},
		// 清除用户信息
		async clearUserInfos({ commit }, data: object) {
			return new Promise<void>((resolve, reject) => {
				// clearTokenApi({}).then((res)=>{
				// 	Local.clear()
				// 	resolve()
				// }).catch(()=>{
				// 	Local.clear()
				// 	reject()}
				// 	)
			})
		},
	},
};

export default userInfosModule;
