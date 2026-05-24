"use client";

import {
  columnsDepartment,
  type Department,
} from "@/components/department/columns";
import { DataTableDepartment } from "@/components/department/data-table";
import FormAddDeparment from "@/components/department/FormAddDeparment";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useDepartments } from "@/hooks/department/useDepartments";
import { DepartmentFilters } from "@/types/departments";
import { Building2, Plus, Users2 } from "lucide-react";
import React, { useState } from "react";

export default function DepartmentPage() {
  const [page, setPage] = useState<number>(1);
  const [searchFilter, setSearchFilter] = useState<string>("");
  const limit = 10;

  const filters: DepartmentFilters = searchFilter.trim()
    ? { name: searchFilter.trim() }
    : {};

  const {
    data: ListDepartments,
    isPending,
    error,
  } = useDepartments({ filters, page, limit });

  console.log("department = ", ListDepartments);

  const departments = (ListDepartments?.data ?? []) as Department[];
  const meta = ListDepartments?.meta;
  const totalDepartments = meta?.total ?? departments.length;
  const totalPages = meta?.total_pages ?? 1;
  const currentPage = meta?.page ?? page;

  return (
    <div>
      <div className="flex items-center justify-between">
        <div className="">
          <h1 className="text-3xl font-bold text-black">Department</h1>
          <p className="text-secondary mt-2 text-sm">
            Manage your workforce, roles, and leave permissions.
          </p>
        </div>
        <div className="flex items-center gap-3">
          <Dialog>
            <DialogTrigger asChild>
              <Button size={"sm"}>
                <Plus />
                Add Department
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-106.25 md:max-w-1/2">
              <DialogHeader>
                <DialogTitle>Add New Department</DialogTitle>
                <DialogDescription>
                  Enter employee credentials and contract details.
                </DialogDescription>
              </DialogHeader>
              <FormAddDeparment />
            </DialogContent>
          </Dialog>
        </div>
      </div>
      <div className="mt-8 grid grid-cols-1 gap-6 md:grid-cols-3">
        <div className="border-border rounded-md border bg-white p-6">
          <div className="flex items-center justify-between">
            <div className="space-y-2">
              <p className="text-sm font-semibold text-black uppercase">
                Total Departments
              </p>
              <h2 className="text-primary text-2xl font-bold">
                {totalDepartments}
              </h2>
            </div>
            <div className="">
              <div className="bg-primary/10 flex size-20 items-center justify-center rounded-lg">
                <Building2 className="text-primary size-10" />
              </div>
            </div>
          </div>
        </div>
        <div className="border-border rounded-md border bg-white p-6">
          <div className="flex items-center justify-between">
            <div className="space-y-2">
              <p className="text-sm font-semibold text-black uppercase">
                Total Employees
              </p>
              <h2 className="text-primary text-2xl font-bold">64</h2>
              <p className="text-xs text-black">Across all active sectors</p>
            </div>
            <div className="">
              <div className="bg-primary/10 flex size-20 items-center justify-center rounded-lg">
                <Users2 className="text-primary size-10" />
              </div>
            </div>
          </div>
        </div>
        <div className="border-border rounded-md border bg-white p-6"></div>
      </div>
      <div className="border-border mt-8 rounded-md border bg-white p-6">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <Select>
              <SelectTrigger>
                <SelectValue placeholder="All Locations" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value="light">Light</SelectItem>
                  <SelectItem value="dark">Dark</SelectItem>
                  <SelectItem value="system">System</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <Select>
              <SelectTrigger>
                <SelectValue placeholder="Sort by Name" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value="light">Light</SelectItem>
                  <SelectItem value="dark">Dark</SelectItem>
                  <SelectItem value="system">System</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>
          <div className="">
            <Input
              name="searchFilter"
              value={searchFilter}
              onChange={(e) => {
                setPage(1);
                setSearchFilter(e.target.value);
              }}
              placeholder="Search name, head of, or parent department"
            />
          </div>
        </div>
        <div className="mt-4">
          <DataTableDepartment columns={columnsDepartment} data={departments} />
          {error ? (
            <p className="text-destructive mt-4 text-sm">
              Gagal memuat daftar departemen. Silakan coba lagi.
            </p>
          ) : null}
        </div>

        <div className="mt-4 flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
          <p className="text-secondary text-sm">
            Page {currentPage} of {totalPages} • {totalDepartments} departments
          </p>
          <div className="flex items-center gap-2">
            <Button
              size="sm"
              variant="outline"
              disabled={currentPage <= 1 || isPending}
              onClick={() => setPage((prev) => Math.max(1, prev - 1))}
            >
              Previous
            </Button>
            <Button
              size="sm"
              variant="outline"
              disabled={currentPage >= totalPages || isPending}
              onClick={() => setPage((prev) => Math.min(totalPages, prev + 1))}
            >
              Next
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
