"use client";
import { useEffect, useState } from "react";
import { CldImage } from "next-cloudinary";
import { useParams } from "next/navigation";
import AppTemplate from "@/app/ui/template/appTemplate";
import { AvatarUploader } from "@/app/ui/avatar-uploader/avatar-uploader";

import Image from "next/image";
import { revalidatePath } from "next/cache";

export default function Setting() {
  const [email, setEmail] = useState("");
  const [publicID, setPublicID] = useState("/user_img/default.jpg");
  const [imageSource, setImageSource] = useState("/user_img/default.jpg");
  const params = useParams();

  // const cloudName = process.env.NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME;
  // const cloudUploadPresent = process.env.CLOUDINARY_UPLOAD_PRESET;
  // console.log("Cloudinary Cloud Name:", cloudName);
  //  src={imageData.secureUrl} alt="Uploaded"
  const signatureEndpoint = "/api/sign-cloudinary-params";
  const saveAvatar = async (publicID, url) => {
    setImageSource(url);
    setPublicID(publicID);
    console.log("Uploaded image url:", url);
    console.log("Uploaded public id:", publicID);
    // revalidatePath("/");
  };

  // Fetch publicID to backendserver to do storage

  // Fetch user email
  // const handleUploadSuccess = (result) => {
  //   // Log the result for debugging purposes
  //   console.log("Upload result", result);
  //   // You can access public_id and secure_url here
  //   const { public_id, secure_url } = result.info;
  //   // Save or use public_id and secure_url as needed
  //   setImageData({
  //     publicId: public_id,
  //     secureUrl: secure_url,
  //   });
  //   setImageSource(secure_url);
  // };

  // const handleUploadError = (error) => {
  //   console.error("Upload failed", error);
  //   alert("An error occurred while uploading. Please try again.");
  // };

  const fetchUserProfile = async () => {
    const request = new Request("http://localhost:8080/api/getUserProfile", {
      method: "GET", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });

    try {
      const response = await fetch(request);
      if (!response.ok) {
        throw new Error(`Failed to fetch: ${response.statusText}`);
      }

      const responseData = await response.json();
      console.log("Server Response:", responseData);
      console.log("Server Response:", responseData.user_email);
      setEmail(responseData.user_email);
    } catch (error) {
      console.error("Error fetching user profile:", error);
    }
  };

  const fetchUpdateNewPw = async (currentPw, newPw) => {
    const data = {
      currentPw: currentPw,
      newPw: newPw,
    };
    const request = new Request("http://localhost:8080/api/updatePassword", {
      method: "POST", // HTTP method
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    try {
      const response = await fetch(request);
      console.log(response.status);
      if (!response.ok) {
        // Not Successfully update password
        if (response.status == 401) {
          router.push("/unauthorized-link");
          return;
        } else {
          // For other errors, handle as text (non-JSON responses)
          const errorText = await response.text();
          throw new Error(`Failed to fetch: ${errorText}`);
        }
      }

      const responseData = await response.json();
      console.log("Server Response:", responseData);
      alert("Password has been successfully updated.");
      // Successfully update password
    } catch (error) {
      console.error("Error fetching update password page:", error);
      // Not Successfully update password
      alert(
        "Fail to update user password. Please try a stronger password pattern.",
      );
    }
  };

  const handleUpdatePw = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const currentPw = formData.get("current_password");
    const newPw = formData.get("new_password");
    const confirmedNewPw = formData.get("confirm_new_password");

    if (newPw != confirmedNewPw) {
      alert("Please enter smae password into the confirm password field.");
      return;
    }
    fetchUpdateNewPw(currentPw, newPw);
  };

  useEffect(() => {
    fetchUserProfile();
  }, []);
  return (
    <div>
      <AppTemplate>
        <div className="block-center flex-col pb-5">
          {publicID == "/user_img/default.jpg" ? (
            <Image
              src="/user_img/default.jpg"
              className="h-32 rounded-full"
              alt={params.username}
              width={128}
              height={128}
            />
          ) : (
            <CldImage
              src={publicID}
              className="h-32 rounded-full"
              alt={params.username}
              width="128"
              height="128"
              crop="fill"
              gravity="face"
            />
          )}
          <div className="btn-violet px-3 py-1">
            <AvatarUploader onUploadSuccess={saveAvatar} />
            {/* <CldUploadWidget
              
              uploadPreset={cloudUploadPresent}
              signatureEndpoint={signatureEndpoint}
              onSuccess={(result)=>
                if(typeof result.info =="object" && "secure_url" in result.info){

                }
                {handleUploadSuccess}}
              onFailure={handleUploadError}
              onQueuesEnd={(result, { widget }) => {
                widget.close();
              }}
            >
              {({ open }) => {
                function handleOnClick(e) {
                  e.preventDefault();
                  open();
                }
                return (
                  <button id="upload_widget" onClick={handleOnClick}>
                    Edit
                  </button>
                );
              }}
            </CldUploadWidget> */}
          </div>
        </div>
        {/* <p>PublicID: {publicID}</p> */}
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
              {email}
            </p>
          </div>
        </div>
        <hr />
        <h3 className="mt-5 justify-items-start text-gray-400">
          Change Password
        </h3>
        <form className="mt-5" onSubmit={handleUpdatePw}>
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
              type="password"
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
              type="password"
            />
          </div>
          <div className="items-left mb-4 flex flex-col">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="confirm_new_password"
            >
              Confirm New Password:
            </label>
            <input
              name="confirm_new_password"
              className="focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus-within:outline-4 focus-within:outline-indigo-600"
              id="confirm_new_password"
              type="password"
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
