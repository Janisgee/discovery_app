"use client";
import { useParams } from "next/navigation";
import PlaceTemplate from "@/app/_ui/template/placeTemplate";

export default function ActivityPlace() {
  const { username, location, activity } = useParams();

  return (
    <PlaceTemplate
      username={username}
      location={location}
      place={activity}
      catagory="activity"
    />
  );
}
