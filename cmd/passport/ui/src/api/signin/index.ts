
import { get, post } from '/@/utils/request';
import baseURL from '..';

const secondaryURL = '/signin'
/**
 * SkU
 * @param params 要传的参数值
 * @returns 返回接口数据
 */
export const signInPasswordApi = (p: object) => post(baseURL+secondaryURL +"/", p);
export const signInPhoneApi = (p: object) => post(baseURL+secondaryURL +"/sms", p);
export const signInEmailApi = (p: object) => post(baseURL+secondaryURL +"/email", p);

