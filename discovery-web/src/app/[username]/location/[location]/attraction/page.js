import AppTemplate from "@/app/_ui/template/appTemplate";
import CatagoryTemplate from "@/app/_ui/template/catagoryTemplate";

export default function Attraction() {
  return (
    <div>
      <AppTemplate>
        <CatagoryTemplate catagory="attraction"></CatagoryTemplate>
      </AppTemplate>
    </div>
  );
}
