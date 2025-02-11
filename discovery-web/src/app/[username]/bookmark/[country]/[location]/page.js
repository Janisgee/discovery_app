"use client";

import { useState, useEffect } from "react";
// import { useRouter } from "next/compat/router";
import { useParams } from "next/navigation";
import BookmarkTemplate from "@/app/ui/template/bookmarkTemplate";

// import Link from "next/link";

export default function BookmarkLocation() {
  const [display, setDisplay] = useState(false);
  const [content, setContent] = useState([]);
  const params = useParams();
  console.log(params);

  const bookmarkLocation = params.location.toUpperCase().replaceAll("%20", " ");
  // console.log(content);
  const spaceLocation = params.location.replaceAll("%20", " ");

  const fetchAllBookmarkByCity = async () => {
    const data = {
      city: spaceLocation,
    };
    const request = new Request(
      "http://localhost:8080/api/getAllBookmarkByCity",
      {
        method: "POST", // HTTP method
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(data),
      },
    );

    try {
      const response = await fetch(request);
      console.log("Response status:", response.status);
      if (response.ok) {
        const htmlContent = await response.json(); // Use json() to handle HTML response
        console.log("Received content:", htmlContent);
        if (htmlContent != null) {
          setContent(htmlContent);
        }
      } else {
        if (response.status == 401) {
          alert(
            `Please login again as 10 mins session expired without taking action.`,
          );
          router.push(`/login`);
        }
        console.error(
          `Error fetching bookmark by city ${bookmarkLocation}:`,
          response.statusText,
        );
      }
    } catch (error) {
      console.error(
        `Error fetching bookmark by city ${bookmarkLocation}:`,
        error,
      );
    }
  };
  //
  useEffect(() => {
    fetchAllBookmarkByCity();
    setDisplay(true);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div>
      {display && (
        <BookmarkTemplate
          username={params.username}
          country={params.country}
          city={bookmarkLocation}
          placeList={content.PlaceList != undefined && content.PlaceList}
        ></BookmarkTemplate>
      )}
    </div>
  );
}
