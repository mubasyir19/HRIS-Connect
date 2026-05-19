"use client";

import { ColumnDef } from "@tanstack/react-table";

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type Department = {
  id: string;
  code: string;
  name: string;
  headOfDepartmentID: string;
  parentDepartmentID: string;
  budgetCode: string;
  totalEmployee: string;
  location: string;
  status: string;
};

export const columnsDepartment: ColumnDef<Department>[] = [
  {
    accessorKey: "name",
    header: "Department Name",
  },
  {
    accessorKey: "headOfDepartmentID",
    header: "Head Of Department",
  },
  {
    accessorKey: "totalEmployee",
    header: "Total Employee",
  },
  {
    accessorKey: "location",
    header: "Location",
  },
  {
    accessorKey: "status",
    header: "Status",
  },
];
