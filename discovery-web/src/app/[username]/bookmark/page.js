"use client";

import { useState, useEffect } from "react";
import { useParams } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft } from "@fortawesome/free-solid-svg-icons";
import AppTemplate from "@/app/ui/template/appTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";
import Link from "next/link";
import { fetchAllBookmark } from "@/app/ui/fetchAPI/fetchBookmark";

export default function Bookmark() {
  const [content, setContent] = useState(null);
  const params = useParams();

  const fetchData = async () => {
    try {
      const data = await fetchAllBookmark();
      console.log(data);
      setContent(data);
    } catch {
      console.error("Error fetching all booking data:", error);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  let itemList = [];
  if (content != null && content.BookmarkedPlace.length > 1) {
    // Filter the country be unique exclude duplicate
    const uniqueCountries = Object.values(
      content.BookmarkedPlace.reduce((acc, item) => {
        if (!acc[item.Country]) {
          acc[item.Country] = item;
        }
        return acc;
      }, {}),
    );
    console.log(uniqueCountries);
    // Sort country assending order
    const sortedCountry = uniqueCountries.sort((a, b) =>
      a.Country.localeCompare(b.Country),
    );
    sortedCountry.forEach((item, index) => {
      itemList.push(
        <Link href={`/${params.username}/bookmark/${item.Country}`} key={index}>
          <CardTemplete
            imageSource="/catagory_img/attraction.jpg"
            text={item.Country}
          />
        </Link>,
      );
    });
  }

  return (
    <div>
      <AppTemplate>
        <div className="flex items-center justify-center gap-3 pt-5">
          <Link href={`/${params.username}/home`}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </Link>
          <h2 className="text-center">Bookmark Country</h2>
        </div>
        <div className="w-full overflow-auto rounded-lg">
          <ul>{itemList}</ul>
        </div>
      </AppTemplate>
    </div>
  );
}
