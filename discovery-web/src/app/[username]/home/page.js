"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useParams } from "next/navigation";
import { CldImage } from "next-cloudinary";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons";

import { fetchProfileImage } from "@/app/ui/fetchAPI/fetchProfileImage";
import { fetchAllBookmark } from "@/app/ui/fetchAPI/fetchBookmark";
import { fetchPlaceImage } from "@/app/ui/fetchAPI/fetchPlaceImage";

import HomeTemplate from "@/app/ui/template/homeTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";
import Link from "next/link";

export default function Home() {
  const [bookmarkNum, setBookmarkNum] = useState(0);
  const [userPublicID, setUserPublicID] = useState("");
  const [itemList, setItemList] = useState([]);

  const params = useParams();
  const router = useRouter();
  const countryGroupNumber = 3;
  const defaultImage =
    "https://res.cloudinary.com/dopxvbeju/image/upload/v1740039540/kphottt1vhiuyahnzy8y.jpg";

  const fetchCountryImage = async () => {
    let itemList = [];
    for (let i = 0; i < countryGroupNumber; i++) {
      const randomCountry = require("random-country");
      const country = randomCountry({ full: true });
      try {
        const imageURL = await fetchPlaceImage(country, "country");
        itemList.push(
          <Link
            href={`/${params.username}/location/${encodeURIComponent(country)}`}
            key={i}
          >
            <CardTemplete imageSource={imageURL} text={country} />
          </Link>,
        );
      } catch (error) {
        console.error("Error fetching image for country:", error);
      }
    }
    setItemList(itemList);
    console.log(itemList);
  };

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const searchData = formData.get("search");
    if (!searchData) {
      alert("Please enter a location");
      return;
    }
    router.push(
      `/${params.username}/location/${encodeURIComponent(searchData)}`,
    );
  };

  const fetchData = async (router) => {
    try {
      const data = await fetchAllBookmark(router);
      console.log(data);
      if (data.BookmarkedPlace != null) {
        setBookmarkNum(data.BookmarkedPlace.length);
      }
      const imageData = await fetchProfileImage(router);
      console.log(imageData);
      if (imageData != null) {
        setUserPublicID(imageData.publicID);
      }
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    fetchData(router);
    fetchCountryImage();
    console.log(itemList);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  console.log("upid", userPublicID);

  return (
    <div>
      <HomeTemplate>
        <div className="block-center flex-col pb-8">
          {!userPublicID ? (
            <CldImage
              src={defaultImage}
              className="h-32 rounded-full"
              alt={params.username}
              width="128"
              height="128"
              crop="fill"
              gravity="face"
            />
          ) : (
            <CldImage
              src={userPublicID}
              className="h-32 rounded-full"
              alt={params.username}
              width="128"
              height="128"
              crop="fill"
              gravity="face"
            />
          )}
          <h6>{params.username}</h6>
          <p className="text-color-dark_grey">{`${bookmarkNum} bookmark places`}</p>
        </div>
        <div className="block-center">
          <form
            onSubmit={handleSearchSubmit}
            className="grid grid-cols-8 gap-x-2"
          >
            <span className="col-span-2 justify-self-end p-2">
              <FontAwesomeIcon icon={faMagnifyingGlass} size="2x" />
            </span>
            <span className="inline-center col-span-4 p-2">
              <label
                htmlFor="default-search"
                className="sr-only mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >
                Search
              </label>
              <input
                name="search"
                type="search"
                id="default-search"
                className="w-full rounded-full border border-gray-300 bg-gray-50 p-2 ps-5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder:text-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
                placeholder="Type location to search"
                required
              />
            </span>
          </form>
        </div>
        <div className="block-center my-5 flex-col ">
          <h3>Popular Destination</h3>

          <div className="size-full overflow-auto rounded-lg">
            <ul>{itemList != [] && itemList}</ul>
          </div>
        </div>
      </HomeTemplate>
    </div>
  );
}
