"use client";

import { useParams } from "next/navigation";
import AppTemplate from "@/app/ui/template/appTemplate";

import CardTemplete from "@/app/ui/template/cardTemplate";

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
              imageSource="/catagory_img/attraction.jpg"
              text="Attraction"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/restaurant`}
          >
            <CardTemplete
              imageSource="/catagory_img/restaurant.jpg"
              text="Restaurant"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/shopping`}
          >
            <CardTemplete
              imageSource="/catagory_img/shopping.jpg"
              text="Shopping"
            />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/activity`}
          >
            <CardTemplete
              imageSource="/catagory_img/activity.jpg"
              text="Activity"
            />
          </Link>
          <Link href={`/${params.username}/location/${params.location}/hotel`}>
            <CardTemplete imageSource="/catagory_img/hotel.jpg" text="Hotel" />
          </Link>
          <Link
            href={`/${params.username}/location/${params.location}/petrol_station`}
          >
            <CardTemplete
              imageSource="/catagory_img/petrol_station.jpg"
              text="Petrol Station"
            />
          </Link>
        </div>
      </AppTemplate>
    </div>
  );
}
