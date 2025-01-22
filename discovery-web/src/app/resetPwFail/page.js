"use client";

import { Button } from "@/app/ui/buttons";

export default function ResetPwFail() {
  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h2 className="mb-10 text-center text-gray-700">
          Sorry, We are unable to reset your password!
        </h2>

        <div className="mb-4 rounded-xl bg-slate-100 py-6  text-center">
          <p className="text-s pb-4 text-gray-500 ">
            Please contact our techical support team:
            <span className="text-blue-500">
              support_discoveryApp@gmail.com
            </span>{" "}
            to request for further assistance.
          </p>
        </div>

        <p className="text-center text-xs text-gray-500">
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}
