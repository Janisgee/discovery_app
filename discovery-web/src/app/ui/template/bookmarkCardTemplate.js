import Image from "next/image";

export default function BookmarkCardTemplate() {
  return (
    <div>
      <div className="mt-5 flex items-center justify-between">
        <div>
          <h6 className="inline-block size-8 rounded-full border-2 border-black text-center">
            1
          </h6>
          <Image
            src="/catagory_img/attraction.jpg"
            className="ml-2 inline-block size-24 rounded-lg object-cover"
            alt="Image of place"
            width={90}
            height={90}
          />
        </div>
        <div className="">
          <h6>Tsim Sha Tsui Star</h6>
          <p className="max-w-48">
            UNESCO-listed prison offering intriguing convict history tours.
          </p>
        </div>
      </div>
    </div>
  );
}
