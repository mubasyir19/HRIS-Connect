import { api } from "@/lib/axiosInstance";
import { DepartmentRequest } from "@/schemas/departmentSchema";

export const getAllDepartment = async (
  filter,
  page: number,
  limit: number,
) => {};

export const getByIdDepartment = async (id: string) => {};

export const addNewDepartment = async (req: DepartmentRequest) => {
  const res = await api.post("/department/add", req);
  return res.data;
};

export const updateDepartment = async (id: string) => {};

export const deleteDepartment = async (id: string) => {};
