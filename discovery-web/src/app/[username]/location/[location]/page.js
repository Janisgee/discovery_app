"use client";

import { useParams } from "next/navigation";
import AppTemplate from "@/app/_ui/template/appTemplate";

import CardTemplete from "@/app/_ui/template/cardTemplate";

import Link from "next/link";

export default function LocationPlace() {
  const params = useParams();
  const location = decodeURIComponent(params.location).toUpperCase();

  return (
    <div>
      <AppTemplate>
        <h1 className="text-center">{location}</h1>

        <div className="w-full overflow-auto rounded-lg">
          <Link
            href={`/${params.username}/location/${params.location}/attraction`}
          >
            <CardTemplete
              imageSource="https://res.cloudinary.com/dopxvbeju/image/upload/v1740377395/attraction_kpgvjr.jpg"
              text="Attraction"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/restaurant`}
          >
            <CardTemplete
              imageSource="https://res.cloudinary.com/dopxvbeju/image/upload/v1740377385/restaurant_rl4pkm.jpg"
              text="Restaurant"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/shopping`}
          >
            <CardTemplete
              imageSource="https://res.cloudinary.com/dopxvbeju/image/upload/v1740377384/shopping_ii3p7t.jpg"
              text="Shopping"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/activity`}
          >
            <CardTemplete
              imageSource="https://res.cloudinary.com/dopxvbeju/image/upload/v1740377384/activity_bnahh9.jpg"
              text="Activity"
            />
          </Link>
          <Link href={`/${params.username}/location/${params.location}/hotel`}>
            <CardTemplete
              imageSource="https://res.cloudinary.com/dopxvbeju/image/upload/v1740377856/hotel_lfx8zv.jpg"
              text="Hotel"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/petrol_station`}
          >
            <CardTemplete
              imageSource="https://res.cloudinary.com/dopxvbeju/image/upload/v1740377382/petrol_station_ecxs86.jpg"
              text="Petrol Station"
            />
          </Link>
        </div>
      </AppTemplate>
    </div>
  );
}
