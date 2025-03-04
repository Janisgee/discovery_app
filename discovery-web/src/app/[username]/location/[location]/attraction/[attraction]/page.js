"use client";
import { useParams } from "next/navigation";
import PlaceTemplate from "@/app/_ui/template/placeTemplate";

export default function AttractionPlace() {
  const { username, location, attraction } = useParams();

  return (
    <PlaceTemplate
      username={username}
      location={location}
      place={attraction}
      catagory="attraction"
    />
  );
}
