import z from "zod";

export const departmentSchema = z.object({
  code: z
    .string()
    .min(1, "Code is required")
    .max(10, "Code must be at most 10 characters"),
  name: z
    .string()
    .min(1, "Name is required")
    .max(50, "Name must be at most 50 characters"),
  budgetCode: z.string().min(50, "Budget code must be at most 50 characters"),
  headOfDepartment: z
    .string()
    .uuid("invalid head of department")
    .optional()
    .nullable(),
  parentDepartment: z
    .string()
    .uuid("invalid parent department")
    .optional()
    .nullable(),
});

export const departmentResponseSchema = z.object({
  id: z.string().uuid(),
  code: z.string(),
  name: z.string(),
  headOfDepartmentId: z.string().uuid().nullable().optional(),
  parentDepartmentId: z.string().uuid().nullable().optional(),
  budgetCode: z.string().nullable().optional(),
  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),

  // Relationships
  headOfDepartment: z.any().nullable().optional(),
  parentDepartment: z.any().nullable().optional(),
  subDepartments: z.array(z.any()).optional(),
  employees: z.array(z.any()).optional(),
});

export const createDepartmentSchema = departmentSchema;

export const updateDepartmentSchema = departmentSchema.partial().extend({
  id: z.string().uuid("Invalid department ID format"),
});

export type DepartmentRequest = z.infer<typeof departmentSchema>;
export type DepartmentResponse = z.infer<typeof departmentResponseSchema>;
export type CreateDepartmentRequest = z.infer<typeof createDepartmentSchema>;
export type UpdateDepartmentRequest = z.infer<typeof updateDepartmentSchema>;
