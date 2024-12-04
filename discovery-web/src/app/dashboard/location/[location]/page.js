"use client";

import AppTemplate from "@/app/ui/template/appTemplate";

import { useParams } from "next/navigation";
import CardTemplete from "@/app/ui/template/cardTemplate";

import Link from "next/link";

export default function LocationPlace() {
  const params = useParams();
  console.log(params.location);
  const location = params.location.toUpperCase().replaceAll("%20", " ");
  return (
    <div>
      <AppTemplate>
        <h1 className="text-center">{location}</h1>

        <div className="w-full overflow-auto rounded-lg">
          <Link href="\">
            <CardTemplete
              imageSource="/catagory_img/attraction.jpg"
              text="Attraction"
            />
          </Link>
          <Link href="">
            <CardTemplete
              imageSource="/catagory_img/restaurant.jpg"
              text="Restaurant"
            />
          </Link>
          <CardTemplete
            imageSource="/catagory_img/shopping.jpg"
            text="Shopping"
          />
          <CardTemplete
            imageSource="/catagory_img/activity.jpg"
            text="Activity"
          />
          <CardTemplete imageSource="/catagory_img/hotel.jpg" text="Hotel" />
          <CardTemplete
            imageSource="/catagory_img/petrol_station.jpg"
            text="Petrol Station"
          />
        </div>
      </AppTemplate>
    </div>
  );
}
