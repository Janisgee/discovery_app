"use client";

import { useEffect, useState } from "react";

import ItemTemplete from "@/app/ui/template/itemTemplate";
import { useParams, useRouter } from "next/navigation";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft } from "@fortawesome/free-solid-svg-icons";

import { LoadingSpinner } from "@/app/ui/loading-spinner";
import Link from "next/link";

export default function CatagoryTemplate({ catagory }) {
  const [isPending, setIsPending] = useState(false);
  const [content, setContent] = useState([]);
  const params = useParams();

  const location = decodeURIComponent(params.location).toUpperCase();
  const twoWordCatagory = catagory.replaceAll("_", " ");

  const fetchSearchCountry = async () => {
    const data = { country: location, catagory: twoWordCatagory };
    console.log("country:", location, "catagory:", twoWordCatagory);
    setIsPending(true);

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
        console.error("Error fetching search country:", response.statusText);
        throw error;
      }
    } catch (error) {
      console.error("Error fetching search country:", error);
    } finally {
      setIsPending(false);
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
      {isPending ? (
        <div className="mt-72 flex items-center justify-center">
          <LoadingSpinner size={48} color="text-violet-600" />
        </div>
      ) : (
        <ul>{itemList}</ul>
      )}
    </div>
  );
}
