import AppTemplate from "@/app/ui/template/appTemplate";
import CatagoryTemplate from "@/app/ui/template/catagoryTemplate";

export default function Shopping() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="shopping"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
