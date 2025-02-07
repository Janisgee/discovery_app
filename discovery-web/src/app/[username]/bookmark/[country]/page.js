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
      } else {
        if (response.status == 401) {
          alert(
            `Please login again as 10 mins session expired without taking action.`,
          );
          router.push(`/login`);
        }
        console.error("Error fetching bookmark place:", response.statusText);
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
      content.BookmarkedPlace.filter((item) => item.Country == params.country),
    );
    console.log(filterByCountry);
    // Sort city assending order
    const sortedCity = filterByCountry.sort((a, b) =>
      a.City.localeCompare(b.City),
    );

    sortedCity.forEach((item, index) => {
      itemList.push(
        <Link
          href={`/${params.username}/bookmark/${item.Country}/${item.City}`}
          key={index}
        >
          <CardTemplete
            imageSource="/catagory_img/attraction.jpg"
            text={item.City}
          />
        </Link>,
      );
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
