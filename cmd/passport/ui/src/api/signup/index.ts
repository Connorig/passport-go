
import { get, post } from '/@/utils/request';
import baseURL from '..';

const secondaryURL = '/signup'
/**
 * SkU
 * @param params 要传的参数值
 * @returns 返回接口数据
 */
export const signUpPasswordApi = (p: object) => post(baseURL+secondaryURL +"/", p);
export const signUpPhoneApi = (p: object) => post(baseURL+secondaryURL +"/phoneRegister", p);
export const signUpEmailApi = (p: object) => post(baseURL+secondaryURL +"/emailRegister", p);

