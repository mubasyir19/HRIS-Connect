import { addNewDepartment } from "@/services/department";
import { QueryClient, useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { toast } from "sonner";

interface ApiErrorResponse {
  message: string;
  statusCode?: number;
  error?: string;
}

export function useAddDepartment() {
  const queryClient = new QueryClient();

  return useMutation({
    mutationFn: addNewDepartment,
    onSuccess: () => {
      toast.success("successfully add new department");
      queryClient.invalidateQueries({ queryKey: ["departments"] });
    },
    onError: (error: AxiosError<ApiErrorResponse>) => {
      toast.error(error?.response?.data?.message || "Failed to add department");
      console.error("Error adding department:", error);
    },
  });
}
