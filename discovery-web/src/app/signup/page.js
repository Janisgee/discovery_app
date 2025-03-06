"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/app/_ui/buttons";

export default function Signup() {
  const [suggest, setSuggest] = useState("");
  const [alertUsername, setalertUsername] = useState(false);
  const [alertEmail, setalertEmail] = useState(false);
  const [alertPassword, setalertPassword] = useState(false);

  const router = useRouter();
  const handleSignupData = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const input_username = formData.get("username");
    const input_email = formData.get("signup_email");
    const input_password = formData.get("signup_password");

    if (!input_username || !input_email || !input_password) {
      if (!input_username) setalertUsername(true);
      if (!input_email) setalertEmail(true);
      if (!input_password) setalertPassword(true);
      setSuggest("");
      return;
    }
    fetchSignupData(input_username, input_email, input_password);
  };

  const handleInputOnChange = (e) => {
    e.preventDefault();
    if (e.currentTarget.id == "username" && e.currentTarget.value != "") {
      setalertUsername(false);
    }
    if (e.currentTarget.id == "email" && e.currentTarget.value != "") {
      setalertEmail(false);
    }
    if (e.currentTarget.id == "password" && e.currentTarget.value != "") {
      setalertPassword(false);
    }
  };

  const fetchSignupData = async (username, email, password) => {
    const data = { username: username, email: email, password: password };

    const request = new Request(
      `${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/signup`,
      {
        method: "POST", // HTTP method
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      },
    );

    try {
      const response = await fetch(request);
      if (!response.ok) {
        // Get error response from password validation
        const errorResponse = await response.json();
        console.log(errorResponse);
        setSuggest(capitalizeFirstLetter(errorResponse.error));
        throw new Error(`Failed to fetch: ${response.statusText}`);
      }

      router.push(`/login`);
    } catch (error) {
      console.error("Error fetching user sign up:", error);
    }
  };

  const capitalizeFirstLetter = (sentence) => {
    return String(sentence).charAt(0).toUpperCase() + String(sentence).slice(1);
  };

  return (
    <div className="block-center p-20">
      <div className="w-full max-w-sm">
        <div className="mb-10 text-right">
          <Button useFor="Back" link="/" color="btn-grey" />
        </div>
        <h2 className="mb-10 text-center text-gray-500">Sign Up</h2>
        <form
          className="mb-4 rounded bg-white px-8 pb-8 pt-6 shadow-md"
          onSubmit={handleSignupData}
        >
          <div className="mb-4">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="username"
            >
              Username
            </label>
            <input
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertUsername ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none`}
              name="username"
              id="username"
              type="text"
              placeholder="Username"
              onChange={handleInputOnChange}
            />{" "}
            {alertUsername ? (
              <span className="text-xs italic text-red-500">
                (Please fill in your username.)
              </span>
            ) : (
              <span> </span>
            )}
          </div>
          <div className="mb-4">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="email"
            >
              Email
            </label>
            <input
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertEmail ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none`}
              name="signup_email"
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
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertPassword ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none`}
              name="signup_password"
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
            {suggest && (
              <p className="mt-5 text-xs italic text-red-500">{suggest}</p>
            )}{" "}
          </div>
          <div className="flex items-center">
            <button className="btn-violet">Sign Up</button>
          </div>
        </form>
        <p className="text-center text-xs text-gray-500">
          &copy;2024 Discovery App. All rights reserved.
        </p>
      </div>
    </div>
  );
}
