import { login } from "@/services/auth";
import { QueryClient, useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

export function useLogin() {
  const queryClient = new QueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: login,
    onSuccess: (data) => {
      queryClient.setQueryData(["user"], data.data);
      toast.success("Login successfully");
      router.push("/dashboard");
    },
  });
}
