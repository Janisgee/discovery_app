"use client";

import { Button } from "@/app/ui/buttons";

export default function EmailSent() {
  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h1 className="mb-10 w-full text-center text-gray-700">
          Please Check Your Email for Account Recovery
        </h1>
        <h3 className="mb-10 w-full text-gray-500">
          If no email is received in a few minutes, please check email address
          was entered correctly.
        </h3>
        <h3 className="mb-10 w-full text-gray-500">
          Recovery emails are valid for 45 minutes.
        </h3>

        <p className="text-center text-xs text-gray-500">
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}
