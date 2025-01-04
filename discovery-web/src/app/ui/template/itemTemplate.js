import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faFileLines, faHeart } from "@fortawesome/free-solid-svg-icons";

export default function ItemTemplete({ imageSource, title, text }) {
  return (
    <div className="mb-10 max-w-xl">
      <div className="mb-8 flex items-center justify-between gap-3">
        <div className="size-28">
          <Image
            src={imageSource}
            className="size-full rounded-lg object-cover"
            alt="Image of place"
            width={125}
            height={125}
          />
        </div>
        <span className="flex flex-col justify-between space-y-2">
          <h3>{title}</h3>
          <p>{text}</p>
        </span>
        <span>
          <FontAwesomeIcon icon={faHeart} size="2x" />
        </span>
      </div>
      <hr />
    </div>
  );
}
