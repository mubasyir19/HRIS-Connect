"use client";

import {
  Bell,
  CircleQuestionMark,
  LogOut,
  Search,
  Settings,
  User,
} from "lucide-react";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Button } from "../ui/button";
import { useLogout } from "@/hooks/auth/useLogout";

export default function Navbar() {
  const { mutate, isPending } = useLogout();

  const handleLogout = () => {
    mutate();
  };
  return (
    <nav className="border-border flex items-center justify-between border-b-2 bg-white px-6 py-3">
      <div className="">
        <div className="flex items-center gap-3 rounded-md bg-slate-100 px-4 py-2">
          <Search className="text-secondary size-4" />
          <input
            type="text"
            name="search"
            placeholder="Search employees, document, reports"
            className="w-full border-none text-sm outline-none placeholder:text-sm"
          />
        </div>
      </div>
      <div className="md:hidden">
        <div className="bg-primary size-7 rounded-full"></div>
      </div>
      <div className="hidden items-center gap-6 md:flex">
        <button className="relative cursor-pointer">
          <Bell className="text-secondary size-5" />
          <div className="absolute top-0 right-0 size-0.5 rounded-full bg-red-500 p-0.5"></div>
        </button>
        <button className="cursor-pointer">
          <CircleQuestionMark className="text-secondary size-5" />
        </button>
        <button className="cursor-pointer">
          <Settings className="text-secondary size-5" />
        </button>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div className="flex items-center gap-3">
              <div className="">
                <p className="text-end text-sm font-semibold text-black">
                  Admin HRIS
                </p>
                <p className="text-secondary text-end text-xs uppercase">
                  ADMIN
                </p>
              </div>
              <div className="bg-primary size-8 rounded-lg"></div>
            </div>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
              <User /> Profile
            </DropdownMenuItem>
            <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
              <Dialog>
                <DialogTrigger asChild>
                  <Button size={"sm"} variant={"outline"} className="w-full">
                    <LogOut /> Logout
                  </Button>
                </DialogTrigger>
                <DialogContent>
                  <DialogHeader>
                    <DialogTitle>Confirm Logout</DialogTitle>
                    <DialogDescription>
                      Are you sure you want to log out of your account?
                    </DialogDescription>
                  </DialogHeader>
                  <div className="mt-4 flex justify-end gap-3">
                    <DialogClose asChild>
                      <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button variant="destructive" onClick={handleLogout}>
                      {isPending ? "Loading..." : "Logout"}
                    </Button>
                  </div>
                </DialogContent>
              </Dialog>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </nav>
  );
}
