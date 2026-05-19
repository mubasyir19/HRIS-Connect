"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Plus } from "lucide-react";
import React from "react";

export default function EmployeePage() {
  return (
    <div>
      <div className="flex items-center justify-between">
        <div className="">
          <h1 className="text-3xl font-bold text-black">Employee Management</h1>
          <p className="text-secondary mt-2 text-sm">
            Manage your departments.
          </p>
        </div>
        <div className="flex items-center gap-3">
          <Dialog>
            <DialogTrigger asChild>
              <Button size={"sm"}>
                <Plus />
                Add Employee
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Add New Employee</DialogTitle>
                <DialogDescription>
                  Enter employee credentials and contract details.
                </DialogDescription>
              </DialogHeader>
              <form>
                <p>this is form add new employee</p>
              </form>
            </DialogContent>
          </Dialog>
        </div>
      </div>
    </div>
  );
}
