"use client";

import { useParams } from "next/navigation";
import BookmarkTemplate from "@/app/ui/template/bookmarkTemplate";
import ItemTemplete from "@/app/ui/template/itemTemplate";
import Link from "next/link";

export default function BookmarkLocation() {
  const params = useParams();
  console.log(params);
  const bookmarkLocation = params.location.toUpperCase().replaceAll("%20", " ");
  return (
    <div>
      <BookmarkTemplate place={bookmarkLocation}>
        <h3 className="py-4 text-center text-gray-500">Attraction</h3>
        <hr />
        {/* <div className="pt-8">
          <Link
            href={`/dashboard/location/${params.location}/attraction/fremantle%20arts%20centre`}
          >
            <ItemTemplete
              imageSource="/user_img/default.jpg"
              text="Fremantle Arts Centre"
            ></ItemTemplete>
          </Link>
          <Link
            href={`/dashboard/location/${params.location}/attraction/fremantle%20arts%20centre`}
          >
            <ItemTemplete
              imageSource="/user_img/default.jpg"
              text="Fremantle Arts Centre"
            ></ItemTemplete>
          </Link>{" "}
          <Link
            href={`/dashboard/location/${params.location}/attraction/fremantle%20arts%20centre`}
          >
            <ItemTemplete
              imageSource="/user_img/default.jpg"
              text="Fremantle Arts Centre"
            ></ItemTemplete>
          </Link>
        </div> */}
        <p className="pt-8 text-center ">No Bookmark</p>
      </BookmarkTemplate>
    </div>
  );
}
