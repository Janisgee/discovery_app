"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft } from "@fortawesome/free-solid-svg-icons";
import AppTemplate from "@/app/ui/template/appTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";
import Link from "next/link";

export default function BookmarkCountry() {
  const [content, setContent] = useState(null);
  const params = useParams();
  const decodeCountryName = decodeURIComponent(params.country);
  const fetchAllBookmark = async () => {
    const request = new Request("http://localhost:8080/api/getAllBookmark", {
      method: "GET", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
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
      console.error("Error fetching bookmark place:", error);
    }
  };
  //
  useEffect(() => {
    fetchAllBookmark();
  }, []);

  let itemList = [];
  if (content != null && content.BookmarkedPlace.length > 1) {
    // Filter the city be unique exclude duplicate
    const filterByCountry = Object.values(
      content.BookmarkedPlace.filter(
        (item) => item.Country == decodeCountryName,
      ),
    );
    // Filter the city be unique exclude duplicate
    const uniqueCity = Object.values(
      filterByCountry.reduce((acc, item) => {
        if (!acc[item.City]) {
          acc[item.City] = item;
        }
        return acc;
      }, {}),
    );
    console.log(uniqueCity);
    // Sort city assending order
    const sortedCity = uniqueCity.sort((a, b) => a.City.localeCompare(b.City));

    sortedCity.forEach((item, index) => {
      if (item.City) {
        itemList.push(
          <Link
            href={`/${params.username}/bookmark/${item.Country}/${item.City}`}
            key={index}
          >
            <CardTemplete
              imageSource=""
              text={item.City}
              searchFor="city"
              country={decodeCountryName}
            />
          </Link>,
        );
      } else {
        itemList.push(
          <Link
            href={`/${params.username}/bookmark/${item.Country}/${item.Country}`}
            key={index}
          >
            <CardTemplete
              imageSource=""
              text={item.City}
              searchFor="city"
              country={decodeCountryName}
            />
          </Link>,
        );
      }
    });
  }

  return (
    <div>
      <AppTemplate>
        <div className="flex items-center justify-center gap-3 pt-5">
          <Link href={`/${params.username}/bookmark`}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </Link>
          <h2 className="text-center">Bookmark City</h2>
        </div>
        <div className="w-full overflow-auto rounded-lg">
          <ul>{itemList}</ul>
        </div>
      </AppTemplate>
    </div>
  );
}
