"use client";

import AppTemplate from "@/app/ui/template/appTemplate";
import { Button } from "@/app/ui/buttons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBinoculars } from "@fortawesome/free-solid-svg-icons";
import { useParams } from "next/navigation";

export default function TripLocation() {
  const params = useParams();
  console.log(params);
  const tripLocation = params.location.toUpperCase().replaceAll("%20", " ");

  return (
    <div>
      <AppTemplate>
        <div className="text-center">
          <h2>Trip to {tripLocation}</h2>
          <p className="mt-2"> 15 Dec - 31 Dec</p>
        </div>
        <div>
          <div className="my-5 grid grid-cols-2">
            <h6>Note:</h6>
            <span className="text-right">
              <Button useFor="Add Note" link="/" color="btn-violet" />
            </span>
          </div>
          <p>UNESCO-listed prison offering intriguing convict history tours.</p>
        </div>
        <div className="my-5">
          <hr />
        </div>
        <div className="grid grid-cols-2">
          <p> 15 Dec</p>
          <span className="text-right">
            <Button useFor="Add Place" link="/" color="btn-violet" />
          </span>
        </div>
        <div className="my-5">
          <hr />
        </div>
        <div className="grid grid-cols-2">
          <p> 16 Dec</p>
          <span className="text-right">
            <Button useFor="Add Place" link="/" color="btn-violet" />
          </span>
          <div>
            <h6 className="size-8 rounded-full border-2 border-black text-center">
              1
            </h6>
          </div>
        </div>
      </AppTemplate>
    </div>
  );
}
