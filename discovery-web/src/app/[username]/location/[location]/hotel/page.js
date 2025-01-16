import AppTemplate from "@/app/ui/template/appTemplate";
import CatagoryTemplate from "@/app/ui/template/catagoryTemplate";

export default function Hotel() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="hotel"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
