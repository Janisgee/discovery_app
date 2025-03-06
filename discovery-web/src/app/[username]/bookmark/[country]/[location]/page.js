"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";
import BookmarkTemplate from "@/app/_ui/template/bookmarkTemplate";

export default function BookmarkLocation() {
  const [display, setDisplay] = useState(false);
  const [content, setContent] = useState([]);
  const params = useParams();

  const bookmarkLocation = params.location.toUpperCase().replaceAll("%20", " ");

  const spaceLocation = params.location.replaceAll("%20", " ");

  const fetchAllBookmarkByCity = async () => {
    const data = {
      city: spaceLocation,
    };
    const request = new Request(
      `${process.env.NEXT_PUBLIC_API_SERVER_BASE}/api/getAllBookmarkByCity`,
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

      if (response.ok) {
        const htmlContent = await response.json(); // Use json() to handle HTML response

        if (htmlContent != null) {
          setContent(htmlContent);
        }
      } else {
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
