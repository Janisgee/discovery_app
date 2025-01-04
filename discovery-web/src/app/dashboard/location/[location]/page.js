"use client";
import { useState } from "react";
import { useParams } from "next/navigation";
import AppTemplate from "@/app/ui/template/appTemplate";

import CardTemplete from "@/app/ui/template/cardTemplate";

import Link from "next/link";

export default function LocationPlace() {
  const [catagory, setCatagory] = useState("");
  const params = useParams();
  const location = params.location.toUpperCase().replaceAll("%20", " ");

  // const fetchSearchCountry = async () => {
  //   const data = { country: location };

  //   const request = new Request("http://localhost:8080/searchCountry", {
  //     method: "POST", // HTTP method
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify(data),
  //   });

  //   try {
  //     const response = await fetch(request);

  //     if (response.ok) {
  //       const htmlContent = await response.text(); // Use text() to handle HTML response
  //       console.log(htmlContent);
  //       router.push(`/dashboard/location/${encodeURIComponent(searchData)}`);
  //     } else {
  //       console.error("Error fetching search country:", response.statusText);
  //     }
  //   } catch (error) {
  //     console.error("Error fetching search country:", error);
  //   }
  // };

  return (
    <div>
      <AppTemplate>
        <h1 className="text-center">{location}</h1>

        <div className="w-full overflow-auto rounded-lg">
          <Link href={`/dashboard/location/${params.location}/attraction`}>
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
