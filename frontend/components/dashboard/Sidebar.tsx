"use client";

import {
  Building,
  Building2,
  Calendar,
  CalendarX,
  ChartColumnBig,
  LayoutDashboard,
  LogOut,
  Settings,
  UsersRound,
} from "lucide-react";
import SidebarItem from "./SidebarItem";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";

export default function Sidebar() {
  return (
    <aside className="border-border hidden h-screen w-64 overflow-y-auto border-r-2 py-4 md:block">
      <div className="flex items-center gap-3 px-4">
        <div className="bg-primary flex size-10 items-center justify-center rounded-md">
          <Building className="size-6 text-white" />
        </div>
        <div className="">
          <h4 className="text-lg font-bold text-black uppercase">
            HRIS Portal
          </h4>
          <p className="text-secondary text-sm">Management Suite</p>
        </div>
      </div>
      <ul className="mt-8 flex flex-col gap-2">
        <SidebarItem
          link="/dashboard"
          label="Dashboard"
          icon={LayoutDashboard}
        />
        <SidebarItem link="/employee" label="Employee" icon={UsersRound} />
        <SidebarItem link="/department" label="Department" icon={Building2} />
        <SidebarItem
          link="/leave-request"
          label="Leave Requests"
          icon={CalendarX}
        />
        <SidebarItem link="/calendar" label="Calendar" icon={Calendar} />
        <SidebarItem link="/reports" label="Reports" icon={ChartColumnBig} />
        <SidebarItem link="/settings" label="Settings" icon={Settings} />
        <Dialog>
          <DialogTrigger asChild>
            <button
              className={`group hover:bg-primary/5 hover:border-primary flex items-center gap-3 border-l-4 border-transparent px-4 py-2 transition-all duration-200`}
            >
              <LogOut
                className={`text-secondary group-hover:text-primary size-6 transition-all duration-200`}
              />
              <p
                className={`text-secondary group-hover:text-primary font-medium transition-all duration-200`}
              >
                Logout
              </p>
            </button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle></DialogTitle>
              <DialogDescription></DialogDescription>
            </DialogHeader>
            <p>logout dialog</p>
          </DialogContent>
        </Dialog>
      </ul>
    </aside>
  );
}
