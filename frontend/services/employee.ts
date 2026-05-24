import { api } from "@/lib/axiosInstance";
import { EmployeeFilters, EmployeeResponse } from "@/types/employee";

export const getListEmployee = async (
  page: number = 1,
  limit: number = 10,
  filters?: EmployeeFilters,
) => {
  const params: Record<string, any> = {
    page,
    limit,
  };

  if (filters) {
    if (filters.name) params.name = filters.name;
    if (filters.departmentId) params.departmentId = filters.departmentId;
    if (filters.position) params.position = filters.position;
    if (filters.status) params.status = filters.status;
  }

  const res = await api.get("/employee/all", { params });
  return res.data;
};

export const getListEmployeeWithQuery = async (
  page: number = 1,
  limit: number = 10,
  search?: string,
): Promise<EmployeeResponse> => {
  const queryString = new URLSearchParams({
    page: page.toString(),
    limit: limit.toString(),
    ...(search && { search }),
  }).toString();

  const res = await api.get(`/employee/all?${queryString}`);
  return res.data;
};
