import { logout } from "@/services/auth";
import { QueryClient, useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";

export function useLogout() {
  const queryClient = new QueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: logout,
    onSuccess: () => {
      queryClient.clear();
      router.push("/login");
    },
    onError: (error) => {
      console.log("error", error);
      queryClient.clear();
      router.push("/login");
    },
  });
}
