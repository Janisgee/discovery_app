"use client";

import { useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { Button } from "@/app/ui/buttons";

export default function ForgetPassword() {
  const router = useRouter();

  const searchParams = useSearchParams();
  const pwResetCode = searchParams.get("evpw");

  const handleInputEmail = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    const input_newPw = formData.get("new_password");
    const input_confirmPw = formData.get("confirm_password");
    if (!input_newPw || !input_confirmPw) {
      alert("Please enter fill in a new password.");
      return;
    }
    if (input_newPw != input_confirmPw) {
      alert("Please enter same password into the confirm password field.");
      return;
    }
    fetchResetNewPw(input_newPw, input_confirmPw, pwResetCode);
  };

  const fetchResetNewPw = async (newPw, confirmPw, pwResetCode) => {
    const data = {
      newPw: newPw,
      confirmPw: confirmPw,
      pwResetCode: pwResetCode,
    };
    const request = new Request("http://localhost:8080/api/resetPassword", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request);
      console.log(response.status);
      if (!response.ok) {
        // Handle error response, check for Unauthorized status
        if (response.status == 401) {
          router.push("/unauthorized-link");
          return;
        } else {
          // For other errors, handle as text (non-JSON responses)
          const errorText = await response.text();
          throw new Error(`Failed to fetch: ${errorText}`);
        }
      }

      const responseData = await response.json();
      console.log("Server Response:", responseData);

      router.push(`/resetPwSuccess`);
    } catch (error) {
      console.error("Error fetching reset password page:->", error);

      // If an error occurs, redirect to the fail page
      router.push(`/resetPwFail`);
    }
  };
  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h2 className="mb-10 text-center text-gray-700">Reset your password</h2>
        <p className="mb-10 text-center text-gray-500">
          Enter a new password for your account.&apos;
        </p>
        <form
          className="mb-4 rounded bg-white px-8 pb-8 pt-6 shadow-md"
          onSubmit={handleInputEmail}
        >
          <div className="mb-6">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="password"
            >
              New Password
            </label>
            <input
              name="new_password"
              className="focus:shadow-outline mb-3 w-full appearance-none rounded border border-red-500 px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
              id="new_password"
              type="password"
              placeholder="******************"
            />
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="password"
            >
              Confirm Password
            </label>
            <input
              name="confirm_password"
              className="focus:shadow-outline mb-3 w-full appearance-none rounded border border-red-500 px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
              id="confirm_password"
              type="password"
              placeholder="******************"
            />
            <p className="text-xs italic text-red-500">
              Please type your new password.
            </p>
          </div>
          <div className="flex items-center justify-between">
            <button className="btn-violet w-full font-space_mono">
              Reset Password
            </button>
          </div>
        </form>

        <p className="text-center text-xs text-gray-500">
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}
