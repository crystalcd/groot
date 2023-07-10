import axios, { AxiosInstance, AxiosResponse, AxiosError } from 'axios';

class ApiService {
  private axiosInstance: AxiosInstance;

  constructor(baseURL: string) {
    this.axiosInstance = axios.create({
      baseURL: baseURL
    });

    // 添加拦截器等其他配置
    this.axiosInstance.interceptors.response.use(
      this.handleResponse,
      this.handleError
    );
  }

  private handleResponse(response: AxiosResponse) {
    // 处理响应数据
    return response.data;
  }

  private handleError(error: AxiosError) {
    // 处理请求错误
    throw error;
  }

  public async get<T>(url: string): Promise<T> {
    try {
      const response = await this.axiosInstance.get<T>(url);
      return response;
    } catch (error) {
      throw error;
    }
  }

  public async post<T>(url: string, data: any): Promise<T> {
    try {
      const response = await this.axiosInstance.post<T>(url, data);
      return response;
    } catch (error) {
      throw error;
    }
  }

  // 添加其他请求方法，如 put、delete 等

  // 可以根据需要定义更多自定义方法和拦截器等

}

export default ApiService;
