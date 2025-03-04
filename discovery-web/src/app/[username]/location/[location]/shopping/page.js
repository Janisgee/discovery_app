import AppTemplate from "@/app/_ui/template/appTemplate";
import CatagoryTemplate from "@/app/_ui/template/catagoryTemplate";

export default function Shopping() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="shopping"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
