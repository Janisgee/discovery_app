"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
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
}) {
  const [emptyHeart, setEmptyHeart] = useState(true);
  const params = useParams();

  const handlePlaceBookmark = (e) => {
    e.preventDefault();
    setEmptyHeart(!emptyHeart);
    if (emptyHeart == false) {
      alert(`Unbookmark place: ${title}!`);
    } else {
      alert(`Bookmark place: ${title}!`);
      fetchBookmark();
    }
  };
  const fetchBookmark = async () => {
    const data = {
      username: params.username,
      place_name: title,
      place_id: placeID,
      place_text: text,
      catagory: catagory,
    };

    const request = new Request("http://localhost:8080/api/bookmark", {
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
        console.log(htmlContent);
      } else {
        console.error("Error fetching search country:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching search country:", error);
    }
  };

  return (
    <div className="mb-10 max-w-xl">
      <div className="mb-8 flex items-center justify-between gap-3">
        <Link
          href={`/${params.username}/location/${params.location}/${catagory}/${title}`}
        >
          <div className="flex items-center justify-between gap-3">
            <div className="size-28">
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
