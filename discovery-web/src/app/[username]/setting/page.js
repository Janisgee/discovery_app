"use client";
import { useEffect, useState } from "react";
import { CldImage } from "next-cloudinary";
import { useParams } from "next/navigation";
import AppTemplate from "@/app/_ui/template/appTemplate";
import { AvatarUploader } from "@/app/_ui/avatar-uploader/avatar-uploader";
import { fetchProfileImage } from "@/app/_ui/fetchAPI/fetchProfileImage";
import { useRouter } from "next/navigation";

export default function Setting() {
  const [email, setEmail] = useState("");
  const [publicID, setPublicID] = useState("");
  const [alertCurrentPw, setAlertCurrentPw] = useState(false);
  const [alertNewPw, setAlertNewPw] = useState(false);
  const [alertConfirmPw, setAlertConfirmPw] = useState(false);
  const [alertUpdateError, setAlertUpdateError] = useState(false);
  const [alertUpdateSuccess, setAlertUpdateSuccess] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");

  const params = useParams();

  const router = useRouter();

  const defaultImage =
    "https://res.cloudinary.com/dopxvbeju/image/upload/v1740039540/kphottt1vhiuyahnzy8y.jpg";

  const saveAvatar = async (publicID, url) => {
    setPublicID(publicID);
    fetchUpdateUserProfileImage(publicID, url);
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
        console.error(
          "Error fetching user profile image:",
          response.statusText,
        );
        throw new Error(response.statusText);
      }
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
      if (response.ok) {
        const responseData = await response.json();

        setEmail(responseData.user_email);
      } else {
        console.error("Error fetching user profile:", response.statusText);
        throw new Error(response.statusText);
      }
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

      if (!response.ok) {
        // For other errors
        const responseError = await response.json();
        console.log(responseError.error);
        setErrorMsg(capitalizeFirstLetter(responseError.error));
        setAlertUpdateError(true);
        throw new Error(`Failed to fetch: ${responseError.error}`);
      }

      setAlertUpdateSuccess("Password has been successfully updated.");
      // Successfully update password
    } catch (error) {
      console.error("Error fetching update password page:", error);
      // Not Successfully update password
    }
  };

  const fetchData = async () => {
    try {
      const imageData = await fetchProfileImage();

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

    if (!currentPw || !newPw || !confirmedNewPw) {
      if (!currentPw) setAlertCurrentPw(true);
      if (!newPw) setAlertNewPw(true);
      if (!confirmedNewPw) setAlertConfirmPw(true);
      return;
    }
    if (newPw != confirmedNewPw) {
      alert("Please enter same password into the confirm password field.");
      return;
    }
    fetchUpdateNewPw(currentPw, newPw);
  };

  const handleInputOnChange = (e) => {
    e.preventDefault();
    console.log(e.currentTarget.id);
    if (
      e.currentTarget.id == "current_password" &&
      e.currentTarget.value != ""
    ) {
      setAlertCurrentPw(false);
    }
    if (e.currentTarget.id == "new_password" && e.currentTarget.value != "") {
      setAlertNewPw(false);
    }
    if (
      e.currentTarget.id == "confirm_new_password" &&
      e.currentTarget.value != ""
    ) {
      setAlertConfirmPw(false);
    }
  };

  const capitalizeFirstLetter = (sentence) => {
    return String(sentence).charAt(0).toUpperCase() + String(sentence).slice(1);
  };

  useEffect(() => {
    fetchData(router);
    fetchUserProfile();
    // eslint-disable-next-line react-hooks/exhaustive-deps
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
        {}
        <form className="mt-5" onSubmit={handleUpdatePw}>
          {alertUpdateError && (
            <span className="text-s mb-8 italic text-red-500">{errorMsg}</span>
          )}
          {alertUpdateSuccess && (
            <span className="text-s mb-8 italic text-green-500">
              {alertUpdateSuccess}
            </span>
          )}
          <div className="items-left mb-4 flex flex-col">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="current_password"
            >
              Current Password:{" "}
              {alertCurrentPw ? (
                <span className="text-xs italic text-red-500">
                  (Please fill in your current password.)
                </span>
              ) : (
                <span> </span>
              )}
            </label>
            <input
              name="current_password"
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertCurrentPw ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus-within:outline-4 focus-within:outline-indigo-600`}
              id="current_password"
              type="password"
              onChange={handleInputOnChange}
            />
          </div>
          <div className="items-left mb-4 flex flex-col">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="new_password"
            >
              New Password:{" "}
              {alertNewPw ? (
                <span className="text-xs italic text-red-500">
                  (Please fill in your new password.)
                </span>
              ) : (
                <span> </span>
              )}
            </label>
            <input
              name="new_password"
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertNewPw ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus-within:outline-4 focus-within:outline-indigo-600`}
              id="new_password"
              type="password"
              onChange={handleInputOnChange}
            />
          </div>
          <div className="items-left mb-4 flex flex-col">
            <label
              className="mr-3 block font-bold text-gray-700"
              htmlFor="confirm_new_password"
            >
              Confirm New Password:{" "}
              {alertNewPw ? (
                <span className="text-xs italic text-red-500">
                  (Please confirm your new password.)
                </span>
              ) : (
                <span> </span>
              )}
            </label>
            <input
              name="confirm_new_password"
              className={`focus:shadow-outline w-full appearance-none rounded border ${alertConfirmPw ? "border-red-500" : ""} px-3 py-2 leading-tight text-gray-700 shadow focus-within:outline-4 focus-within:outline-indigo-600`}
              id="confirm_new_password"
              type="password"
              onChange={handleInputOnChange}
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
