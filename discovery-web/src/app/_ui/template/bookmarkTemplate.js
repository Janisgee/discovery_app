"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { TimeoutModule } from "@/app/_ui/timeoutModule";
import ItemTemplete from "@/app/_ui/template/itemTemplate";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";

import {
  faFileLines,
  faHeart,
  faHouse,
  faBinoculars,
  faBurger,
  faStore,
  faBaseballBatBall,
  faBed,
  faGasPump,
  faRightFromBracket,
  faCircleArrowLeft,
} from "@fortawesome/free-solid-svg-icons";

export default function BookmarkTemplate({
  username,
  country,
  city,
  placeList,
  children,
}) {
  const [catagoryClicked, setCatagoryClicked] = useState("attraction");
  console.log(username);
  console.log(city);
  console.log(placeList);

  const params = useParams();
  let itemList = [];
  if (placeList.length > 0) {
    // Filter the city with different catagory with click
    const filterItem = placeList.filter(
      (list) => list.Catagory == catagoryClicked,
    );

    filterItem.forEach((item, index) => {
      itemList.push(
        <Link
          href={`/${username}/location/${city}/${item.Catagory}/${item.PlaceName}=${item.PlaceID}`}
          key={index}
        >
          <ItemTemplete
            imageSource={item.PlaceDetail.image_url}
            title={item.PlaceName}
            text={item.PlaceText}
            placeID={item.PlaceID}
            catagory={item.Catagory}
            hasbookmark={true}
          ></ItemTemplete>
        </Link>,
      );
    });
  }

  const handleCatagoryClick = (e, catagory) => {
    e.preventDefault();
    setCatagoryClicked(catagory);
  };
  return (
    <div className="background-yellow font-inter">
      <TimeoutModule username={params.username} />
      <div className="p-5">
        <div>
          <div className="mb-5 flex items-center justify-between">
            <span>
              <span className="mr-5">
                <Link href={`/${username}/home`}>
                  <FontAwesomeIcon icon={faHouse} size="3x" />
                </Link>
              </span>
            </span>
            <span className="flex items-center">
              <Link href={`/${username}/bookmark`}>
                <FontAwesomeIcon icon={faHeart} size="3x" />
              </Link>
              <span className="ml-5">
                <Link href="/">
                  <FontAwesomeIcon icon={faRightFromBracket} size="3x" />
                </Link>
              </span>
            </span>
          </div>
          <div className="flex justify-center gap-5 pt-5">
            <Link href={`/${username}/bookmark/${country}`}>
              <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
            </Link>
            <h3 className="mb-5 text-center">Bookmark in {city}</h3>
          </div>
          <div className="text-center">
            <button
              className={`${catagoryClicked == "attraction" ? "btn-attraction-active" : "btn-white"}`}
              onClick={(e) => handleCatagoryClick(e, "attraction")}
            >
              <FontAwesomeIcon icon={faBinoculars} size="xl" />
            </button>
            <button
              className="btn-white"
              onClick={(e) => handleCatagoryClick(e, "restaurant")}
            >
              <FontAwesomeIcon icon={faBurger} size="xl" />
            </button>
            <button
              className="btn-white"
              onClick={(e) => handleCatagoryClick(e, "shopping")}
            >
              <FontAwesomeIcon icon={faStore} size="xl" />
            </button>
            <button
              className="btn-white"
              onClick={(e) => handleCatagoryClick(e, "activity")}
            >
              <FontAwesomeIcon icon={faBaseballBatBall} size="xl" />
            </button>
            <button
              className="btn-white"
              onClick={(e) => handleCatagoryClick(e, "hotel")}
            >
              <FontAwesomeIcon icon={faBed} size="xl" />
            </button>
            <button
              className="btn-white"
              onClick={(e) => handleCatagoryClick(e, "petrol_station")}
            >
              <FontAwesomeIcon icon={faGasPump} size="xl" />
            </button>
          </div>
        </div>
      </div>
      <div className="rounded-t-3xl bg-white ">
        <div className="px-10 pb-10">
          <h3 className="py-4 text-center text-gray-500">
            {catagoryClicked.toUpperCase()}
          </h3>
          <hr />
          {itemList ? (
            <div className="w-full pt-8">{itemList}</div>
          ) : (
            <p className="pt-8 text-center ">No Bookmark</p>
          )}
          {children}
        </div>
      </div>
    </div>
  );
}
