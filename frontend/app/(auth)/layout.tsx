import ReactQueryClient from "@/providers/ReactQueryClient";
import React from "react";

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ReactQueryClient>
      <div className="">{children}</div>
    </ReactQueryClient>
  );
}
