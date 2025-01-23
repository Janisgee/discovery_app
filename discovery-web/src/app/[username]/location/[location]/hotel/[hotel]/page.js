"use client";
import { useParams } from "next/navigation";
import PlaceTemplate from "@/app/ui/template/placeTemplate";

export default function HotelPlace() {
  const { username, location, hotel } = useParams();

  return (
    <PlaceTemplate
      username={username}
      location={location}
      place={hotel}
      catagory="hotel"
    />
  );
}
