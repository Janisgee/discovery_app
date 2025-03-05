"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/app/_ui/buttons";
import Link from "next/link";

export default function Login() {
  const [alertEmail, setalertEmail] = useState(false);
  const [alertPassword, setalertPassword] = useState(false);
  const [alertLoginError, setalertLoginError] = useState(false);
  const router = useRouter();
  const handleLoginData = (e) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    const input_email = formData.get("login_email");
    const input_password = formData.get("login_password");

    if (!input_email || !input_password) {
      if (!input_email) setalertEmail(true);
      if (!input_password) setalertPassword(true);
      setalertLoginError(false);
      return;
    }
    fetchLoginData(input_email, input_password);
  };

  const handleInputOnChange = (e) => {
    e.preventDefault();
    console.log(e.currentTarget.id);
    if (e.currentTarget.id == "email" && e.currentTarget.value != "") {
      setalertEmail(false);
    }
    if (e.currentTarget.id == "password" && e.currentTarget.value != "") {
      setalertPassword(false);
    }
  };

  const fetchLoginData = async (loginEmail, loginPassword) => {
    const data = { email: loginEmail, password: loginPassword };

    const request = new Request("http://localhost:8080/api/login", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    try {
      await new Promise((resolve) => setTimeout(resolve, 2000));
      const response = await fetch(request);
      if (!response.ok) {
        throw new Error(`Failed to fetch: ${response.statusText}`);
      }

      const responseData = await response.json();

      if (responseData.username != "") {
        router.push(`/${encodeURIComponent(responseData.username)}/home`);
      }
    } catch (error) {
      setalertLoginError(true);
    }
  };

  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h2 className="mb-10 text-center text-gray-500">Login</h2>
        <form
          className="mb-4 rounded bg-white px-8 pb-8 pt-6 shadow-md"
          onSubmit={handleLoginData}
        >
          <div className="mb-4">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="email"
            >
              Email{" "}
            </label>
            <input
              name="login_email"
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertEmail ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none`}
              id="email"
              type="text"
              placeholder="Email"
              onChange={handleInputOnChange}
            />{" "}
            {alertEmail ? (
              <span className="text-xs italic text-red-500">
                (Please fill in your email address.)
              </span>
            ) : (
              <span> </span>
            )}
          </div>
          <div className="mb-6">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="password"
            >
              Password
            </label>
            <input
              name="login_password"
              className={`focus:shadow-outline mb-3 w-full appearance-none rounded border ${alertPassword ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none`}
              id="password"
              type="password"
              placeholder="******************"
              onChange={handleInputOnChange}
            />
            {alertPassword ? (
              <span className="text-xs italic text-red-500">
                (Please fill in your password.)
              </span>
            ) : (
              <span> </span>
            )}
            {alertLoginError ? (
              <span className="text-xs italic text-red-500">
                We couldn’t log you in. Please check your email and password and
                try again.
              </span>
            ) : (
              <span> </span>
            )}
          </div>
          <div className="flex items-center justify-between">
            <button className="btn-violet font-space_mono">Login</button>
            <Link
              href="/forgot-password"
              className="inline-block align-baseline text-sm font-bold text-blue-500 hover:text-blue-800"
            >
              Forgot Password?
            </Link>
          </div>
        </form>
        <div className="mb-4 rounded-xl bg-slate-100 py-6  text-center">
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
