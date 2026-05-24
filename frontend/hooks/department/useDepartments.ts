import { getAllDepartment } from "@/services/department";
import { getListEmployee } from "@/services/employee";
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
    // queryFn: async () => {
    //   // Fetch departments and employees in parallel
    //   const [deptRes, empRes] = await Promise.all([
    //     getAllDepartment({}, page, limit),
    //     getListEmployee({}, 1, 10000), // Get all employees
    //   ]);

    //   // Count employees per department
    //   const employeeCount = new Map();
    //   empRes.data.forEach((emp: any) => {
    //     const deptId = emp.departmentId;
    //     employeeCount.set(deptId, (employeeCount.get(deptId) || 0) + 1);
    //   });

    //   // Enrich departments with employee count
    //   const enrichedData = deptRes.data.map((dept: any) => ({
    //     ...dept,
    //     totalEmployees: employeeCount.get(dept.id) || 0,
    //   }));

    //   return {
    //     ...deptRes,
    //     data: enrichedData,
    //   };
    // },
  });
}
