import { getListEmployee } from "@/services/employee";
import { EmployeeFilters } from "@/types/employee";
import { useQuery } from "@tanstack/react-query";

interface UseEmployeeDataProps {
  page: number;
  limit: number;
  filters: EmployeeFilters;
}

export const useEmployeeList = ({
  page,
  limit,
  filters,
}: UseEmployeeDataProps) => {
  return useQuery({
    queryKey: ["employees", page, limit, filters],
    queryFn: () => getListEmployee(page, limit, filters),
    // keepPreviousData: true, // Keep previous data while fetching new page
    staleTime: 5000, // Data considered fresh for 5 seconds
  });
};
