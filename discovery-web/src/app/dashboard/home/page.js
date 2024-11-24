"use client";
import { useRouter } from "next/navigation";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons";

import { Button } from "@/app/ui/buttons";
import HomeTemplate from "@/app/ui/template/homeTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";

export default function Home() {
  const router = useRouter();

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const searchData = formData.get("search");
    if (!searchData) {
      alert("Please enter a location");
      return;
    }
    router.push(`/dashboard/location/${encodeURIComponent(searchData)}`);
  };

  return (
    <div>
      <HomeTemplate>
        <div className="block-center flex-col pb-10">
          <Image
            src="/user_img/default.jpg"
            className="mb-5 h-32 rounded-full"
            alt="default profile picture"
            width={128}
            height={128}
          />
          <h6 className="">Janis Chan</h6>
          <p className="text-color-dark_grey">7 bookmark places</p>
        </div>
        <div className="block-center flex-col">
          <Button useFor="✒️ Plan New Trip" link="/" color="btn-violet" />

          <form onSubmit={handleSearchSubmit}>
            <div className="block-center flex-row">
              <FontAwesomeIcon icon={faMagnifyingGlass} size="2x" />
              <label
                htmlFor="default-search"
                className="sr-only mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >
                Search
              </label>
              <input
                name="search"
                type="search"
                id="default-search"
                className="ml-2 mt-2 block w-full rounded-full border border-gray-300 bg-gray-50 p-2 ps-5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder:text-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
                placeholder="Type location to search"
                required
              />
            </div>
          </form>
        </div>
        <div className="block-center my-5 flex-col ">
          <h3>Popular Destination</h3>
          {/* <div className='w-full h-96 overflow-hidden inline-block rounded-lg'> */}
          <div className="h-96 w-full overflow-auto rounded-lg">
            <CardTemplete
              imageSource="/place_img/paris-france.jpg"
              text="Japan"
            />
            <CardTemplete
              imageSource="/place_img/paris-france.jpg"
              text="Korea"
            />
            <CardTemplete
              imageSource="/place_img/paris-france.jpg"
              text="France"
            />
          </div>
        </div>
      </HomeTemplate>
    </div>
  );
}
