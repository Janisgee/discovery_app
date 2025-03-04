import AppTemplate from "@/app/_ui/template/appTemplate";
import CatagoryTemplate from "@/app/_ui/template/catagoryTemplate";

export default function Hotel() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="hotel"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
