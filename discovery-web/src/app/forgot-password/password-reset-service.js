import { useEffect, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";

export function usePasswordReset() {
  const router = useRouter();
  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState("");

  function doReset(email) {
    if (isPending) return;

    const parts = email.split("@");

    if (parts.length !== 2) {
      setError("Enter a valid email");
      return;
    }

    setIsPending(true);
    setError("");

    fetch(`${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/forgetPassword`, {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email }),
    })
      .then((response) => {
        if (response.ok) {
          router.push(`/forgot-password/sent`);
        } else {
          setError(`Failed to reset: ${response.statusText}`);
        }
      })
      .catch(() => {
        setError("Something went wrong, please try again");
      })
      .finally(() => {
        setIsPending(false);
      });
  }

  return [doReset, isPending, error];
}

export function useSetNewPassword() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const pwResetCode = searchParams.get("evpw");

  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!pwResetCode) {
      router.push("/forgot-password/expired");
    }
  }, [pwResetCode, router]);

  function doSetNewPassword(newPw, confirmPw) {
    if (!newPw || !confirmPw) {
      setError("Enter and confirm new password");
    }
    if (newPw !== confirmPw) {
      setError("Passwords do not match");
      return;
    }
    setIsPending(true);
    setError("");

    fetch(`${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/resetPassword`, {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        newPw: newPw,
        confirmPw: confirmPw,
        pwResetCode: pwResetCode,
      }),
    })
      .then((response) => {
        if (response.ok) {
          router.push("/forgot-password/success");
        } else if (response.status === 401) {
          router.push("/forgot-password/expired");
        } else {
          return response.text().then((msg) => setError(msg));
        }
      })
      .catch(() => {
        setError("An error occurred, please try again");
      })
      .finally(() => {
        setIsPending(false);
      });
  }

  return [doSetNewPassword, isPending, error];
}
