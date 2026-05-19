import { getToken } from "@/helpers/getToken";
import axios, { AxiosError, type AxiosRequestConfig } from "axios";

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API || "http://localhost:8080/api/v1",
  withCredentials: true,
});

api.interceptors.request.use((config) => {
  const accessToken = getToken();
  if (accessToken) {
    config.headers = config.headers ?? {};
    config.headers["X-Access-Token"] = accessToken;
  }

  return config;
});

api.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    const originalRequest = error.config as AxiosRequestConfig & {
      _retry?: boolean;
    };

    if (!originalRequest || originalRequest._retry) {
      return Promise.reject(error);
    }

    if (error.response?.status === 401) {
      if (originalRequest.url?.includes("/auth/refresh")) {
        return Promise.reject(error);
      }

      originalRequest._retry = true;

      try {
        await api.post("/auth/refresh", null, {
          withCredentials: true,
        });

        return api(originalRequest);
      } catch (refreshError) {
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  },
);
