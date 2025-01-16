"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/app/ui/buttons";

export default function Signup() {
  const [passwordSuggest, setpasswordSuggest] = useState("");

  const router = useRouter();
  const handleSignupData = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const input_username = formData.get("username");
    const input_email = formData.get("signup_email");
    const input_password = formData.get("signup_password");

    if (!input_username) {
      alert("Please enter a username.");
      return;
    }
    if (!input_email) {
      alert("Please enter an email.");
      return;
    }
    if (!input_password) {
      alert("Please enter a password.");
      return;
    }
    fetchSignupData(input_username, input_email, input_password);
  };

  const fetchSignupData = async (username, email, password) => {
    const data = { username: username, email: email, password: password };
    console.log(data);
    const request = new Request("http://localhost:8080/api/signup", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request);
      if (!response.ok) {
        // Get error response from password validation
        const responseData = await response.json();
        console.log("Server Error Response for signup:", responseData);
        setpasswordSuggest(capitalizeFirstLetter(responseData.error));
        throw new Error(`Failed to fetch: ${response.statusText}`);
      }

      const responseData = await response.json();
      console.log("Server Response for signup:", responseData);

      router.push(`/${encodeURIComponent(responseData.username)}/home`);

      // router.push(`/dashboard/home`);
    } catch (error) {
      console.error("Error fetching user sign up:", error);
      // alert("Error fetching user home page. Please try again later.");
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
              className="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
              name="username"
              id="username"
              type="text"
              placeholder="Username"
            />
          </div>
          <div className="mb-4">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="email"
            >
              Email
            </label>
            <input
              className="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
              name="signup_email"
              id="email"
              type="text"
              placeholder="Email"
            />
          </div>
          <div className="mb-6">
            <label
              className="mb-2 block text-sm font-bold text-gray-700"
              htmlFor="password"
            >
              Password
            </label>
            <input
              className="focus:shadow-outline mb-3 w-full appearance-none rounded border border-red-500 px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none"
              name="signup_password"
              id="password"
              type="password"
              placeholder="******************"
            />
            <p className="text-xs italic text-red-500">
              {passwordSuggest == ""
                ? "Please choose a password."
                : passwordSuggest}
            </p>
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
