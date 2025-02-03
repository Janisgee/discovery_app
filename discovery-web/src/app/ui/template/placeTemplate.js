"use client";

import { useEffect, useState } from "react";

import AppTemplate from "@/app/ui/template/appTemplate";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft, faHeart } from "@fortawesome/free-solid-svg-icons";
import Image from "next/image";
import Link from "next/link";

export default function PlaceTemplate({ username, location, place, catagory }) {
  const [content, setContent] = useState([]);
  const decodeURIPlace = decodeURIComponent(place).toUpperCase();

  const fetchSearchPlaceDetails = async () => {
    const data = { place: place };

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

        setContent(htmlContent);
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
          <Link href={`/${username}/location/${location}/${catagory}`}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </Link>
          <h1 className="text-center text-xl">{decodeURIPlace}</h1>
          <FontAwesomeIcon icon={faHeart} size="2x" />
        </div>
        <Image
          src="/place_img/paris-france.jpg"
          className="mt-5 h-32 w-full rounded-lg object-cover"
          alt="Image of place"
          width={330}
          height={130}
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
      </AppTemplate>
    </div>
  );
}
