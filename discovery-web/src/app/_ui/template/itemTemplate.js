"use client";

import { useState } from "react";
import { useParams, useRouter } from "next/navigation";
import Link from "next/link";

import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHeart as faHeartSolid } from "@fortawesome/free-solid-svg-icons";
import { faHeart as faHeartRegular } from "@fortawesome/free-regular-svg-icons";

export default function ItemTemplete({
  imageSource,
  title,
  text,
  placeID,
  catagory,
  hasbookmark,
}) {
  const [emptyHeart, setEmptyHeart] = useState(!hasbookmark);
  const params = useParams();
  const router = useRouter();

  const handlePlaceBookmark = (e) => {
    e.preventDefault();
    if (!emptyHeart) {
      alert(`Unbookmark place: ${title}!`);
      fetchBookmark("http://localhost:8080/api/unBookmark");
    } else {
      alert(`Bookmark place: ${title}!`);
      fetchBookmark("http://localhost:8080/api/bookmark");
    }
    setEmptyHeart(!emptyHeart);
  };
  const fetchBookmark = async (requestLinkToServer) => {
    const data = {
      username: params.username,
      place_name: title,
      place_id: placeID,
      place_text: text,
      catagory: catagory,
    };

    const request = new Request(requestLinkToServer, {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request);
      if (response.ok) {
        const htmlContent = await response.json(); // Use json() to handle HTML response
      } else {
        console.error("Error fetching bookmark request:", response.statusText);
        throw error;
      }
    } catch (error) {
      console.error("Error fetching bookmark request:", error);
    }
  };

  return (
    <div className="mb-10 max-w-xl">
      <div className="mb-8 flex items-center justify-between gap-3">
        <Link
          href={`/${params.username}/location/${params.location}/${catagory}/${title}=${placeID}`}
        >
          <div className="flex items-center justify-between gap-3">
            <div className="">
              <Image
                src={imageSource}
                className="size-full rounded-lg object-cover"
                alt="Image of place"
                width={125}
                height={125}
              />
            </div>
            <span className="flex flex-col justify-between space-y-2">
              <h3>{title}</h3>
              <p>{text}</p>
            </span>
          </div>
        </Link>
        <span>
          <button onClick={handlePlaceBookmark}>
            <FontAwesomeIcon
              icon={emptyHeart ? faHeartRegular : faHeartSolid}
              size="2x"
            />
          </button>
        </span>
      </div>
      <hr />
    </div>
  );
}
