import React from "react";

export default function ProtectedProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  return <div>{children}</div>;
}
