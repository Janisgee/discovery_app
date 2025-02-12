"use client";

import { useEffect, useState } from "react";

import ItemTemplete from "@/app/ui/template/itemTemplate";
import { useParams, useRouter } from "next/navigation";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft } from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";

export default function CatagoryTemplate({ catagory }) {
  const [content, setContent] = useState([]);
  const params = useParams();
  const router = useRouter();

  const location = decodeURIComponent(params.location).toUpperCase();
  const twoWordCatagory = catagory.replaceAll("_", " ");

  const fetchSearchCountry = async () => {
    const data = { country: location, catagory: twoWordCatagory };
    console.log("country:", location, "catagory:", twoWordCatagory);

    const request = new Request("http://localhost:8080/searchCountry", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request, { next: { revalidate: 600 } });

      if (response.ok) {
        const htmlContent = await response.json(); // Use json() to handle HTML response
        console.log(htmlContent);
        setContent(htmlContent);
      } else {
        if (response.status == 401) {
          alert(
            `Please login again as 10 mins session expired without taking action.`,
          );
          router.push(`/login`);
        }

        console.error("Error fetching search country:", response.statusText);
      }
    } catch (error) {
      console.error("Error fetching search country:", error);
    }
  };

  useEffect(() => {
    fetchSearchCountry();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  let itemList = [];
  if (content.length > 1) {
    content.forEach((item, index) => {
      itemList.push(
        <li key={index}>
          <ItemTemplete
            imageSource={item.image}
            title={item.name}
            text={item.description}
            placeID={item.place_id}
            catagory={catagory}
            hasbookmark={item.hasbookmark}
          ></ItemTemplete>
        </li>,
      );
    });
  }

  return (
    <div>
      <div className="mb-8 text-center">
        <h2>{location}</h2>
        <div className="grid grid-cols-4">
          <Link href={`/${params.username}/location/${params.location}`}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </Link>
          <p className="col-span-2 text-2xl font-bold">
            [ {twoWordCatagory.toUpperCase()} ]
          </p>
        </div>
      </div>
      <ul>{itemList}</ul>
    </div>
  );
}
