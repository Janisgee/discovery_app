"use client";
import { useParams } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";

import {
  faFileLines,
  faHeart,
  faHouse,
  faRightFromBracket,
} from "@fortawesome/free-solid-svg-icons";

export default function AppTemplate({ children }) {
  const params = useParams();
  return (
    <div className="background-yellow font-inter">
      <div className="p-5">
        <div className="block-center">
          <div className="flex w-full max-w-sm justify-between">
            <div>
              <span className="mr-5">
                <Link href={`/${params.username}/home`}>
                  <FontAwesomeIcon icon={faHouse} size="3x" />
                </Link>
              </span>
              <Link href={`/${params.username}/trips`}>
                <FontAwesomeIcon icon={faFileLines} size="3x" />
              </Link>
            </div>
            <div className="flex items-center justify-center">
              <Link href={`/${params.username}/bookmark`}>
                <FontAwesomeIcon icon={faHeart} size="3x" />
              </Link>
              <span className="ml-5">
                <Link href="/">
                  <FontAwesomeIcon icon={faRightFromBracket} size="3x" />
                </Link>
              </span>
            </div>
          </div>
        </div>
      </div>
      <div className="rounded-t-3xl bg-white ">
        <div className="px-10 pb-10">{children}</div>
      </div>
    </div>
  );
}
