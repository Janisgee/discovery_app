"use client";

import { Button } from "@/app/_ui/buttons";
import { useSetNewPassword } from "@/app/forgot-password/_hooks/password-reset-service";
import { LoadingSpinner } from "@/app/_ui/loading-spinner";

export default function SetNewPassword() {
  const [doSetNewPassword, isPending, error] = useSetNewPassword();

  const handleInputEmail = async (e) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    const newPw = formData.get("new_password");
    const confirmPw = formData.get("confirm_password");

    doSetNewPassword(newPw, confirmPw);
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
          noValidate
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
            <p className="min-h-[16px] text-xs italic text-red-500">{error}</p>
          </div>
          <div className="flex items-center justify-between">
            <button className="btn-violet w-full font-space_mono">
              {isPending ? <LoadingSpinner /> : "Reset Password"}
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
