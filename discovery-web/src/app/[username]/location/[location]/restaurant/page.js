import AppTemplate from "@/app/_ui/template/appTemplate";
import CatagoryTemplate from "@/app/_ui/template/catagoryTemplate";

export default function Restaurant() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="restaurant"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
