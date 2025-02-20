// components/avatar-uploader.tsx
"use client";

import { CldUploadWidget } from "next-cloudinary";

export function AvatarUploader({ onUploadSuccess }) {
  return (
    <CldUploadWidget
      uploadPreset={process.env.NEXT_PUBLIC_CLOUDINARY_UPLOAD_PRESET}
      signatureEndpoint="/api/sign-cloudinary-params"
      onSuccess={(result) => {
        if (typeof result.info === "object" && "secure_url" in result.info) {
          console.log(result.info);
          onUploadSuccess(result.info.public_id, result.info.secure_url);
        }
      }}
      options={{
        singleUploadAutoClose: true,
      }}
    >
      {({ open }) => {
        return (
          <button type="button" onClick={() => open()}>
            Upload Avatar
          </button>
        );
      }}
    </CldUploadWidget>
  );
}
