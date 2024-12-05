"use client";

import AppTemplate from "@/app/ui/template/appTemplate";
import ItemTemplete from "@/app/ui/template/itemTemplate";
import { useParams } from "next/navigation";
import Link from "next/link";

export default function Attraction() {
  const params = useParams();
  console.log(params);
  return (
    <div>
      <AppTemplate>
        <div className="mb-8 text-center">
          <h2>{params.location}</h2>
          <p className="text-2xl font-bold">[ Attraction ]</p>
        </div>
        <Link
          href={`/dashboard/location/${params.location}/attraction/fremantle_arts_centre`}
        >
          <ItemTemplete
            imageSource="/user_img/default.jpg"
            text="Fremantle Arts Centre"
          ></ItemTemplete>
        </Link>
        <ItemTemplete
          imageSource="/user_img/default.jpg"
          text="Fremantle Arts Centre"
        ></ItemTemplete>
        <ItemTemplete
          imageSource="/user_img/default.jpg"
          text="Fremantle Arts Centre"
        ></ItemTemplete>
        <ItemTemplete
          imageSource="/user_img/default.jpg"
          text="Fremantle Arts Centre"
        ></ItemTemplete>
        <ItemTemplete
          imageSource="/user_img/default.jpg"
          text="Fremantle Arts Centre"
        ></ItemTemplete>
      </AppTemplate>
    </div>
  );
}
