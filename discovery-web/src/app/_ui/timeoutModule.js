"use client";

import { useEffect, useState, useRef } from "react";
import Link from "next/link";

export function TimeoutModule({ username }) {
  const [contentItem, setContentItem] = useState(null);
  const warningTimeoutID = useRef(null);

  const events = ["click", "mousemove", "mousedown", "keydown"];

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
    } catch (error) {
      console.error("Error logging out:", error);
    }
  };

  useEffect(() => {
    // Check if the username is in the current path
    if (window.location.pathname.includes(`${username}`)) {
      // Set a timeout to show the session warning

      warningTimeoutID.current = setTimeout(
        callTimeoutFunc,
        process.env.NEXT_PUBLIC_AUTOLOGOUT_TIME,
      ); // 5 seconds for demo purposes

      // Add event listeners for user activity
      events.forEach((event) => {
        window.addEventListener(event, eventHandler);
      });
    }

    // Cleanup event listeners when the component unmounts
    return () => {
      events.forEach((event) => {
        window.removeEventListener(event, eventHandler);
      });
      if (warningTimeoutID.current) clearTimeout(warningTimeoutID.current);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [username]);

  // Function to show the timeout warning
  function callTimeoutFunc() {
    setContentItem(
      <div className="fixed inset-0 z-50 flex items-center justify-center bg-gray-900/50">
        <div
          id="alert-additional-content-2"
          className="mb-4 rounded-lg border border-red-300 bg-red-50 p-4 text-red-800 dark:border-red-800 dark:bg-gray-800 dark:text-red-400"
          role="alert"
        >
          <div className="flex items-center">
            <svg
              className="me-2 size-4 shrink-0"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z" />
            </svg>
            <span className="sr-only">Info</span>
            <h3 className="text-lg font-medium">Session expired</h3>
          </div>
          <div className="mb-4 mt-2 text-sm">
            You have been inactive for some time and now you are being logged
            out. Please log in again.
          </div>
          <div className="flex">
            <Link href={`/login`}>
              <button
                type="button"
                className="rounded-lg border border-red-800 bg-transparent px-3 py-1.5 text-center text-xs font-medium text-red-800 hover:bg-red-900 hover:text-white focus:outline-none focus:ring-4 focus:ring-red-300 dark:border-red-600 dark:text-red-500 dark:hover:bg-red-600 dark:hover:text-white dark:focus:ring-red-800"
                data-dismiss-target="#alert-additional-content-2"
                aria-label="Close"
              >
                Dismiss
              </button>
            </Link>
          </div>
        </div>
      </div>,
    );
    fetchLogoutRequest();
  }

  // Reset the warning timeout and logout timeout on any user interaction
  function eventHandler() {
    clearTimeout(warningTimeoutID.current);
    warningTimeoutID.current = setTimeout(
      callTimeoutFunc,
      process.env.NEXT_PUBLIC_AUTOLOGOUT_TIME,
    ); // reset the timeout
  }

  return <>{contentItem}</>;
}
