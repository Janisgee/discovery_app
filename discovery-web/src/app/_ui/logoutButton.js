"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faRightFromBracket } from "@fortawesome/free-solid-svg-icons";

export default function LogoutButton() {
  const router = useRouter();
  const fetchLogoutRequest = async () => {
    const request = new Request(
      `${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/logout`,
      {
        credentials: "include",
      },
    );

    try {
      const response = await fetch(request);
      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Logout failed: ${errorText}`);
      }

      localStorage.setItem("logout", Date.now().toString());
      router.push("/");
    } catch (error) {
      console.error("Error logging out:", error);
    }
  };

  const onStorageChange = (e) => {
    if (e.key == "logout") {
      router.push("/");
    }
  };

  useEffect(() => {
    window.addEventListener("storage", onStorageChange);
    return () => {
      window.removeEventListener("storage", onStorageChange);
    };
  });
  return (
    <>
      <button onClick={fetchLogoutRequest}>
        <FontAwesomeIcon icon={faRightFromBracket} size="2x" />
      </button>
    </>
  );
}
