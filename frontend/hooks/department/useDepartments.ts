import { getAllDepartment } from "@/services/department";
import { DepartmentFilters } from "@/types/departments";
import { useQuery } from "@tanstack/react-query";

interface DepartmentMeta {
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

interface DepartmentListResponse {
  data: unknown[];
  meta: DepartmentMeta;
}

interface UseDepartmentListDataProps {
  filters?: DepartmentFilters;
  page: number;
  limit: number;
}

export function useDepartments({
  page,
  limit,
  filters,
}: UseDepartmentListDataProps) {
  return useQuery<DepartmentListResponse>({
    queryKey: ["departments", filters, page, limit],
    queryFn: () => getAllDepartment(page, limit, filters),
    staleTime: 5000,
  });
}
