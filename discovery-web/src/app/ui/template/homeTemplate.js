"use client";

import { useParams, useRouter } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faGear,
  faHeart,
  faRightFromBracket,
} from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";
import { Button } from "../buttons";

export default function HomeTemplate({ children }) {
  const params = useParams();
  const router = useRouter();

  const fetchLogoutRequest = async () => {
    const request = new Request("http://localhost:8080/api/logout", {
      method: "POST", // HTTP method
      credentials: "include",
    });

    try {
      const response = await fetch(request);
      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Logout failed: ${errorText}`);
      }
      console.log("Logged out successfully");
      router.push("/");
    } catch (error) {
      console.error("Error logging out:", error);
      // Show a user-friendly message (optional)
      alert("An error occurred while logging out. Please try again.");
    }
  };
  return (
    <div className="background-yellow font-inter">
      <div className="px-10 py-5">
        <div className="block-center">
          <div className="flex w-full max-w-sm justify-between">
            <FontAwesomeIcon icon={faGear} size="2x" />
            <div className="flex items-center justify-between">
              <Link href={`/${params.username}/bookmark`}>
                <FontAwesomeIcon icon={faHeart} size="2x" />
              </Link>
              <span className="ml-5">
                <button onClick={fetchLogoutRequest}>
                  <FontAwesomeIcon icon={faRightFromBracket} size="2x" />
                </button>
              </span>
            </div>
          </div>
        </div>
      </div>
      <div className="rounded-t-3xl bg-white ">
        <div className="px-10 pb-10">{children}</div>
      </div>
    </div>
  );
}
