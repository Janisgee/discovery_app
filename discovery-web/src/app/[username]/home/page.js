"use client";
import { useEffect, useState } from "react";

import { useRouter } from "next/navigation";
import { useParams } from "next/navigation";
import { CldImage } from "next-cloudinary";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons";

import { fetchProfileImage } from "@/app/_ui/fetchAPI/fetchProfileImage";
import { fetchAllBookmark } from "@/app/_ui/fetchAPI/fetchBookmark";
import { fetchPlaceImage } from "@/app/_ui/fetchAPI/fetchPlaceImage";
import { fetchInputControl } from "@/app/_ui/fetchAPI/fetchInputControl";

import { LoadingSpinner } from "@/app/_ui/loading-spinner";
import HomeTemplate from "@/app/_ui/template/homeTemplate";
import CardTemplete from "@/app/_ui/template/cardTemplate";
import Link from "next/link";

export default function Home() {
  const [isPending, setIsPending] = useState(false);
  const [bookmarkNum, setBookmarkNum] = useState(0);
  const [userPublicID, setUserPublicID] = useState("");
  const [itemList, setItemList] = useState([]);
  const [searchWord, setSearchWord] = useState("");
  const [optionResult, setOptionResult] = useState([]);
  const [resultArray, setResultArray] = useState([]);

  const params = useParams();
  const router = useRouter();
  const countryGroupNumber = 3;
  const defaultImage =
    "https://res.cloudinary.com/dopxvbeju/image/upload/v1740039540/kphottt1vhiuyahnzy8y.jpg";

  const fetchCountryImage = async () => {
    let itemList = [];
    setIsPending(true);
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
    setIsPending(false);
  };

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    if (resultArray.length > 0) {
      // const formData = new FormData(e.currentTarget);
      // const searchData =/ formData.get("search");
      if (!resultArray) {
        alert("Please enter a location");
        return;
      }
      router.push(
        `/${params.username}/location/${encodeURIComponent(resultArray[0].name)}`,
      );
    }
  };

  const handleInputChange = async (e) => {
    e.preventDefault();
    setSearchWord(e.target.value);
    try {
      let optionList = [];
      const result = await fetchInputControl(e.target.value);

      if (result) {
        for (let i = 0; i < result.length; i++) {
          optionList.push(
            <option value={result[i].name}>{result[i].name}</option>,
          );
        }
        setOptionResult(optionList);
        setResultArray(result);
      }
    } catch (error) {
      console.error("Error fetching word data:", error);
    }
  };

  const fetchData = async (router) => {
    try {
      const data = await fetchAllBookmark(router);

      if (data.BookmarkedPlace != null) {
        setBookmarkNum(data.BookmarkedPlace.length);
      }
      const imageData = await fetchProfileImage();

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

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div>
      <HomeTemplate>
        {/* {params.username != undefined && (
          <TimeoutModule username={params.username} />
        )} */}
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
                list="searchList"
                className="w-full rounded-full border border-gray-300 bg-gray-50 p-2 ps-5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder:text-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
                placeholder="Type location to search"
                value={searchWord}
                required
                onChange={handleInputChange}
              />
              <datalist id="searchList">{optionResult}</datalist>
            </span>
          </form>
        </div>
        <div className="block-center my-5 flex-col ">
          <h3>Popular Destination</h3>
          {isPending ? (
            <div className="mt-40 flex items-center justify-center">
              <LoadingSpinner size={48} color="text-violet-600" />
            </div>
          ) : (
            <div className="size-full overflow-auto rounded-lg">
              <ul>{itemList.length > 0 && itemList}</ul>
            </div>
          )}
        </div>
      </HomeTemplate>
    </div>
  );
}
