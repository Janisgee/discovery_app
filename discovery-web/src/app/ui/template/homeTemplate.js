"use client";

import { useParams, useRouter } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faFileLines,
  faHeart,
  faRightFromBracket,
} from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";

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
        throw new Error(`Failed to logout: ${response.statusText}`);
      }
      console.log("Logged out successfully");
      router.push("/");
    } catch (error) {
      console.error("Error logging out:", error);
    }
  };
  return (
    <div className="background-yellow font-inter">
      <div className="px-10 py-5">
        <div className="block-center">
          <div className="flex w-full max-w-sm justify-between">
            <Link href={`/${params.username}/trips`}>
              <FontAwesomeIcon icon={faFileLines} size="3x" />
            </Link>
            <div className="flex items-center justify-between">
              <Link href={`/${params.username}/bookmark`}>
                <FontAwesomeIcon icon={faHeart} size="3x" />
              </Link>
              <span className="ml-5">
                <button onClick={fetchLogoutRequest}>
                  <FontAwesomeIcon icon={faRightFromBracket} size="3x" />
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
