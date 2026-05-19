import { LucideIcon } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

interface SidebarItemProps {
  label: string;
  link: string;
  icon: LucideIcon;
}

export default function SidebarItem({
  label,
  link,
  icon: Icon,
}: SidebarItemProps) {
  const pathname = usePathname();
  console.log("path route = ", pathname);
  const isActive = pathname === link;

  return (
    <Link href={link}>
      <li
        className={`group flex items-center gap-3 border-l-4 px-4 py-2 transition-all duration-200 ${isActive ? "border-primary bg-primary/5" : "hover:bg-primary/5 hover:border-primary border-transparent"}`}
      >
        <Icon
          className={`size-6 transition-all duration-200 ${isActive ? "text-primary" : "text-secondary group-hover:text-primary"}`}
        />
        <p
          className={`font-medium transition-all duration-200 ${isActive ? "text-primary" : "text-secondary group-hover:text-primary"}`}
        >
          {label}
        </p>
      </li>
    </Link>
  );
}
