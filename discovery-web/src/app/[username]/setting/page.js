"use client";
import { useEffect, useState } from "react";
import { CldImage } from "next-cloudinary";
import { useParams } from "next/navigation";
import AppTemplate from "@/app/ui/template/appTemplate";
import { AvatarUploader } from "@/app/ui/avatar-uploader/avatar-uploader";
import { fetchProfileImage } from "@/app/ui/fetchAPI/fetchProfileImage";

// import { revalidatePath } from "next/cache";

export default function Setting() {
  const [email, setEmail] = useState("");
  const [publicID, setPublicID] = useState("");
  const [imageSource, setImageSource] = useState("");
  const params = useParams();

  const defaultImage =
    "https://res.cloudinary.com/dopxvbeju/image/upload/v1740039540/kphottt1vhiuyahnzy8y.jpg";

  const saveAvatar = async (publicID, url) => {
    setImageSource(url);
    setPublicID(publicID);
    console.log("Uploaded image url:", url);
    console.log("Uploaded public id:", publicID);
    fetchUpdateUserProfileImage(publicID, url);
    // revalidatePath("/");
  };

  // Fetch publicID to backendserver to do storage
  const fetchUpdateUserProfileImage = async (publicID, url) => {
    const data = {
      public_id: publicID,
      secure_url: url,
    };
    const request = new Request(
      "http://localhost:8080/api/updateUserProfileImage",
      {
        method: "POST", // HTTP method
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(data),
      },
    );

    try {
      const response = await fetch(request);
      if (!response.ok) {
        throw new Error(`Failed to fetch: ${response.statusText}`);
      }

      const responseData = await response.json();
      console.log("Server Response:", responseData);
    } catch (error) {
      console.error("Error fetching user new profile image:", error);
    }
  };

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
      console.error(
        "Error fetching and store user profile into database:",
        error,
      );
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

  const fetchData = async () => {
    try {
      const imageData = await fetchProfileImage();
      console.log(imageData);
      if (imageData != null) {
        setPublicID(imageData.publicID);
        setImageSource(imageData.secureURL);
      }
    } catch (error) {
      console.error("Error fetching data:", error);
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
    fetchData();
    fetchUserProfile();
  }, []);
  return (
    <div>
      <AppTemplate>
        <div className="block-center flex-col pb-5">
          {!publicID ? (
            <CldImage
              src={defaultImage}
              className="h-32 rounded-full"
              alt={params.username}
              width="128"
              height="128"
              crop="fill"
              gravity="face"
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
