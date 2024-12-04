import Image from "next/image";

export default function CardTemplete({ imageSource, text }) {
  return (
    <div className="relative mx-auto max-w-xl">
      <Image
        src={imageSource}
        className="mt-5 h-32 w-full rounded-lg object-cover"
        alt="Image of place"
        width={330}
        height={125}
      />
      <div className="absolute inset-0 rounded-lg bg-gray-700 opacity-40"></div>
      <div className="absolute inset-0 flex items-center justify-center">
        <h2 className="text-center text-white">{text}</h2>
      </div>
    </div>
  );
}
