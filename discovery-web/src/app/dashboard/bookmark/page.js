import AppTemplate from "@/app/ui/template/appTemplate";
import CardTemplete from "@/app/ui/template/cardTemplate";
import Link from "next/link";

export default function Bookmark() {
  return (
    <div>
      <AppTemplate>
        <h2 className="text-center">Bookmark</h2>
        <div className="w-full overflow-auto rounded-lg">
          <Link href={`/dashboard/bookmark/perth`}>
            <CardTemplete
              imageSource="/catagory_img/attraction.jpg"
              text="Perth"
            />
          </Link>
          <Link href={`/dashboard/bookmark/perth`}>
            <CardTemplete
              imageSource="/catagory_img/attraction.jpg"
              text="Japan"
            />
          </Link>
          <Link href={`/dashboard/bookmark/hong_kong`}>
            <CardTemplete
              imageSource="/catagory_img/attraction.jpg"
              text="Hong Kong"
            />
          </Link>
          <Link href={`/dashboard/bookmark/germany`}>
            <CardTemplete
              imageSource="/catagory_img/attraction.jpg"
              text="Germany"
            />
          </Link>
        </div>
      </AppTemplate>
    </div>
  );
}
