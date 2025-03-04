import { Button } from "@/app/_ui/buttons";

export default function ForgotPasswordExpired() {
  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h2 className="mb-10 w-full text-center text-gray-700">
          Whoops, that&apos;s an expired link
        </h2>
        <p className="text-s mb-10 text-gray-500">
          For security reasons, password reset links expire after a little
          while. If you still need to reset your password, you can request a new
          reset email.
        </p>

        <div className="mb-10 text-center">
          <Button
            useFor="Request a new reset email"
            link="/forgot-password"
            color="btn-violet"
          />
        </div>
        <p className="text-center text-xs text-gray-500">
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}
