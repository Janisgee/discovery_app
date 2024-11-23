import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import { faFileLines, faHeart } from "@fortawesome/free-solid-svg-icons";

export default function AppTemplate({ children }) {
  return (
    <div className="background-yellow font-inter">
      <div className="px-10 py-5">
        <div className="block-center">
          <div className="flex w-full max-w-sm justify-between">
            <FontAwesomeIcon icon={faFileLines} size="3x" />
            <FontAwesomeIcon icon={faHeart} size="3x" />
          </div>
        </div>
      </div>
      <div className="rounded-t-3xl bg-white ">
        <div className="px-10 pb-10">{children}</div>
      </div>
    </div>
  );
}
