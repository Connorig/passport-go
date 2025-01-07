import { get, post } from '/@/utils/request';
import baseURL from '..';

const secondaryURL = '/send'
/**
 * SkU
 * @param params 要传的参数值
 * @returns 返回接口数据
 */
export const signInSendCodeApi = (p: object) => post(baseURL+secondaryURL +"/signInSendCode", p);
export const signUpSendCodeApi = (p: object) => post(baseURL+secondaryURL +"/signUpSendCode", p);

