import { v2 as cloudinary } from "cloudinary";

cloudinary.config({
  cloud_name: process.env.NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME,
  api_key: process.env.CLOUDINARY_API_KEY,
  api_secret: process.env.CLOUDINARY_API_SECRET,
});

// This route is used to generate the signed upload data
export async function POST(req) {
  const body = await req.json();
  const { paramsToSign } = body;

  // Generate the signature using Cloudinary's SDK
  const signature = cloudinary.utils.api_sign_request(
    paramsToSign,
    process.env.CLOUDINARY_API_SECRET,
  );

  // Send back the signature, timestamp, and upload preset to the client
  return Response.json({ signature });
}
