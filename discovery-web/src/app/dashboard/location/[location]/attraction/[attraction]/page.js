"use client";

import AppTemplate from "@/app/ui/template/appTemplate";
import { useParams } from "next/navigation";

export default function AttractionPlace() {
  const params = useParams();
  console.log(params);
  const attraction = params.attraction.toUpperCase().replaceAll("%20", " ");
  return (
    <div>
      <AppTemplate>
        <div>
          <h1 className="text-center text-xl">{attraction}</h1>
        </div>
      </AppTemplate>
    </div>
  );
}
