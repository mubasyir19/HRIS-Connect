"use client";

import React, { useState } from "react";
import { Input } from "../ui/input";
import { useLogin } from "@/hooks/auth/useLogin";

interface FormLogin {
  email: string;
  password: string;
}

export default function FormLogin() {
  const { mutate, isPending, error } = useLogin();

  const [formData, setFormData] = useState<FormLogin>({
    email: "",
    password: "",
  });

  const handleChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.SubmitEvent) => {
    e.preventDefault();
    console.log("form data login = ", formData);

    mutate(formData);
  };
  return (
    <form onSubmit={handleSubmit} className="space-y-3 md:mt-10">
      <div className="input-group">
        <label htmlFor="email" className="block">
          Email
        </label>
        <Input
          type="email"
          name="email"
          placeholder="admin@example.com"
          value={formData.email}
          onChange={handleChange}
          required
        />
      </div>
      <div className="input-group">
        <label htmlFor="" className="block">
          Password
        </label>
        <Input
          type="password"
          name="password"
          placeholder="********"
          value={formData.password}
          onChange={handleChange}
        />
      </div>
      <button
        type="submit"
        className="bg-primary w-full rounded-sm px-4 py-2 text-center font-medium text-white"
      >
        {isPending ? "Loading..." : "Login"}
      </button>
      <div className="bg-primary/10 flex items-start justify-center gap-3 rounded-md border border-gray-400 p-4">
        <div className="bg-primary mt-1 flex size-5 shrink-0 items-center justify-center rounded-full">
          <p className="text-xs font-medium text-white">i</p>
        </div>
        <div className="">
          <p className="text-base text-black">Demo Environment Access</p>
          <p className="mt-2 text-sm text-black">
            Use admin@example.com / password to explore the management suite.
          </p>
        </div>
      </div>
    </form>
  );
}
