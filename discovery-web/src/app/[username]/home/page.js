"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useParams } from "next/navigation";

import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons";

import { fetchAllBookmark } from "@/app/ui/fetchAPI/fetchBookmark";

import HomeTemplate from "@/app/ui/template/homeTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";
import Link from "next/link";

export default function Home() {
  const [bookmarkNum, setBookmarkNum] = useState(0);
  const [country1, setCountry1] = useState("");
  const [country2, setCountry2] = useState("");
  const [country3, setCountry3] = useState("");
  // const [searchCountry, setSearchCountry] = useState("");

  const params = useParams();
  const router = useRouter();

  const generateRandomCountry = () => {
    const randomCountry = require("random-country");
    const country1 = randomCountry({ full: true });
    if (country1 != "") {
      setCountry1(country1);
    }
    const country2 = randomCountry({ full: true });
    if (country2 != "") {
      setCountry2(country2);
    }
    const country3 = randomCountry({ full: true });
    if (country3 != "") {
      setCountry3(country3);
    }
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

  const fetchData = async () => {
    try {
      const data = await fetchAllBookmark();
      console.log(data);
      if (data.BookmarkedPlace != null) {
        setBookmarkNum(data.BookmarkedPlace.length);
      }

      // setBookmarkNum(len(data.BookmarkedPlace));
    } catch (error) {
      console.error("Error fetching all booking data:", error);
    }
  };

  useEffect(() => {
    fetchData();
    generateRandomCountry();
  }, []);

  return (
    <div>
      <HomeTemplate>
        <div className="block-center flex-col pb-8">
          <Image
            src="/user_img/default.jpg"
            className="mb-5 h-32 rounded-full"
            alt="default profile picture"
            width={128}
            height={128}
          />
          <h6 className="">{params.username}</h6>
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
          {/* <div className='w-full h-96 overflow-hidden inline-block rounded-lg'> */}
          <div className="h-96 w-full overflow-auto rounded-lg">
            <Link
              href={`/${params.username}/location/${encodeURIComponent(country1)}`}
            >
              <CardTemplete
                imageSource="/place_img/paris-france.jpg"
                text={country1}
              />
            </Link>
            <Link
              href={`/${params.username}/location/${encodeURIComponent(country2)}`}
            >
              <CardTemplete
                imageSource="/place_img/paris-france.jpg"
                text={country2}
              />
            </Link>
            <Link
              href={`/${params.username}/location/${encodeURIComponent(country3)}`}
            >
              <CardTemplete
                imageSource="/place_img/paris-france.jpg"
                text={country3}
              />
            </Link>
          </div>
        </div>
      </HomeTemplate>
    </div>
  );
}
