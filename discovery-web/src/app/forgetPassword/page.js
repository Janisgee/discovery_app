"use client";

import { Button } from "@/app/ui/buttons";
import Link from "next/link";

export default function forgetPassword() {
  const handleResetPassword = async (e) => {
    e.preventDefault();
  };
  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h2 className="mb-10 text-center text-gray-700">Reset your password</h2>
        <p className="mb-10 text-gray-500">
          Enter the email associated with your account and we&apos;ll send you
          password reset instructions.
        </p>
        <form
          className="mb-4 rounded bg-white px-8 pb-8 pt-6 shadow-md"
          onSubmit={handleResetPassword}
        >
          <div className="mb-6">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="password"
            >
              Your Email
            </label>
            <input
              name="login_password"
              className="focus:shadow-outline mb-3 w-full appearance-none rounded border border-red-500 px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
              id="email"
              type="email"
            />
            <p className="text-xs italic text-red-500">
              Please type your email.
            </p>
          </div>
          <div className="flex items-center justify-between">
            <button className="btn-violet w-full font-space_mono">
              Send Reset Instructions
            </button>
          </div>
        </form>
        <div className="mb-4 rounded-xl bg-slate-100 py-6 text-center">
          <p className="text-s pb-4 text-gray-500 ">
            Don&apos;t have an account?
          </p>
          <Link href={`/signup`}>
            <p className="inline-block align-baseline text-sm font-bold text-blue-500 hover:text-blue-800">
              Sign up now!
            </p>
          </Link>
        </div>
        <p className="text-center text-xs text-gray-500">
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}
