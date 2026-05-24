"use client";

import { ColumnDef } from "@tanstack/react-table";

type HeadOfDepartmentInfo = {
  id: string;
  fullname: string;
};

export type Department = {
  id: string;
  code: string;
  name: string;
  headOfDepartmentId: string | null;
  headOfDepartment: HeadOfDepartmentInfo | null;
  parentDepartmentId: string | null;
  budgetCode: string;
  totalEmployee: number;
  location?: string;
  status?: string;
};

export const columnsDepartment: ColumnDef<Department>[] = [
  {
    accessorKey: "name",
    header: "Department Name",
  },
  {
    accessorKey: "headOfDepartment.fullname",
    header: "Head Of Department",
    cell: ({ row }) => {
      const headOfDept = row.original.headOfDepartment;
      return (
        <div className="font-medium text-gray-700">
          {headOfDept?.fullname || "-"}
        </div>
      );
    },
  },
  {
    accessorKey: "code",
    header: "Department Code",
  },
  {
    accessorKey: "totalEmployee",
    header: "Total Employee",
    cell: ({ row }) => {
      const total = row.getValue("totalEmployee") as number;
      return <div className="font-medium text-black">{total}</div>;
    },
  },
  {
    // accessorKey: "status",
    header: "....",
  },
];
