"use client";

import React, { useState } from "react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Textarea } from "../ui/textarea";
import { Button } from "../ui/button";
import { useAddDepartment } from "@/hooks/department/useAddDepartment";

export default function FormAddDeparment() {
  const { mutate, isPending } = useAddDepartment();
  const [formData, setFormData] = useState({
    name: "",
    code: "",
    parentDepartment: "",
    headOfDepartment: "",
    budgetCode: "",
  });

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.SubmitEvent) => {
    e.preventDefault();
    console.log("formData: ", formData);

    mutate(formData);
  };

  const handleSelectChange = (name: string, value: string) => {
    setFormData((prev) => ({ ...prev, [name]: value }));
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
          <Label>Head of Deparment</Label>
          <Select
            value={formData.headOfDepartment}
            name="headOfDepartment"
            onValueChange={(value) =>
              handleSelectChange("headOfDepartment", value)
            }
            disabled={isPending}
          >
            <SelectTrigger className="w-full">
              <SelectValue placeholder="Select Employee" />
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
