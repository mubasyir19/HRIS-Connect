"use client";

import { columnsDepartment } from "@/components/department/columns";
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
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Building2, Plus, Users2 } from "lucide-react";
import React from "react";

export default function DepartmentPage() {
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
              <h2 className="text-primary text-2xl font-bold">8</h2>
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
        <div className="mt-4">
          <DataTableDepartment columns={columnsDepartment} data={[]} />
        </div>
      </div>
    </div>
  );
}
