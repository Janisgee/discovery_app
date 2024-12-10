"use client";

import AppTemplate from "@/app/ui/template/appTemplate";
import ItemTemplete from "@/app/ui/template/itemTemplate";
import { useParams } from "next/navigation";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft, faHeart } from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";

export default function Attraction() {
  const params = useParams();
  console.log(params);
  const location = params.location.toUpperCase().replaceAll("%20", " ");
  return (
    <div>
      <AppTemplate>
        <div className="mb-8 text-center">
          <h2>{location}</h2>
          <div className="grid grid-cols-4">
            <Link href={`/dashboard/location/${params.location}`}>
              <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
            </Link>
            <p className="col-span-2 text-2xl font-bold">[ Attraction ]</p>
          </div>
        </div>
        <Link
          href={`/dashboard/location/${params.location}/attraction/fremantle%20arts%20centre`}
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
