"use client";
import { useParams } from "next/navigation";
import PlaceTemplate from "@/app/ui/template/placeTemplate";

export default function PetrolStationPlace() {
  const { username, location, petrol_station } = useParams();

  return (
    <PlaceTemplate
      username={username}
      location={location}
      place={petrol_station}
      catagory="petrol_station"
    />
  );
}
