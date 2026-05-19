import { Button } from "@/components/ui/button";
import {
  Banknote,
  CalendarCheck,
  CalendarDays,
  ChartNoAxesCombined,
  Clock,
  Contact,
  FileText,
  Plus,
  Users,
} from "lucide-react";
import React from "react";

export default function DashboardPage() {
  return (
    <div>
      <div className="flex flex-col items-start gap-4 md:flex-row md:items-center md:justify-between">
        <div className="">
          <h1 className="text-3xl font-bold text-black">
            Welcome Back, Admin!
          </h1>
          <p className="text-secondary mt-2 text-sm">
            Today is Tuesday, May 21th, 2026
          </p>
        </div>
        <div className="flex items-center gap-3">
          <Button size={"sm"}>Export Data</Button>
          <Button size={"sm"}>
            <Plus />
            Add Employee
          </Button>
        </div>
      </div>
      <div className="mt-8 grid grid-cols-1 gap-6 md:grid-cols-3 lg:grid-cols-4">
        <div
          id="card-dash"
          className="border-border flex items-start justify-between rounded-md border bg-white p-4 xl:p-6"
        >
          <div className="space-y-2">
            <p className="text-sm font-medium text-black uppercase">
              Total Employee
            </p>
            <h2 className="text-2xl font-bold text-black">150</h2>
            <p className="text-xs text-green-500">+12% from last month</p>
          </div>
          <div className="">
            <Users className="text-secondary size-7" />
          </div>
        </div>
        <div
          id="card-dash"
          className="border-border flex items-start justify-between rounded-md border bg-white p-4 xl:p-6"
        >
          <div className="space-y-2">
            <p className="text-sm font-medium text-black uppercase">
              Present Today
            </p>
            <h2 className="text-2xl font-bold text-black">142</h2>
            <div className="relative h-1 w-full rounded-full bg-gray-300">
              <div className="bg-primary absolute left-0 h-1 w-[75%] rounded-full"></div>
            </div>
            <p className="text-xs text-green-500">94.6% occupancy rate</p>
          </div>
          <div className="">
            <CalendarCheck className="text-secondary size-7" />
          </div>
        </div>
        <div
          id="card-dash"
          className="border-border flex items-start justify-between rounded-md border bg-white p-4 xl:p-6"
        >
          <div className="space-y-2">
            <p className="text-sm font-medium text-black uppercase">
              Pending Leave
            </p>
            <h2 className="text-2xl font-bold text-red-900">5</h2>
            <p className="text-xs text-red-900">Action required</p>
          </div>
          <div className="">
            <Clock className="size-7 text-red-900" />
          </div>
        </div>
        <div
          id="card-dash"
          className="border-border flex items-start justify-between rounded-md border bg-white p-4 xl:p-6"
        >
          <div className="space-y-2">
            <p className="text-sm font-medium text-black uppercase">
              Avg Leave
            </p>
            <h2 className="text-2xl font-bold text-black">12 days</h2>
            <p className="text-xs text-black">Per employee /year</p>
          </div>
          <div className="">
            <CalendarDays className="text-secondary size-7" />
          </div>
        </div>
      </div>
      <div className="mt-8 grid grid-cols-1 gap-6 md:grid-cols-5 lg:grid-cols-8">
        <div className="border-border rounded-md border bg-white p-6 md:col-span-3 lg:col-span-6">
          <div className="flex items-center justify-between">
            <div className="">
              <h3 className="text-xl font-bold text-black">Attendance Trend</h3>
              <p className="text-secondary text-sm">
                Daily presence over the last 7 days
              </p>
            </div>
            <div className=""></div>
          </div>
        </div>
        <div className="border-border w-full shrink-0 rounded-md border bg-white p-6 md:col-span-1 lg:col-span-2">
          <h3 className="text-xl font-bold text-black">Recent Activity</h3>
        </div>
      </div>
      <div className="mt-8 grid grid-cols-1 gap-6 md:grid-cols-2">
        <div className="border-border rounded-md border bg-white p-6">
          <h3 className="text-xl font-bold text-black">
            Recent Leave Requests
          </h3>
        </div>
        <div className="border-border rounded-md border bg-white p-6">
          <h3 className="text-xl font-bold text-black">Quick Actions</h3>
          <div className="mt-4 grid grid-cols-1 gap-3 md:grid-cols-2">
            <div className="flex items-center justify-center rounded-lg border border-dashed border-gray-300 p-4">
              <div className="space-y-2">
                <Contact className="text-primary mx-auto size-8" />
                <p className="text-center font-semibold text-black">
                  Onboard New
                </p>
              </div>
            </div>
            <div className="flex items-center justify-center rounded-lg border border-dashed border-gray-300 p-4">
              <div className="space-y-2">
                <FileText className="text-primary mx-auto size-8" />
                <p className="text-center font-semibold text-black">
                  Bulk Upload
                </p>
              </div>
            </div>
            <div className="flex items-center justify-center rounded-lg border border-dashed border-gray-300 p-4">
              <div className="space-y-2">
                <Banknote className="text-primary mx-auto size-8" />
                <p className="text-center font-semibold text-black">
                  Run Payroll
                </p>
              </div>
            </div>
            <div className="flex items-center justify-center rounded-lg border border-dashed border-gray-300 p-4">
              <div className="space-y-2">
                <ChartNoAxesCombined className="text-primary mx-auto size-8" />
                <p className="text-center font-semibold text-black">
                  Analytics
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
