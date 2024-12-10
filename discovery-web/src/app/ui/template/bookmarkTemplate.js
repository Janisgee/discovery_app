"use client";

import { useParams } from "next/navigation";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";

import {
  faFileLines,
  faHeart,
  faHouse,
  faBinoculars,
  faBurger,
  faStore,
  faBaseballBatBall,
  faBed,
  faGasPump,
} from "@fortawesome/free-solid-svg-icons";

import { Button } from "@/app/ui/buttons";

export default function BookmarkTemplate({ place, children }) {
  return (
    <div className="background-yellow font-inter">
      <div className="p-5">
        <div>
          <div className="mb-5 flex items-center justify-between">
            <span>
              <span className="mr-5">
                <Link href="/dashboard/home">
                  <FontAwesomeIcon icon={faHouse} size="3x" />
                </Link>
              </span>
              <Link href={`/dashboard/trips`}>
                <FontAwesomeIcon icon={faFileLines} size="3x" />
              </Link>
            </span>
            <span className="flex items-center">
              <Link href={`/dashboard/bookmark`}>
                <FontAwesomeIcon icon={faHeart} size="3x" />
              </Link>
              <span className="ml-2 inline items-start text-xl">
                <Button useFor="JG" link="/dashboard/home" color="btn-grey" />
              </span>
            </span>
          </div>
          <h3 className="mb-5 text-center">Bookmark in {place}</h3>
          <div className="text-center">
            <Button
              useFor={<FontAwesomeIcon icon={faBinoculars} size="xl" />}
              link="/"
              color="btn-white"
            />
            <Button
              useFor={<FontAwesomeIcon icon={faBurger} size="xl" />}
              link="/"
              color="btn-white"
            />
            <Button
              useFor={<FontAwesomeIcon icon={faStore} size="xl" />}
              link="/"
              color="btn-white"
            />
            <Button
              useFor={<FontAwesomeIcon icon={faBaseballBatBall} size="xl" />}
              link="/"
              color="btn-white"
            />
            <Button
              useFor={<FontAwesomeIcon icon={faBed} size="xl" />}
              link="/"
              color="btn-white"
            />
            <Button
              useFor={<FontAwesomeIcon icon={faGasPump} size="xl" />}
              link="/"
              color="btn-white"
            />
          </div>
        </div>
      </div>
      <div className="rounded-t-3xl bg-white ">
        <div className="px-10 pb-10">{children}</div>
      </div>
    </div>
  );
}
