"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faRightFromBracket } from "@fortawesome/free-solid-svg-icons";

export default function LogoutButton() {
  const router = useRouter();
  const fetchLogoutRequest = async () => {
    const request = new Request("http://localhost:8080/api/logout", {
      credentials: "include",
    });

    try {
      const response = await fetch(request);
      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Logout failed: ${errorText}`);
      }
      console.log("Logged out successfully");

      localStorage.setItem("logout", Date.now().toString());
      router.push("/");
    } catch (error) {
      console.error("Error logging out:", error);
    }
  };

  const onStorageChange = (e) => {
    console.log(e);
    console.log(e.key);
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
