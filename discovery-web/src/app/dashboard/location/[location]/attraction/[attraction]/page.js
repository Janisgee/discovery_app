"use client";

import AppTemplate from "@/app/ui/template/appTemplate";
import { useParams, useEffect } from "next/navigation";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleArrowLeft, faHeart } from "@fortawesome/free-solid-svg-icons";
import Image from "next/image";
import Link from "next/link";

export default function AttractionPlace() {
  const params = useParams();
  console.log(params);
  const attraction = params.attraction.toUpperCase().replaceAll("%20", " ");
  const location = params.location.toUpperCase().replaceAll("%20", " ");

  const fetchSearchCountry = async () => {
    const data = { country: location };

    const request = new Request("http://localhost:8080/searchCountry", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request);

      if (response.ok) {
        const htmlContent = await response.text(); // Use text() to handle HTML response
        console.log(htmlContent);
        router.push(`/dashboard/location/${encodeURIComponent(searchData)}`);
      } else {
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

  return (
    <div>
      <AppTemplate>
        <div className="flex items-center justify-between gap-3 pt-5">
          <Link href={`/dashboard/location/${params.location}/attraction`}>
            <FontAwesomeIcon icon={faCircleArrowLeft} size="2x" />
          </Link>
          <h1 className="text-center text-xl">{attraction}</h1>
          <FontAwesomeIcon icon={faHeart} size="2x" />
        </div>
        <Image
          src="/place_img/paris-france.jpg"
          className="mt-5 h-32 w-full rounded-lg object-cover"
          alt="Image of place"
          width={330}
          height={130}
        />
        <div className="mt-5">
          <p>
            Fremantle Markets is one of Perth’s most iconic and historic
            attractions. Here’s a detailed look at the market
          </p>
          <div className="mt-5">
            <h6>Location:</h6>
            <ul>
              <li>
                Address: Corner of South Terrace and Henderson Street,
                Fremantle, WA.
              </li>
              <li>Opening Hours:</li>
            </ul>
          </div>{" "}
          <div className="mt-5">
            <h6>History:</h6>
            <ul>
              <li>Established: 1897</li>
              <li>
                Originally built as a fresh produce market, the Fremantle
                Markets were an essential trading spot for local farmers and
                traders. Over the years, it transitioned into a vibrant cultural
                and social space, housing stalls for local artisans,
                craftspeople, and food vendors.
              </li>
            </ul>
          </div>
          <div className="mt-5">
            <h6>Key Features:</h6>
            <ul>
              <li>
                Stalls: The markets house over 150 stalls offering an eclectic
                mix of fresh food, artisan products, clothing, homewares, and
                handmade goods.
              </li>
              <li>
                Food and Drink: Expect a variety of local, organic, and
                international cuisines, with a focus on fresh and sustainable
                options. Popular food items include Asian street food, fresh
                juices, gourmet coffee, and local snacks.
              </li>
            </ul>
          </div>
          <div className="mt-5">
            <p>
              Fremantle Markets remains a must-visit for both tourists and
              locals, offering a dynamic and vibrant atmosphere while reflecting
              the local culture and heritage.
            </p>
          </div>
        </div>
      </AppTemplate>
    </div>
  );
}
