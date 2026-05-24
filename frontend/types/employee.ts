export interface Employee {
  id: string;
  fullname: string;
  email: string;
  position: string;
  departmentId: string;
  // tambahkan field lain sesuai model Employee Anda
}

export interface MetaData {
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

export interface EmployeeResponse {
  data: Employee[];
  meta: MetaData;
}

export interface EmployeeFilters {
  name?: string;
  departmentId?: string;
  position?: string;
  status?: string;
}
