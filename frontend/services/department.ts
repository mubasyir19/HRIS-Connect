import { api } from "@/lib/axiosInstance";
import { DepartmentRequest } from "@/schemas/departmentSchema";
import { DepartmentFilters } from "@/types/departments";

export const getAllDepartment = async (
  page: number = 1,
  limit: number = 10,
  filters?: DepartmentFilters,
) => {
  const params: Record<string, string | number | undefined> = {
    page,
    limit,
  };

  if (filters) {
    if (filters.code) params.code = filters.code;
    if (filters.name) params.name = filters.name;
    if (filters.headOfDepartmentId)
      params.headOfDepartmentId = filters.headOfDepartmentId;
    if (filters.parentOfDepartmentId)
      params.parentOfDepartmentId = filters.parentOfDepartmentId;
  }

  const res = await api.get("/department/all", {
    params,
  });
  return res.data;
};

export const getByIdDepartment = async (_id: string) => {
  void _id;
  throw new Error("Not implemented");
};

export const addNewDepartment = async (req: DepartmentRequest) => {
  const res = await api.post("/department/add", req);
  return res.data;
};

export const updateDepartment = async (_id: string) => {
  void _id;
  throw new Error("Not implemented");
};

export const deleteDepartment = async (_id: string) => {
  void _id;
  throw new Error("Not implemented");
};
