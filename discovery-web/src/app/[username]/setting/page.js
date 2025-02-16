"use client";
import { CldUploadButton } from "next-cloudinary";
import { useParams } from "next/navigation";

import AppTemplate from "@/app/ui/template/appTemplate";

import Image from "next/image";

export default function Setting() {
  const params = useParams();
  const cloudName = process.env.NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME;
  console.log("Cloudinary Cloud Name:", cloudName);
  return (
    <div>
      <AppTemplate>
        <div className="block-center flex-col pb-8">
          <Image
            src="/user_img/default.jpg"
            className="h-32 rounded-full"
            alt="default profile picture"
            width={128}
            height={128}
          />
          <div className="btn-violet px-3 py-1">
            <CldUploadButton uploadPreset="<Upload Preset>">
              Edit
            </CldUploadButton>
          </div>
        </div>
        <h3 className="justify-items-start text-gray-400">Profile Settings</h3>
        <div className="mb-8 mt-5">
          <div className="mb-2 flex items-center">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="username"
            >
              Username:
            </label>
            <p className="w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none">
              {params.username}
            </p>
          </div>
          <div className="mb-2 flex items-center">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="username"
            >
              Email:
            </label>
            <p className="w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:outline-none">
              {params.username}*****
            </p>
          </div>

          {/* <div className="flex items-center rounded-md bg-white pl-3 outline-2 outline-sky-500 focus-within:outline-4 focus-within:outline-indigo-600">
            <label
              htmlFor="username"
              className="block font-medium text-gray-900"
            >
              Username:
            </label>
            <input
              type="text"
              name="username"
              id="username"
              className="ml-5 w-full py-1.5 pl-5 pr-3 text-base text-gray-900 placeholder:text-gray-400 focus:outline-none"
              placeholder={params.username}
            ></input>
          </div> */}
        </div>
        <hr />
        <h3 className="mt-8 justify-items-start text-gray-400">
          Change Password
        </h3>
        <form className="mt-5">
          {" "}
          <div className="items-left mb-4 flex flex-col">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="current_password"
            >
              Current Password:
            </label>
            <input
              name="current_password"
              className="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus-within:outline-4 focus-within:outline-indigo-600"
              id="current_password"
              type="text"
              placeholder="******************"
            />
          </div>
          <div className="items-left mb-4 flex flex-col">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="new_password"
            >
              New Password:
            </label>
            <input
              name="new_password"
              className="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus-within:outline-4 focus-within:outline-indigo-600"
              id="new_password"
              type="text"
            />
          </div>
          <div className="flex justify-center">
            <button className="btn-violet">Update Password</button>
          </div>
        </form>
      </AppTemplate>
    </div>
  );
}
