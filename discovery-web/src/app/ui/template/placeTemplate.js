"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

import AppTemplate from "@/app/ui/template/appTemplate";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft } from "@fortawesome/free-solid-svg-icons";
import { faHeart as faHeartSolid } from "@fortawesome/free-solid-svg-icons";
import { faHeart as faHeartRegular } from "@fortawesome/free-regular-svg-icons";
import Image from "next/image";
import Link from "next/link";

export default function PlaceTemplate({ username, location, place, catagory }) {
  const [content, setContent] = useState([]);
  const [hasBookmark, setHasBookmark] = useState(false);

  // Split location_name & placeID
  const placeArray = place.split("%3D");
  const placeName = placeArray[0];
  const placeID = placeArray[1];

  const router = useRouter();
  const decodeURIPlace = decodeURIComponent(placeName);
  const toUpperCasePlace = decodeURIPlace.toUpperCase();

  const backwardClick = (e) => {
    e.preventDefault();
    history.back();
  };

  const handlePlaceBookmark = (e) => {
    e.preventDefault();
    if (hasBookmark) {
      alert(`Unbookmark place: ${decodeURIPlace}!`);
      fetchBookmarkActionInPlaceDetail("http://localhost:8080/api/unBookmark");
    } else {
      alert(`Bookmark place: ${decodeURIPlace}!`);
      fetchBookmarkActionInPlaceDetail(
        "http://localhost:8080/api/bookmarkByPlaceName",
      );
    }
    setHasBookmark(!hasBookmark);
  };

  const fetchBookmarkActionInPlaceDetail = async (requestLinkToServer) => {
    const data = {
      username: username,
      place_name: decodeURIPlace,
      place_id: placeID,
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
        console.log(htmlContent);
      } else {
        console.error("Error fetching bookmark request:", response.statusText);
        throw error;
      }
    } catch (error) {
      console.error("Error fetching bookmark request:", error);
    }
  };

  const fetchSearchPlaceDetails = async () => {
    const data = { place: placeName, catagory: catagory };

    const request = new Request("http://localhost:8080/searchPlace", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request);
      console.log("Response status:", response.status);
      if (response.ok) {
        const htmlContent = await response.json(); // Use json() to handle HTML response
        console.log("Received content:", htmlContent);

        setContent(htmlContent.PlaceInfo);

        if (htmlContent.HasBookmark) {
          setHasBookmark(true);
        } else {
          setHasBookmark(false);
        }
      } else {
        console.error("Error fetching search place:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching search place:", error);
    }
  };
  //
  useEffect(() => {
    fetchSearchPlaceDetails();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div>
      <AppTemplate>
        <div className="flex items-center justify-between gap-3 pt-5">
          <button onClick={backwardClick}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </button>
          {/* <Link href={`/${username}/location/${location}/${catagory}`}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </Link> */}
          <h1 className="text-center text-xl">{toUpperCasePlace}</h1>

          <button onClick={handlePlaceBookmark}>
            <FontAwesomeIcon
              icon={hasBookmark ? faHeartSolid : faHeartRegular}
              size="2x"
            />
          </button>
        </div>
        {content.length != 0 && (
          <div>
            <Image
              src={content.image_url}
              className="mt-5 h-80 w-full rounded-lg object-cover"
              alt={`Image of ${decodeURIPlace}`}
              width={330}
              height={330}
            />
            <div className="mt-5">
              <p>{content.description}</p>
              <div className="mt-5">
                <h6>Location:</h6>
                <ul>
                  <li>{content.location}</li>
                  <li>Opening Hours: {content.opening_hours}</li>
                </ul>
              </div>{" "}
              <div className="mt-5">
                <h6>History:</h6>
                <ul>
                  <li>{content.history}</li>
                </ul>
              </div>
              <div className="mt-5">
                <h6>Key Features:</h6>
                <ul>
                  <li>{content.key_features}</li>
                </ul>
              </div>
              <div className="mt-5">
                <p>{content.conclusion}</p>
              </div>
            </div>
          </div>
        )}
      </AppTemplate>
    </div>
  );
}
