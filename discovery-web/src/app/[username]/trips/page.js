"use client";
import { useParams } from "next/navigation";
import { DatePicker } from "@nextui-org/date-picker";
import AppTemplate from "@/app/ui/template/appTemplate";
import { today } from "@internationalized/date";
import { Button } from "@/app/ui/buttons";

export default function Trips() {
  const params = useParams();
  return (
    <div>
      <AppTemplate>
        <div>
          <h2 className="mb-16 text-center">Plan a new Trip</h2>

          <form className="mx-auto mb-7 max-w-sm ">
            <label
              htmlFor="countries"
              className="mb-2 block text-lg font-medium text-gray-900 dark:text-white"
            >
              Where to :
            </label>
            <select
              id="countries"
              className="block w-full rounded-lg border border-gray-300 bg-gray-50 p-4 text-lg text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder:text-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
            >
              <option>United States</option>
              <option>Canada</option>
              <option>France</option>
              <option>Germany</option>
            </select>
          </form>

          <div className="flex flex-row rounded-lg border-2 p-4">
            <div className="flex w-full flex-col">
              <DatePicker
                className=""
                label="Date from:"
                defaultValue={today}
              />{" "}
            </div>
            <div className="flex w-full flex-col">
              <DatePicker
                className=""
                // defaultValue={parseDate("2024-04-04")}
                label="Date to:"
              />
            </div>
          </div>
          <div className="mt-24 text-center">
            <Button
              useFor="✒️ Start Planning"
              link={`/${params.username}/trips/hong%20kong`}
              color="btn-violet"
            />
          </div>
        </div>
      </AppTemplate>
    </div>
  );
}
