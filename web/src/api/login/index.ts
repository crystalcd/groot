import { post } from '@/http'

export interface ILogin {
  mobile: number | string
  password: string
  code?: string | number
}

enum URL {
  login = '/auth/login',
  logout = '/auth/logout'
}

export function login(params: ILogin) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      const randomNumber = Math.random();
      resolve({ "data": { "token": "test" }, "status": { "code": 0, "message": "success" } }); // 异步操作成功，传递结果
    }, 2000);
  });
  return post(URL.login, params)
}

export function logout() {
  return post(URL.logout)
}
