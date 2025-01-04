import AppTemplate from "@/app/ui/template/appTemplate";
import CatagoryTemplate from "@/app/ui/template/catagoryTemplate";

export default function Restaurant() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="restaurant"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
