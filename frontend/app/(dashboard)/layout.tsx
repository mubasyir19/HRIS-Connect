import Navbar from "@/components/dashboard/Navbar";
import Sidebar from "@/components/dashboard/Sidebar";
import ReactQueryClient from "@/providers/ReactQueryClient";
import React from "react";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ReactQueryClient>
      <div className="flex min-h-full w-full flex-row flex-nowrap items-start">
        <Sidebar />
        <main className="min-h-screen w-full bg-[#FAF8FF]">
          <Navbar />
          <div className="px-4 py-4 md:px-6">{children}</div>
        </main>
      </div>
    </ReactQueryClient>
  );
}
