"use client";
import { useParams } from "next/navigation";
import PlaceTemplate from "@/app/_ui/template/placeTemplate";

export default function ShoppingPlace() {
  const { username, location, shopping } = useParams();

  return (
    <PlaceTemplate
      username={username}
      location={location}
      place={shopping}
      catagory="shopping"
    />
  );
}
