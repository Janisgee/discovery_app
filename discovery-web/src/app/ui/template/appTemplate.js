import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import {
  faFileLines,
  faHeart,
  faHouse,
} from "@fortawesome/free-solid-svg-icons";

import { Button } from "@/app/ui/buttons";

export default function AppTemplate({ children }) {
  return (
    <div className="background-yellow font-inter">
      <div className="p-5">
        <div className="block-center">
          <div className="flex w-full max-w-sm justify-between">
            <div>
              <span className="mr-5">
                <FontAwesomeIcon icon={faHouse} size="3x" />
              </span>
              <FontAwesomeIcon icon={faFileLines} size="3x" />
            </div>
            <div className="flex items-center justify-center">
              <FontAwesomeIcon icon={faHeart} size="3x" />
              <span className=" ml-2 inline items-start text-xl">
                <Button useFor="JG" link="/dashboard/home" color="btn-grey" />
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
