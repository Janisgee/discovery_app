import { LoadingSpinner } from "@/app/_ui/loading-spinner";

export default function loading() {
  return (
    <div className="loading fixed inset-0 z-50 flex items-center justify-center">
      <LoadingSpinner size={64} color="text-violet-600" />
    </div>
  );
}
