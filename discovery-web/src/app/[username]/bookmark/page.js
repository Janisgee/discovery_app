"use client";

import { useState, useEffect } from "react";

import { useRouter, useParams } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft } from "@fortawesome/free-solid-svg-icons";
import AppTemplate from "@/app/ui/template/appTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";
import Link from "next/link";
import { fetchAllBookmark } from "@/app/ui/fetchAPI/fetchBookmark";

export default function Bookmark() {
  const [content, setContent] = useState(null);
  const [noBookmark, setNoBookmark] = useState(true);
  const params = useParams();
  const router = useRouter();

  const fetchData = async (router) => {
    const bookmarks = await fetchAllBookmark(router);

    if (bookmarks) {
      console.log(bookmarks);

      // Handle the bookmarks data here (e.g., display it)
      if (bookmarks.BookmarkedPlace.length > 1) {
        setNoBookmark(false);
        setContent(bookmarks);
        console.log("Fetched bookmarks:", bookmarks);
      } else {
        setNoBookmark(true);
        console.log("No bookmarks available.");
      }
    } else {
      console.log("Error occurred.");
      // Handle error case here (e.g., show error UI)
    }
  };

  useEffect(() => {
    fetchData(router);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  let itemList = [];

  if (content && !noBookmark && content.BookmarkedPlace.length > 1) {
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
          {noBookmark ? (
            <p className="pt-10 text-center text-lg">
              There is no bookmark place from user.
            </p>
          ) : (
            <ul>{itemList}</ul>
          )}
        </div>
      </AppTemplate>
    </div>
  );
}
