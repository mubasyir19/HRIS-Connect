"use client";

import { useAddDepartment } from "@/hooks/department/useAddDepartment";
import { useEmployeeList } from "@/hooks/employee/useEmployeeList";
import React, { useState } from "react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Employee } from "@/types/employee";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Check, ChevronsUpDown } from "lucide-react";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "../ui/command";
import { cn } from "@/lib/utils";

export default function FormAddDepartment() {
  const [search, setSearch] = useState<string>("");
  const [openCombobox, setOpenCombobox] = useState<boolean>(false);
  const { mutate, isPending } = useAddDepartment();
  const [formData, setFormData] = useState({
    name: "",
    code: "",
    parentDepartment: "",
    headOfDepartment: "",
    budgetCode: "",
  });

  const { data: employeeData, isLoading: employeeLoading } = useEmployeeList({
    page: 1,
    limit: 100,
    filters: {
      name: search || undefined,
    },
  });

  const employees = employeeData?.data ?? [];

  const selectedEmployee = employees.find(
    (e: Employee) => e.id === formData.headOfDepartment,
  );

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSelectChange = (name: string, value: string) => {
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    const submitData = {
      code: formData.code,
      name: formData.name,
      budgetCode: formData.budgetCode,
      headOfDepartment: formData.headOfDepartment || null,
      parentDepartment: formData.parentDepartment || null,
    };

    console.log("submitData: ", submitData);
    mutate(submitData);
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <div className="space-y-2">
          <Label>Department Name</Label>
          <Input
            type="text"
            value={formData.name}
            onChange={handleChange}
            name="name"
            placeholder="e.g Tech Acquisition"
            disabled={isPending}
          />
        </div>
        <div className="space-y-2">
          <Label>Department Code</Label>
          <Input
            type="text"
            value={formData.code}
            onChange={handleChange}
            name="code"
            placeholder="e.g HR-TA-01"
            disabled={isPending}
          />
        </div>

        <div className="space-y-2">
          <Label>Head of Department</Label>
          <Popover open={openCombobox} onOpenChange={setOpenCombobox}>
            <PopoverTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={openCombobox}
                className="w-full justify-between font-normal"
                disabled={isPending}
                type="button" // ⚠️ Penting! Cegah trigger submit form
              >
                {/* Tampilkan nama employee terpilih, atau placeholder */}
                {selectedEmployee
                  ? selectedEmployee.fullname
                  : "Search employee..."}
                <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0" align="start">
              <Command shouldFilter={false}>
                {/* shouldFilter=false karena filter dilakukan via API (useEmployeeList) */}
                <CommandInput
                  placeholder="Search employee..."
                  value={search}
                  onValueChange={setSearch} // update state search → trigger re-fetch
                />
                <CommandList>
                  {employeeLoading ? (
                    <CommandEmpty>Loading...</CommandEmpty>
                  ) : employees.length === 0 ? (
                    <CommandEmpty>No employee found.</CommandEmpty>
                  ) : (
                    <CommandGroup>
                      {employees.map((emp: Employee) => (
                        <CommandItem
                          key={emp.id}
                          value={emp.id}
                          onSelect={(currentValue) => {
                            // Toggle: jika sudah dipilih → kosongkan, jika belum → set
                            handleSelectChange(
                              "headOfDepartment",
                              currentValue === formData.headOfDepartment
                                ? ""
                                : currentValue,
                            );
                            setSearch(""); // reset search setelah pilih
                            setOpenCombobox(false); // tutup popover
                          }}
                        >
                          {/* Centang jika employee ini sedang terpilih */}
                          <Check
                            className={cn(
                              "mr-2 h-4 w-4",
                              formData.headOfDepartment === emp.id
                                ? "opacity-100"
                                : "opacity-0",
                            )}
                          />
                          {emp.fullname}
                        </CommandItem>
                      ))}
                    </CommandGroup>
                  )}
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </div>

        <div className="space-y-2">
          <Label>Parent Department</Label>
          <Select
            value={formData.parentDepartment}
            name="parentDepartment"
            onValueChange={(value) =>
              handleSelectChange("parentDepartment", value)
            }
            disabled={isPending}
          >
            <SelectTrigger className="w-full">
              <SelectValue placeholder="None (Top Level)" />
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
      </div>

      <div className="space-y-2">
        <Label>Budget Code</Label>
        <Input
          type="text"
          name="budgetCode"
          onChange={handleChange}
          value={formData.budgetCode}
          placeholder="e.g. BC-2024-X"
          disabled={isPending}
        />
      </div>

      <div className="flex items-center justify-end">
        <Button type="submit">
          {isPending ? "Loading..." : "Save Department"}
        </Button>
      </div>
    </form>
  );
}
