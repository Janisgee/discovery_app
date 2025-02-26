import Image from "next/image";
import { useState, useEffect } from "react";
import { fetchPlaceImage } from "@/app/ui/fetchAPI/fetchPlaceImage";

export default function CardTemplete({
  imageSource,
  text,
  searchFor,
  country,
}) {
  const [displayImage, setDisplayImage] = useState(imageSource);
  const fetchData = async () => {
    try {
      const imageURL = await fetchPlaceImage(text, searchFor, country);
      setDisplayImage(imageURL);
    } catch (error) {
      console.error("Error fetching image for country:", error);
    }
  };

  useEffect(() => {
    if (imageSource == "") {
      fetchData();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <div className="relative mx-auto max-w-xl">
      {displayImage && (
        <Image
          src={displayImage}
          className="mt-5 h-32 w-full rounded-lg object-cover"
          alt="Image of place"
          width={330}
          height={125}
        />
      )}

      <div className="absolute inset-0 rounded-lg bg-gray-700 opacity-40"></div>
      <div className="absolute inset-0 flex items-center justify-center">
        <h2 className="text-center text-white">{text}</h2>
      </div>
    </div>
  );
}
