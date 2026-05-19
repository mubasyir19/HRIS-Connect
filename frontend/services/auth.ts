import { api } from "@/lib/axiosInstance";
import { LoginRequest } from "@/schemas/authSchema";

export const login = async (req: LoginRequest) => {
  const res = await api.post("/auth/login", req);
  return res.data.data;
};

export const logout = async () => {
  const res = await api.post("/auth/logout");
  return res.data;
};

export const profile = async () => {};
