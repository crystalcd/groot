import { get } from '@/http'
import type { IMenuItem } from '@/layout/types'
import type { IUser } from '../system/modal/userModel'

// 菜单
export function getMenus() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve([
        {
          "id": "home",
          "name": "首页",
          "icon": "HomeFilled",
          "path": "/home"
        },
        {
          "id": "base",
          "name": "基础管理",
          "icon": "Management",
          "children": [
            {
              "id": "article",
              "name": "文章管理",
              "path": "/base/article"
            }
          ]
        },
        {
          "id": "system",
          "name": "系统管理",
          "icon": "Management",
          "children": [
            {
              "id": "user",
              "name": "用户管理",
              "path": "/system/user"
            },
            {
              "id": "role",
              "name": "角色管理",
              "path": "/system/role"
            }
          ]
        }
      ])
    }, 2000);
  });
  return get<IMenuItem[]>('/menu')
}

// 获取当前用户信息
export function getUserInfo() {
  return get<IUser>('/auth/profile')
}

export function getTableData(api: string, params: any) {
  return get(api, params)
}

// 获取指定类型 key 的下拉选项
export function getOptions(key: string) {
  return get(`${key}/default/options`)
}

// 获取数据 autocomplete 下拉选项
export function getAutocompleteOptions(key: string, keyword: string) {
  return get(`${key}/autocomplete/options`, { keyword })
}
