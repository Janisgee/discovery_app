"use client";
import { useParams } from "next/navigation";
import PlaceTemplate from "@/app/ui/template/placeTemplate";

export default function RestaurantPlace() {
  const { username, location, restaurant } = useParams();

  return (
    <PlaceTemplate
      username={username}
      location={location}
      place={restaurant}
      catagory="restaurant"
    />
  );
}
